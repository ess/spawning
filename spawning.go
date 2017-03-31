package spawning

import (
	"os/exec"

	"github.com/ess/mockable"
)

type Result struct {
	Command string
	Output  string
	Success bool
}

func NewConcurrentPool() Pool {
	if mockable.Mocked() {
		return &mockedPool{}
	}

	return &concurrentPool{}
}

func NewSequentialPool() Pool {
	if mockable.Mocked() {
		return &mockedPool{}
	}

	return &sequentialPool{}
}

func Run(command string) *Result {
	pool := NewSequentialPool()
	pool.Add(command)
	return pool.Run()[0]
}

func prefixedCommand(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}
