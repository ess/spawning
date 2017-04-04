// Copyright © 2017 Dennis Walters
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
