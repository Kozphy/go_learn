package concurrency

import (
	"fmt"
	"sync"
)

// source: https://go.dev/blog/pipelines

// Introduce
/*
	Go’s concurrency primitives make it easy to construct streaming data pipelines that
	make efficient use of I/O and multiple CPUs. This article presents examples of such
	pipelines, highlights subtleties that arise when operations fail, and introduces
	techniques for dealing with failures cleanly.
*/

// What is a pipeline?
/*
	Informally, a pipeline is a series of stages connected by channels,
	where each stage is a group of goroutines running the same function. In each stage,
	the goroutines:
	- receive value from upstream via inbound channels
	- perform some function on that data, usually producing new values
	- send values downstream via outbound channels

	Each stage has any number of inbound and outbound channels,
	except the first and last stages, which have only outbound or inbound channels,
	respectively. The first stage is sometimes called the "source" or "producer"; the last stage,
	the "sink" or "consumer".
*/

// Squaring numbers
// First stage
/*
	The first stage, gen, is a function that
	converts a list of integers to a channel that emits the integers in the list.
*/
func gen(nums ...int) chan int {
	// add buffer, avoid resource leak
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// second stage
/*
	sq, receives integers from a channel and returns a channel that
	emits the square of each received integer.
*/
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	/*
		The "merge" function converts a list of channels to a single channel by starting a goroutine
		for each inbound channel that copies the values to the sole outbound channel.

		Once all the "output" goroutines have been started,
		"merge" starts one more goroutine to close the outbound channel
		after all sends on that channel are done.
	*/
	var wg sync.WaitGroup
	out := make(chan int)

	// Start on output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed, then calls wg.Done.
	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines
	// are done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// third stage
func Final_state() {
	// convert to channel
	c := gen(2, 3)
	// From input channel extracting number then mutiplying by 2 next assigning to new channel
	// which will finally be returned.
	out := sq(c)

	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	// we can compose inbound and outbound channels, because they are same type,
	for n := range sq(gen(2, 3)) {
		fmt.Println(n)
	}

	// Fan-out, fan-in
	/*
		Multiple functions can read from the same channel until that channel is closed;
		this is called fan-out. This provides a way to distribute work
		amongst a group of workers to parallelize CPU use and I/O.

		A function can read from multiple inputs and proceed until all are closed by multiplexing
		the input channels onto a single channel that’s closed when all the inputs are closed.
		This is called fan-in.

		We can change our pipeline to run two instances of sq,
		each reading from the same input channel.
		We introduce a new function, "merge", to fan in the results:
	*/
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

	// Stopping short
	/*
		There is a pattern to our pipeline functions:
		- stages close their outbound channels when all the send operations are done.
		- stages keep receiving values from inbound channels until those channels are closed.

		This pattern allows each receiving stage to be written as a range loop and
		ensures that all goroutines exit once all values have been successfully sent downstream.

		But in real pipelines, stages don’t always receive all the inbound values.
		More often, a stage exits early because an inbound value represents an error in an earlier stage.

		In either case the receiver should not have to wait for the remaining values to arrive, and
		we want earlier stages to stop producing values that later stages don't need.

	*/

}
