package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	fields    string
	delimiter string
	separated bool
}

type App struct {
	config  Config
	columns []int
}

func NewApp() *App {
	conf := Config{}
	app := App{
		config: conf,
	}

	return &app
}

func (app *App) Init() {
	flag.StringVar(&app.config.fields, "fields", "0", "columns for extract")
	flag.StringVar(&app.config.delimiter, "delimiter", "\t", "delimiter symbol")
	flag.BoolVar(&app.config.separated, "separeated", false, "only srings with delimiter")

	flag.Parse()

	if app.config.fields != "" {
		var col []int
		s := strings.Split(app.config.fields, " ")
		for _, s := range s {
			i, err := strconv.Atoi(s)
			if err != nil {
			} else {
				col = append(col, i)
			}
		}

		app.columns = col
		sort.Ints(app.columns)
	}
}

func (app *App) Scanner() string {

	output := strings.Builder{}

	_, err := os.Stdin.Stat()
	var source io.Reader
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	source = os.Stdin

	lines, err := app.Reader(source)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	for _, line := range lines {
		if app.config.separated && strings.Contains(line, app.config.delimiter) {
			s, err := app.Cut(&line)
			if err != nil {
				os.Exit(1)
			}
			output.WriteString(s + "\n")
		} else {
			s, err := app.Cut(&line)
			if err != nil {
				os.Exit(1)
			}
			output.WriteString(s + "\n")
		}

	}

	return output.String()

	//	s := bufio.NewScanner(os.Stdin)
	//	for s.Scan() && s.Text() != "stop" {
	//		a := s.Text()
	//		l, err := app.Cut(&a)
	//		if err != nil {
	//			os.Exit(1)
	//		}
	//		fmt.Println(l)
	//	}
}

func (app *App) Reader(in io.Reader) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(in)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return lines, nil
}

func (app *App) Cut(string *string) (string, error) {
	b := strings.Builder{}

	split := strings.Split(*string, app.config.delimiter)
	if app.config.separated {
		if len(split) == 1 {
			return "", nil
		}
	}

	//if app.columns[len(app.columns)-1] > len(split) {
	//	return "", errors.New("wrong column number")
	//}

	for _, column := range app.columns {
		if column < len(split) {
			b.WriteString(split[column] + app.config.delimiter)
		}

	}

	return strings.TrimPrefix(b.String(), app.config.delimiter), nil
}

func main() {

	myApp := NewApp()
	myApp.Init()
	out := myApp.Scanner()
	fmt.Println(out)

}
