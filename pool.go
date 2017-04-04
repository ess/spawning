package spawning

// Pool is an interface that describes a process spawn pool.
//type Pool interface {
//Add(string) Pool
//Run() []*Result
//}

// Pool is a collection of commands to spawn.
type Pool struct {
	commands []string
}

// Add appends the provided command to the associated pool's command collection.
// It returns the pool itself, which means that it is a chainable method.
func (pool *Pool) Add(command string) *Pool {
	pool.commands = append(pool.commands, command)

	return pool
}

// Run executes the pool's commands via the provided Runner.
func (pool *Pool) Run(executionModel Runner) []*Result {
	return executionModel.Run(pool.commands)
}
