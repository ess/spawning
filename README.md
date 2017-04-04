The `spawning` package provides a handy API for spawning shell commands on
UNIXish environments.

## Installation ##

I suggest the use of [glide](https://glide.sh) for managing your Go deps, but
you should be able to install it directly without much issue:

```
go get github.com/ess/spawning
```

As mentioned, though, a better idea is to use glide (or another package manager
that supports SemVer).

## Usage ##

Commands (run via `bash -c`) can be spawned in a one-off manner, or one can
add a collection of commands to a Pool, then run the pool either sequentially,
concurrently, or with a custom Runner.

### Spawning One-off Commands ###

Running a one-off command is pretty easy, really:

```go
package main

import (
	"fmt"

	"github.com/ess/spawning"
)

func main() {
	result := spawning.Run("echo 'My sausages turned to gold!'")

	if result.Success {
		fmt.Println("The sausages were successfully turned to gold.")
	}
}
```

### Spawning Multiple Commands ###

Running multiple commands is fairly easy, too, and you can execute them either
Sequentially, Concurrently, or with a custom Runner implementation.

#### Sequentially ####

Effectively, the sequential runner could be used to create minimal shell
scripts. The important caveat there is that as every command is executed via
`bash -c`, there is no shared environment between the commands in a given Pool.

```go
package main

import (
	"fmt"
	"github.com/ess/spawning"
)

func main() {
	script := spawning.NewPool()
	script.Add("mkdir -p output")
	script.Add("date > output/begin")
	script.Add(`echo "The operation has succeeded" > output/result`)
	script.Add("date > output/end")
	script.Add(`for step in begin result end ; do cat output/${step} ; done`)

	// Run the commands sequentially and process the results
	for _, result := range script.Run(spawning.Sequentially()) {
		processResult(result)
	}
}

func processResult(result *spawning.Result) {
	if !result.Success {
		fmt.Println("The following command failed:", result.Command)
		return
	}

	fmt.Println(result.Output)
}
```

#### Concurrently ####

Running commands concurrently isn't at all great for running locally, but it is
quite handy for, say, controlling a cluster via SSH.

```go
package main

import (
	"fmt"
	"github.com/ess/spawning"
)

func main() {

	logins := []string{
		"joe@192.168.1.1",
		"joe@192.168.1.2",
		"joe@192.168.1.3",
		"jim@10.1.1.1",
	}

	pool := spawning.NewPool()

	for _, login := range logins {
		pool.Add(sshCommand(login, "date"))
	}

	// Run the commands concurrently and process the results
	for _, result := range pool.Run(spawning.Concurrently()) {
		processResult(result)
	}
}

func processResult(result *spawning.Result) {
	if !result.Success {
		fmt.Println("The following command failed:", result.Command)
		return
	}

	fmt.Println(result.Output)
}

func sshCommand(login string, command string) string {
	return fmt.Sprintf("ssh %s '$s'", login, command)
}
```

#### Running With A Custom Runner ####

`spawning.Pool`'s `Run()` method requires that one pass in a `Runner`. If
neither `Sequentially` nor `Concurrently` are suitable for your purposes, you
can implement your own Runner:

```go
package main

import (
	"fmt"
	"github.com/ess/spawning"
)

type RoadRunner struct{}

func (runner *RoadRunner) Run(commands []string) []*spawning.Result {
	results := make([]*spawning.Result, 0)

	for _, command := range commands {
		results = append(
			results,
			&spawning.Result{Command: command, Output: "meep meep", Success: true},
		)
	}

	return results
}

func main() {
  runner := &RoadRunner{}

	for _, result := range spawning.NewPool().Add("sudo ls -lah /").Run(runner) {
		processResult(result)
	}
}

func processResult(result *spawning.Result) {
	if !result.Success {
		fmt.Println("The following command failed:", result.Command)
		return
	}

	fmt.Println(result.Output)
}
```

## History ##

* v0.1.1 - I'm a jerk that forgets to include a license
* v0.1.0 - Initial release
