package telnet

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const network = "tcp"

var (
	ErrConnectionClosed = errors.New("connection closed by peer")
	ErrLog              = log.New(os.Stderr, "", 0)
)

type Client struct {
	address string
	timeout time.Duration
	conn    net.Conn
	reader  io.ReadCloser
	writer  io.Writer
}

type IClient interface {
	Dial() error
	Read() error
	Write() error
	Close() error
}

func transferData(in io.Reader, out io.Writer) error {
	_, err := io.Copy(out, in)
	if err != nil {
		return ErrConnectionClosed
	}
	return nil
}

func (c *Client) Dial() error {
	conn, err := net.DialTimeout(network, c.address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	ErrLog.Println("...Connected to " + c.address)
	return nil
}

func (c *Client) Write() error {
	return transferData(c.reader, c.conn)
}

func (c *Client) Read() error {
	return transferData(c.conn, c.writer)
}

func (c *Client) Close() error {
	err := c.reader.Close()
	if err != nil {
		return err
	}
	err = c.conn.Close()
	return err
}

func NewClient(address string, reader io.ReadCloser, writer io.Writer, timeout time.Duration) IClient {
	return &Client{
		address: address,
		reader:  reader,
		writer:  writer,
		timeout: timeout,
	}
}
