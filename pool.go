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
