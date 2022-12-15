package pkg

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type IApp interface {
	Init()
	Start()
	Stop()
	AppMain()
}

// App
type App struct {
	Name    string
	Logger  *log.Logger
	ctx     context.Context
	cancel  context.CancelFunc
	AppMain func(app *App)
}

// NewApp
func NewApp(name string, logger *log.Logger) *App {
	app := &App{Name: name, Logger: logger}
	return app
}

// Init
func (a *App) Init() {

}

// Start
func (a *App) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	a.ctx = ctx
	a.cancel = cancel
	defer a.cancel()

	go a.appRunMain()

	a.аppGracefulShutdown()
	return
}

// Stop
func (a *App) Stop() {
	a.cancel()
}

// аppGracefulShutdown
func (a *App) аppGracefulShutdown() {
	a.Logger.Printf("Application \"%s\" wait shutdown", a.Name)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		a.Logger.Printf("Application \"%s\" %v by signal.Notify \n", a.Name, v)
		a.cancel()
		return
	case done := <-a.ctx.Done():
		a.Logger.Printf("Application \"%s\" root context is done: %v \n", a.Name, done)
		return
	}

	a.Logger.Printf("Application \"%s\"  exited propertly \n", a.Name)
}

func (a *App) appRunMain() {
	var fu func(app *App)
	a.Logger.Printf("Application \"%s\" run main \n", a.Name)
	defer a.cancel()
	if a.AppMain == nil {

		fu = func(app *App) {
			a.Logger.Printf("Application \"%s\" payload is empty, gopher is chill \n", a.Name)
			for i := 0; i < 1000000000000; i++ {
			}
		}
		a.AppMain = fu
	}
	a.AppMain(a)
}
