package builtins

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"wb_l2/develop/dev08/internal"
)

func split(data, delim, ignore string) (result []string) {
	var it1, it2 bool
	var pnt int

	data = strings.Trim(data, delim)
	i := 0
	for ; i < len(data); i++ {
		if strings.ContainsRune(ignore, rune(data[i])) && (i == 0 || it1 || data[i-1] != '\\') {
			it1 = !it1
			it2 = true
		}
		if !it1 {
			if strings.ContainsRune(delim, rune(data[i])) && (i == 0 || it2 || data[i-1] != '\\') {
				result = append(result, data[pnt:i])
				for i+1 < len(data) && strings.ContainsRune(delim, rune(data[i+1])) {
					i++
				}
				pnt = i + 1
			}
			it2 = false
		}
	}
	if i != pnt {
		result = append(result, data[pnt:i])
	}
	return result
}

func CreateCommands(input string, paths []string) (commands []internal.ICommand, err error) {
	const ignore = "\"'"
	var pipex = make([]int, 2)
	var out, in int
	var cmd internal.ICommand

	if err = syscall.Pipe(pipex); err != nil {
		return nil, err
	}

	out = pipex[1]
	groups := split(input, ";", ignore)
	for _, group := range groups {
		in = 0
		pipeSplit := split(group, "|", ignore)
		for _, cmdline := range pipeSplit {
			args := split(cmdline, " ", ignore)
			for i := range args {
				args[i] = strings.TrimFunc(args[i], func(r rune) bool {
					return r == '\'' || r == '"'
				})
			}
			if cmd, err = createCommand(args, paths, out, in); err != nil {
				fmt.Printf("%s\n", err.Error())
				return nil, nil
			}
			commands = append(commands, cmd)
			in = pipex[0]
			if err = syscall.Pipe(pipex); err != nil {
				return nil, err
			}
			out = pipex[1]
		}
		if err = syscall.Close(commands[len(commands)-1].GetWriter()); err != nil {
			log.Fatal(err)
		}
		commands[len(commands)-1].SetWriter(1)
		out = pipex[1]
	}
	return commands, nil
}

func checkFile(ut string) (string, error) {
	stat, err := os.Stat(ut)
	if err != nil {
		return "", err
	} else if stat.IsDir() {
		return "", errors.New(ut + " is directory, can't execute")
	} else if stat.Mode()&0100 == 0 {
		return "", errors.New(ut + " isn't executable, pls make: \n$> chmod +x " + ut)
	}
	return ut, nil
}

func createCommand(args, paths []string, writer, reader int) (internal.ICommand, error) {
	switch args[0] {
	case "cd":
		return &Cd{Command: *internal.NewCommand(args, writer, reader)}, nil
	case "pwd":
		return &Pwd{Command: *internal.NewCommand(args, writer, reader)}, nil
	case "echo":
		return &Echo{Command: *internal.NewCommand(args, writer, reader)}, nil
	case "kill":
		return &Kill{Command: *internal.NewCommand(args, writer, reader)}, nil
	case "ps":
		return &Ps{Command: *internal.NewCommand(args, writer, reader)}, nil
	case "exec":
		return &Exec{Command: *internal.NewCommand(args, writer, reader)}, nil
	}
	for _, v := range paths {
		if _, err := checkFile(v + "/" + args[0]); err == nil {
			args[0] = v + "/" + args[0]
			return internal.NewCommand(args, writer, reader), nil
		}
	}
	return nil, errors.New(args[0] + ": command not found")
}
