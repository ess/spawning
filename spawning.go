// Package spawning provides a handy API for spawning shell commands on UNIXish
// environments.
package spawning

// NewPool returns a brand new Pool to use for spawning multiple commands.
func NewPool() *Pool {
	return &Pool{}
}

// Run takes a single string command, spawns the command, and returns the
// Result of that run.
func Run(command string) *Result {
	return NewPool().
		Add(command).
		Run(Sequentially())[0]
}
