package spawning

import (
	"fmt"
)

// Result represents the actual result of a spawned command
type Result struct {
	Command string
	Output  string
	Success bool
}

func (result *Result) String() string {
	var status string

	if result.Success {
		status = "success"
	} else {
		status = "failure"
	}

	return fmt.Sprintf("%s : %s", result.Command, status)
}
