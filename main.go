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
	basePath string = "."
	errLog   *log.Logger
	infoLog  *log.Logger
)

/*Инициальзация логгеров*/
func initLoggers() {
	logFileInfo, err := os.OpenFile("INFO.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logFileError, err := os.OpenFile("ERROR.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	errLog = log.New(logFileError, "ERROR ", log.Ldate|log.Ltime|log.Llongfile)
	infoLog = log.New(logFileInfo, "INFO ", log.Ldate|log.Ltime)
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
		fmt.Println("Самодельный терминал\nПеред путём к файлу необходимо ставить знак ':'")
		return
	}

	newArgs := strings.Split(args, " ") //Разделение команды на слова

	/*Если команда cd*/
	if newArgs[0] == "cd" {
		infoLog.Printf("CD basePath from %v to %v", basePath, basePath+"/"+newArgs[1])
		basePath = basePath + "/" + newArgs[1]
		return
	}

	/*Замена всех путей на полный путь от программы*/
	for i, v := range newArgs {
		if []rune(v)[0] == ':' {
			newV := []rune(v)[1:]
			newVstr := basePath + "/" + string(newV)
			infoLog.Printf("Command path from %v to %v", newV, newVstr)
			newArgs[i] = newVstr
		}
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
