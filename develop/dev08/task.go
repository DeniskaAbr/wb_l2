package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type App struct {
	path string
}

func NewApp(logger *log.Logger) *App {
	app := App{}

	return &app
}

type Command struct {
	name string
	args []string
}

type prcs struct {
	name string
	*os.Process
}

func (c Command) Execute() (string, error) {
	switch c.name {
	case "cd":
		return "", cd(c.args)
	case "pwd":
		return pwd()
	case "echo":
		return echo(c.args), nil
	case "kill":
		return "", kill(c.args)
	case "ps":
		return ps(), nil
	case "exec":
		return "", execute(c.args)
	case "exit":
		os.Exit(0)
		return "", nil
	default:
		return "", errors.New(c.name + ": command not found")
	}
}

func (a *App) Init() {
	a.UpdatePath()
}

func (a *App) Scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		a.UpdatePath()
		fmt.Print(a.path)

		if scanner.Scan() {
			token := scanner.Text()
			pipe, err := a.ReadInput(token)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			out, err := a.ExecutePipeline(pipe)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if out != "" {
				fmt.Println(out)
			}
		} else {
			fmt.Println()
			break
		}

	}
}

func (a *App) ReadInput(token string) (pipe []Command, err error) {
	commands := strings.FieldsFunc(token, func(r rune) bool {
		return r == '|'
	})
	if len(commands) == 0 {
		return nil, errors.New("empty command")
	}
	for _, p := range commands {
		cmdString := strings.Fields(p)
		switch len(cmdString) {
		case 0:
			return nil, errors.New("command not recognized")
		case 1:
			pipe = append(pipe, Command{
				name: cmdString[0],
			})
		default:
			pipe = append(pipe, Command{
				name: cmdString[0],
				args: cmdString[1:],
			})
		}
	}
	return pipe, nil
}

func (a *App) UpdatePath() {
	path, err := os.Getwd()
	if err != nil {
		a.path = "$ "
	}

	home, err := os.UserHomeDir()
	if err != nil {
		a.path = "$ "
	}
	if strings.HasPrefix(path, home) {
		a.path = "~" + path[len(home):] + " $"
	}
	a.path = path + " $ "
}

func (a *App) ExecutePipeline(pipe []Command) (out string, err error) {
	for _, c := range pipe {

		if out != "" {
			c.args = append(c.args, out)
		}
		out, err = c.Execute()
		if err != nil {
			return "", err
		}
	}
	return out, nil
}

func ReadProcesses() {
	processes, _ = process.Processes()
}

// ***
// Commands

func cd(args []string) (err error) {
	var path string
	switch {
	case len(args) == 0:
		path, err = os.UserHomeDir()
		if err != nil {
			return err
		}
	case strings.HasPrefix(args[0], "~"):
		path, err = os.UserHomeDir()
		if err != nil {
			return err
		}
		path = path + args[0][len("~"):]
	default:
		path = args[0]
	}

	return os.Chdir(path)
}

func pwd() (string, error) {
	return os.Getwd()
}

func echo(args []string) string {
	return strings.Join(args, " ")
}

func kill(args []string) error {
	ReadProcesses()
	if len(args) == 0 {
		return errors.New("process not found")
	}

	var pss []int32

	for _, arg := range args {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		for _, p := range processes {
			if p.Pid == int32(pid) {
				pss = append(pss, p.Pid)
			}
			// return errors.New("Process " + args[i] + " not found")

		}
	}

	for _, pid := range pss {
		for _, p := range processes {
			if p.Pid == pid {
				err := p.Kill()
				if err != nil {
					return err
				}
				return errors.New("Send kill to PID: " + strconv.Itoa(int(p.Pid)) + " no errors")
			}
		}
	}
	//

	return nil
}

func ps() string {
	psString := strings.Builder{}
	defer psString.Reset()
	ReadProcesses()

	for _, process := range processes {
		pid := process.Pid
		name, _ := process.Name()
		s := fmt.Sprintf("%d\t%s\n", pid, name)
		psString.WriteString(s)
	}

	return psString.String()
}

func execute(args []string) error {
	if len(args) == 0 {
		return errors.New("команда не указана")
	}
	cm := exec.Command(args[0], args[1:]...)
	cm.Stdin = os.Stdin
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	return cm.Run()
}

// ***

var processes []*process.Process

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// logger.SetOutput(io.Discard)
	myapp := NewApp(logger)
	myapp.Scan()
}
