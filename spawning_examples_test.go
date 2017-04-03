package spawning_test

import (
	"fmt"

	"github.com/ess/mockable"
	"github.com/ess/spawning"
)

func ExampleRun() {
	result := spawning.Run(`echo "my sausages turned to gold!"`)

	fmt.Println(result.Output)

	// Output:
	// my sausages turned to gold!
	//
}

func ExampleConcurrently() {
	pool := spawning.NewPool()
	pool.Add(`echo "I am slow, so you see me last" ; sleep 0.2`)
	pool.Add(`echo "I am fast, so you see me first" ; sleep 0.1`)
	pool.Add("false")

	for _, result := range pool.Run(spawning.Concurrently()) {
		fmt.Println(result)
	}

	// Output:
	// false : failure
	// echo "I am fast, so you see me first" ; sleep 0.1 : success
	// echo "I am slow, so you see me last" ; sleep 0.2 : success
}

func ExampleSequentially() {
	pool := spawning.NewPool()
	pool.Add(`echo "I am first, so you see me first" ; sleep 0.2`)
	pool.Add(`echo "I am last, so you see me last"`)
	pool.Add("false")

	for _, result := range pool.Run(spawning.Sequentially()) {
		fmt.Println(result)
	}

	// Output:
	// echo "I am first, so you see me first" ; sleep 0.2 : success
	// echo "I am last, so you see me last" : success
	// false : failure
}

func ExamplePool_Run() {
	var allResults []*spawning.Result

	pool := spawning.NewPool().
		Add("sleep 0.1 ; true").
		Add("sleep 0.2 ; false")

	// Run the commands sequentially
	allResults = append(allResults, pool.Run(spawning.Sequentially())...)

	// Run the commands concurrently
	allResults = append(allResults, pool.Run(spawning.Concurrently())...)

	// Just for giggles, verify the mocked runner
	mockable.Enable()
	allResults = append(allResults, pool.Run(spawning.Sequentially())...)
	allResults = append(allResults, pool.Run(spawning.Concurrently())...)
	mockable.Disable()

	for _, result := range allResults {
		fmt.Println(result)
	}

	// Output:
	// sleep 0.1 ; true : success
	// sleep 0.2 ; false : failure
	// sleep 0.1 ; true : success
	// sleep 0.2 ; false : failure
	// sleep 0.1 ; true : success
	// sleep 0.2 ; false : success
	// sleep 0.1 ; true : success
	// sleep 0.2 ; false : success
}
