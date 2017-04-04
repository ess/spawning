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

	"github.com/ess/mockable"
)

// Runner is an interface describing an object that knows how to Run a
// collection of commands. In return, the caller gets a collection of Results.
type Runner interface {
	Run([]string) []*Result
}

// Concurrently returns a Runner. If mocking is enabled, it is a mocked Runner.
// Otherwise, it is a Runner that executes commands concurrently.
func Concurrently() Runner {
	if mockable.Mocked() {
		return &mockedRunner{}
	}

	return &concurrentRunner{}
}

// Sequentially returns a Runner. If mocking is enabled, it is a mocked Runner.
// Otherwise, it is a Runner that executes commands one-at-a-time in the order
// they are received.
func Sequentially() Runner {
	if mockable.Mocked() {
		return &mockedRunner{}
	}

	return &sequentialRunner{}
}

func prefixedCommand(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}
