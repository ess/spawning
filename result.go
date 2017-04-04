package spawning

import (
	"fmt"
)

// Result represents the actual result of a spawned command. The Command member
// contains the actual command that was run. The Output member is the combined
// STDOUT and STDERR of the command. The Success member reflects whether or not
// the command had a successful exit status.
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
