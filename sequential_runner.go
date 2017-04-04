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
	"os/exec"
)

type sequentialRunner struct{}

func (runner *sequentialRunner) getResult(command *exec.Cmd) *Result {
	result := &Result{
		Command: command.Args[len(command.Args)-1],
		Success: true,
	}

	output, err := command.CombinedOutput()
	if err != nil {
		result.Success = false
	}

	result.Output = string(output)

	return result
}

func (runner *sequentialRunner) Run(commands []string) []*Result {
	results := make([]*Result, 0)

	for _, command := range commands {
		results = append(results, runner.getResult(prefixedCommand(command)))
	}

	return results
}
