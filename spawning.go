package spawning

import (
	"github.com/ess/mockable"
)

type Result struct {
	Command string
	Output  string
	Success bool
}

func NewPool() Pool {
	if mockable.Mocked() {
		return &mockedPool{}
	}

	return &realPool{}
}

func Run(command string) *Result {
	pool := NewPool()
	pool.Add(command)
	return pool.Run()[0]
}
