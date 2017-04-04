package spawning

import (
	"os/exec"
)

type sequentialRunner struct{}

func (runner *sequentialRunner) getResult(command *exec.Cmd) *Result {
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

func (runner *sequentialRunner) Run(commands []string) []*Result {
	results := make([]*Result, 0)

	for _, command := range commands {
		results = append(results, runner.getResult(prefixedCommand(command)))
	}

	return results
}
