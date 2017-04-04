package spawning

import (
	"os/exec"
)

type concurrentRunner struct{}

func (runner *concurrentRunner) getResult(command *exec.Cmd, results chan<- *Result) {
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

func (runner *concurrentRunner) Run(commands []string) []*Result {
	results := make(chan *Result, len(commands))

	ret := make([]*Result, 0)

	for _, command := range commands {
		go runner.getResult(prefixedCommand(command), results)
	}

	for range commands {
		ret = append(ret, <-results)
	}

	return ret
}
