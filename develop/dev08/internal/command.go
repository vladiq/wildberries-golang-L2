package internal

import (
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
)

type ICommand interface {
	Run(*sync.WaitGroup) error
	GetWriter() int
	SetWriter(int)
	GetReader() int
	SetReader(int)
	GetPid() uintptr
}

func Execute(executable []ICommand) {
	var wg sync.WaitGroup

	for _, e := range executable {
		wg.Add(1)
		if err := e.Run(&wg); err != nil {
			if err = fmt.Errorf("%v", err); err != nil {
				log.Fatal(err)
			}
		} else if e.GetPid() > 0 {
			if _, err = syscall.Wait4(int(e.GetPid()), nil, 0, nil); err != nil {
				log.Fatal(err)
			}
		}
		wg.Wait()
	}
}

type Command struct {
	args   []string
	writer int
	reader int
	pid    uintptr
}

func (c *Command) GetPid() uintptr {
	return c.pid
}

func (c *Command) SetPid(pid uintptr) {
	c.pid = pid
}

func (c *Command) GetArgs() []string {
	return c.args
}

func (c *Command) GetWriter() int {
	return c.writer
}

func (c *Command) SetWriter(writer int) {
	c.writer = writer
}

func (c *Command) GetReader() int {
	return c.reader
}

func (c *Command) SetReader(reader int) {
	c.reader = reader
}

func NewCommand(args []string, writer, reader int) *Command {
	return &Command{args, writer, reader, 0}
}

func (c *Command) DupAll() (err error) {
	if err = syscall.Dup2(c.writer, 1); err != nil {
		return err
	}
	if err = syscall.Dup2(c.reader, 0); err != nil {
		return err
	}
	c.CloseFds()
	return nil
}

func (c *Command) CloseFds() {
	var err error
	if c.writer != 1 {
		if err = syscall.Close(c.writer); err != nil {
			log.Fatal(err)
		}
	}
	if c.reader != 0 {
		if err = syscall.Close(c.reader); err != nil {
			log.Fatal(err)
		}
	}
}

func (c *Command) Fork() (pid uintptr) {
	pid, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	return pid
}

func (c *Command) Run(wg *sync.WaitGroup) (err error) {
	pid := c.Fork()
	if pid == 0 {
		if err = c.DupAll(); err != nil {
			log.Fatal(err)
		}
		if err = syscall.Exec(c.args[0], c.args, os.Environ()); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	c.CloseFds()
	c.SetPid(pid)
	wg.Done()
	return nil
}
