package spawning

import (
	"os/exec"

	"github.com/ess/mockable"
)

// Runner is an interface describing an object that knows how to Run a
// collection of commands. In return, the caller gets a collection of Results.
type Runner interface {
	Run([]string) []*Result
}

// Concurrently returns a Runner. If mocking is enabled, it is a mocked Runner.
// Otherwise, it is a Runner that executes commands concurrently.
func Concurrently() Runner {
	if mockable.Mocked() {
		return &mockedRunner{}
	}

	return &concurrentRunner{}
}

// Sequentially returns a Runner. If mocking is enabled, it is a mocked Runner.
// Otherwise, it is a Runner that executes commands one-at-a-time in the order
// they are received.
func Sequentially() Runner {
	if mockable.Mocked() {
		return &mockedRunner{}
	}

	return &sequentialRunner{}
}

func prefixedCommand(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}
