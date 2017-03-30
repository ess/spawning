package spawning

import (
	"os/exec"
)

type realPool struct {
	commands []*exec.Cmd
}

func (pool *realPool) Add(cmd string) {
	pool.commands = append(pool.commands, exec.Command("bash", "-c", cmd))
}

func (pool *realPool) getResult(command *exec.Cmd, results chan<- *Result) {
	result := &Result{
		Command: command.Args[len(command.Args)-1],
		Success: true,
	}

	output, err := command.CombinedOutput()
	if err != nil {
		result.Success = false
	}

	result.Output = string(output)

	results <- result
}

func (pool *realPool) Run() []*Result {
	results := make(chan *Result, len(pool.commands))

	ret := make([]*Result, 0)

	for _, command := range pool.commands {
		go pool.getResult(command, results)
	}

	for range pool.commands {
		ret = append(ret, <-results)
	}

	return ret
}
