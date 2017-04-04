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
