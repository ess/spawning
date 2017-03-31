package spawning

import (
	"os/exec"
)

type sequentialPool struct {
	commands []*exec.Cmd
}

func (pool *sequentialPool) Add(cmd string) Pool {
	pool.commands = append(pool.commands, prefixedCommand(cmd))

	return pool
}

func (pool *sequentialPool) getResult(command *exec.Cmd) *Result {
	result := &Result{
		Command: command.Args[len(command.Args)-1],
		Success: true,
	}

	output, err := command.CombinedOutput()
	if err != nil {
		result.Success = false
	}

	result.Output = string(output)

	return result
}

func (pool *sequentialPool) Run() []*Result {
	results := make([]*Result, 0)

	for _, command := range pool.commands {
		results = append(results, pool.getResult(command))
	}

	return results
}
