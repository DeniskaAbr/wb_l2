package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type App struct {
	rootURL string
	logger  *log.Logger
}

func NewApp(rootURL string, logger *log.Logger) *App {
	app := App{
		rootURL: rootURL,
		logger:  logger,
	}

	return &app
}

var url string

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// logger.SetOutput(io.Discard)

	if len(os.Args) < 2 {
		logger.Println("не указан адрес сайта")
		return
	} else {
		url = os.Args[1]
	}

	myApp := NewApp(url, logger)

	data, err := myApp.GetData(myApp.rootURL)
	if err != nil {
		return
	}

	myApp.SaveFile("index.html", data)
	if err != nil {
		return
	}
}

func (app *App) GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		app.logger.Println("ошибка чтения данных сайта")
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		app.logger.Println("ошибка чтения скачаных данных")
		return nil, err
	}

	resp.Body.Close()

	return data, nil
}

func (app *App) SaveFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		app.logger.Println("ошибка сохранения данных в файл")
		return err
	}

	return nil
}
