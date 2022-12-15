package server

import (
	"context"
	"dev11/conf"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Server *http.Server
	Conf   conf.Configuration
	Logger *log.Logger
}

func NewServer(configuration conf.Configuration, logger *log.Logger) *Server {
	s := Server{Conf: configuration, Logger: logger}
	if s.Conf.Port == "" || s.Conf.Host == "" {
		s.Conf.Host = "127.0.0.1"
		s.Conf.Port = ":80"
	}

	s.Server = &http.Server{
		Addr: s.Conf.Host + ":" + s.Conf.Port,
	}

	return &s
}

func (s *Server) Init(sm *http.ServeMux) error {
	s.Server.Handler = sm
	s.Logger.Println("init server")
	return nil
}

func (s *Server) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := s.Server.ListenAndServe()
		if err == http.ErrServerClosed {
			s.Logger.Println(err)
		}
	}()

	s.Logger.Printf("Server is listening on HOST: %s PORT: %s", s.Conf.Host, s.Conf.Port)

	////***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}

	log.Println("Server Exited Property")

	return nil
}

func (s *Server) Stop() error {
	s.Logger.Println("stop server")
	s.Server.Shutdown(context.TODO())
	return nil
}
