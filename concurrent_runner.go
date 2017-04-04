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

type concurrentRunner struct{}

func (runner *concurrentRunner) getResult(command *exec.Cmd, results chan<- *Result) {
	result := &Result{
		Command: command.Args[len(command.Args)-1],
		Success: true,
	}

	output, err := command.CombinedOutput()
	if err != nil {
		result.Success = false
	}

	result.Output = string(output)

	results <- result
}

func (runner *concurrentRunner) Run(commands []string) []*Result {
	results := make(chan *Result, len(commands))

	ret := make([]*Result, 0)

	for _, command := range commands {
		go runner.getResult(prefixedCommand(command), results)
	}

	for range commands {
		ret = append(ret, <-results)
	}

	return ret
}
