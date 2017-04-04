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

// Package spawning provides a handy API for spawning shell commands on UNIXish
// environments.
package spawning

// NewPool returns a brand new Pool to use for spawning multiple commands.
func NewPool() *Pool {
	return &Pool{}
}

// Run takes a single string command, spawns the command, and returns the
// Result of that run.
func Run(command string) *Result {
	return NewPool().
		Add(command).
		Run(Sequentially())[0]
}
