package builtins

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"wb_l2/develop/dev08/internal"
)

type Kill struct {
	internal.Command
}

func (k *Kill) Run(wg *sync.WaitGroup) error {
	var (
		args []string
		proc *os.Process
	)
	k.CloseFds()
	defer wg.Done()
	args = k.GetArgs()
	if len(args) > 1 {
		for _, v := range args[1:] {
			if pid, err := strconv.Atoi(v); err != nil {
				fmt.Println("kill: pid:", v, "is not valid")
			} else {
				if proc, err = os.FindProcess(pid); err != nil {
					fmt.Println("kill:", pid, "not found")
				} else if err = proc.Kill(); err != nil {
					fmt.Println("kill: " + err.Error())
				}
			}
		}
	}

	return nil
}
