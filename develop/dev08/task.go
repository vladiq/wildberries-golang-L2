package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"wb_l2/develop/dev08/internal"
	"wb_l2/develop/dev08/internal/builtins"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качестве аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	ReadCmd()
}

func ReadCmd() {
	var (
		commands []internal.ICommand
		paths    []string
	)
	env := os.Getenv("PATH")
	if env != "" {
		paths = strings.Split(env, ":")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		if line, _, err := reader.ReadLine(); err != nil {
			log.Fatal(err)
		} else if string(line) == "\\quit" {
			break
		} else if commands, err = builtins.CreateCommands(string(line), paths); err != nil {
			log.Fatal(err)
		}
		internal.Execute(commands)
	}
}
