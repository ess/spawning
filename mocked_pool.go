package spawning

type mockedPool struct {
	commands []string
}

func (pool *mockedPool) Add(cmd string) Pool {
	pool.commands = append(pool.commands, cmd)

	return pool
}

func (pool *mockedPool) Run() []*Result {
	ret := make([]*Result, 0)

	for _, command := range pool.commands {
		result := &Result{
			Command: command,
			Output:  command,
			Success: true,
		}

		ret = append(ret, result)
	}

	return ret
}
