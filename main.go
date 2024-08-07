package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	basePath string = "C://"
	errLog   *log.Logger
	infoLog  *log.Logger
)

/*Инициализация логгеров*/
func initLoggers() {
	logFile, err := os.OpenFile("MYTERMINALLOG.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	errLog = log.New(logFile, "ERROR ", log.Ldate|log.Ltime|log.Llongfile)
	infoLog = log.New(logFile, "INFO ", log.Ldate|log.Ltime)
}

/*Функция для запуска команды*/
func cmd(args string) {
	args = strings.TrimSpace(args) //Удаление пробелов
	if args == "" {
		infoLog.Println("Empty command")
		return
	}

	/*команда help - помощь*/
	if args == "help" {
		fmt.Println("Самодельный терминал))")
		return
	}

	newArgs := strings.Split(args, " ") //Разделение команды на слова

	/*Если команда cd*/
	if newArgs[0] == "cd" {
		infoLog.Printf("CD basePath from %v to %v", basePath, newArgs[1])
		basePath = newArgs[1]
		os.Chdir(basePath)
		return
	}

	infoLog.Printf("Full command: %v", newArgs)
	cmd := exec.Command("powershell", newArgs...)
	/*Установка Stdin, Stdout, Stderr для команды*/
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	fmt.Print("<<< - - - - - - - - - - - - - -\n") //Так красивее))
	err := cmd.Run()
	if err != nil {
		errLog.Printf("Error in running: %v", err)
	}
	fmt.Print("<<< - - - - - - - - - - - - - -\n") //Так красивее))
}

func main() {
	initLoggers()
	sc := bufio.NewScanner(os.Stdin)

	/*Сканирование команд в бесконечном цикле*/
	for {
		fmt.Printf("\n%s >>>", basePath)
		sc.Scan()
		cmd(sc.Text())
	}
}
