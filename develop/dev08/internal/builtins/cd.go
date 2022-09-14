package builtins

import (
	"errors"
	"os"
	"sync"
	"wb_l2/develop/dev08/internal"
)

type Cd struct {
	internal.Command
}

func (c Cd) Run(wg *sync.WaitGroup) error {
	args := c.GetArgs()
	c.CloseFds()
	defer wg.Done()
	if len(args) != 2 {
		return errors.New("cd: too many arguments")
	}
	if err := os.Chdir(args[1]); err != nil {
		return errors.New("cd: " + err.Error())
	}
	return nil
}
