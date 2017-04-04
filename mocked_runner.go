package spawning

type mockedRunner struct{}

func (runner *mockedRunner) Run(commands []string) []*Result {
	ret := make([]*Result, 0)

	for _, command := range commands {
		result := &Result{
			Command: command,
			Output:  command,
			Success: true,
		}

		ret = append(ret, result)
	}

	return ret
}
