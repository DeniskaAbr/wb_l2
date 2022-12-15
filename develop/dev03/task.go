package main

import (
	_app "dev03/pkg"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var myapp *_app.App

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// logger.SetOutput(io.Discard)

	conf, err := readParams()
	if err != nil {
		logger.Printf(err.Error())
		os.Exit(2)
	}

	myapp = _app.NewApp("MyTestApp", logger)
	myapp.Init()

	myapp.AppMain = func(a *_app.App) {
		myPayload := Payload{app: myapp, config: conf}
		myPayload.Main()
	}

	myapp.Start()

}

type Config struct {
	columnForSort     int  // -k — указание колонки для сортировки
	sortByNumberValue bool // -n — сортировать по числовому значению
	sortReverse       bool // -r — сортировать в обратном порядке
	excludeDoubles    bool // -u — не выводить повторяющиеся строки

	sortByMonth                   bool // -M — сортировать по названию месяца
	ignoreTrailingSpaces          bool // -b — игнорировать хвостовые пробелы
	checkIfSorted                 bool // -c — проверять отсортированы ли данные
	sortByNumberValueWithSuffixes bool // -h — сортировать по числовому значению с учётом суффиксов

	operatedFilePath string
}

type Payload struct {
	app    *_app.App
	config *Config
}

func (pl *Payload) Main() {
	fmt.Println("my main")

	strings, err := pl.readFile(pl.config.operatedFilePath)
	if err != nil {
		pl.app.Logger.Printf("Application \"%s\" error open file: %v \n", pl.app.Name, pl.config.operatedFilePath)
		pl.app.Stop()
	}

	//fmt.Println(strings)
	//for _, s := range strings {
	//	fmt.Println(s)
	//}

	strings, err = pl.dataProcessing(strings)

	if err != nil {
		pl.app.Logger.Printf("Application \"%s\" error open file: %v \n", pl.app.Name, pl.config.operatedFilePath)
		pl.app.Stop()
	}
}

func (pl *Payload) readFile(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	strings := strings.Split(string(bytes), "\n")

	return strings, nil
}

func (pl *Payload) dataProcessing(ss []string) ([]string, error) {

	//  cfg := pl.config

	return []string{}, nil
}

func readParams() (*Config, error) {
	fmt.Println("read params")
	conf := Config{}

	flag.IntVar(&conf.columnForSort, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&conf.sortByNumberValue, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&conf.sortReverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&conf.excludeDoubles, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&conf.sortByMonth, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&conf.ignoreTrailingSpaces, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&conf.checkIfSorted, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&conf.sortByNumberValueWithSuffixes, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.Parse()

	arguments := flag.Args()

	if len(arguments) > 0 && arguments[0] == "help" || len(arguments) == 0 {
		flag.Usage()
		return &conf, errors.New("attributes parse error, please see manual")
	}

	conf.operatedFilePath = arguments[len(arguments)-1]

	return &conf, nil
}

func removeDoubles(s []string) []string {
	var unicalStringKeys map[string]struct{}

	unical := make([]string, 0)
	unicalStringKeys = make(map[string]struct{})
	for _, str := range s {
		_, ok := unicalStringKeys[str]
		if !ok {
			unical = append(unical, str)
			unicalStringKeys[str] = struct{}{}
		}
	}
	return unical
}

func sorter() {
	
}