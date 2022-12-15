package main

import (
	telnet "dev10/pkg"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeout          time.Duration
	ErrNotEnoughArgs = errors.New("not enough arguments, should be 3 at least")
)

func main() {

	var stringDuration string

	flag.StringVar(&stringDuration, "timeout", "10s", "connection timeout")

	flag.Parse()
	if (len(os.Args) < 3) || (len(os.Args) > 4) {
		log.Fatal(ErrNotEnoughArgs)
	}

	t, err := stringTime(stringDuration)
	if err != nil {
		timeout = 10 * time.Second
	} else {
		timeout = t
	}

	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]

	c := telnet.NewClient(
		net.JoinHostPort(host, port),
		os.Stdin,
		os.Stdout,
		timeout,
	)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	err = c.Dial()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	errs := make(chan error, 1)
	go func() { errs <- c.Write() }()
	go func() { errs <- c.Read() }()

	select {
	case <-quit:
		signal.Stop(quit)
		return
	case err = <-errs:
		if err != nil {
			log.Fatal(err)
		}
		telnet.ErrLog.Println("...EOF")
		return
	}
}

func stringTime(t string) (time.Duration, error) {
	td, err := time.ParseDuration(t)
	return td, err
}
