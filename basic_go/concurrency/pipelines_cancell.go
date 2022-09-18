package concurrency

import (
	"fmt"
	"sync"
)

// third stage
func Final_state() {
	// explicit_cancellation()
	// explicit_cancellation_improve()
	digesting_tree(true)

}

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
func gen(with_buffer bool, nums ...int) chan int {
	if with_buffer {
		return gen_with_buffer(nums...)
	}
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// add buffer, avoid resource leak
/*
	A buffer can hold a fixed number of values; send operations complete immediately if there’s room in the buffer:
*/
func gen_with_buffer(nums ...int) chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func gen_with_done(done chan struct{}, nums ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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

func sq_with_done(done chan struct{}, in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out

}

func merge_with_done(done chan struct{}, cs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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

func merge(with_buffer bool, cs ...chan int) chan int {
	/*
		The "merge" function converts a list of channels to a single channel by starting a goroutine
		for each inbound channel that copies the values to the sole outbound channel.

		Once all the "output" goroutines have been started,
		"merge" starts one more goroutine to close the outbound channel
		after all sends on that channel are done.
	*/
	var wg sync.WaitGroup
	out := make(chan int)
	if with_buffer {
		// enough space for the unread inputs
		out = make(chan int, 1)
	}

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

func explicit_cancellation_improve() {
	/*
		Set up a done channel that's shared by the whole pipeline,
		and close that channel when this pipeline exits, as a signal
		for all the goroutines we started to exit

		Here are the guidelines for pipeline construction:
		- stages close their outbound channels when all the send operations are done
		- stages keep receiving values from inbound channels until those channels are closed or the
		senders are unblocked

		Pipelines unblock senders either by ensuring there’s enough buffer for all the values
		that are sent or by explicitly signalling senders when the receiver may abandon
		the channel.
	*/
	done := make(chan struct{})
	defer close(done)

	in := gen_with_done(done, 2, 3)
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq_with_done(done, in)
	c2 := sq_with_done(done, in)

	// Consume the first value from output.
	out := merge_with_done(done, c1, c2)
	fmt.Println(<-out)

}

func explicit_cancellation() {
	fmt.Println("Explicit cancellation")
	/*
		When main decides to exit without receiving all the values from out,
		it must tell the goroutines in the upstream stages to abandon the values
		they’re trying to send.

		This approach has a problem: each downstream receiver needs to know the number of potentially
		blocked upstream senders and arrange to signal those senders on early return.
		Keeping track of these counts is tedious and error-prone.
	*/

	in := gen(false, 2, 3)
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output
	done := make(chan struct{}, 2)
	out := merge_with_done(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}
}

func square_number() {
	fmt.Println("Squaring numbers")
	// convert to channel
	c := gen(false, 2, 3)
	// From input channel extracting number then mutiplying by 2 next assigning to new channel
	// which will finally be returned.
	out := sq(c)

	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	// we can compose inbound and outbound channels, because they are same type,
	for n := range sq(gen(false, 2, 3)) {
		fmt.Println(n)
	}
	fmt.Println()

}

func fan_in_out() {
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
	fmt.Println("fan-out, fan-in")
	in := gen(false, 2, 3, 7)
	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(false, c1, c2) {
		fmt.Println(n)
	}
	fmt.Println()

}

func stopping_short(resolve bool) {
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
	if !resolve {

		fmt.Println("Stopping short")
		in := gen(false, 3, 4)
		c1 := sq(in)
		c2 := sq(in)
		// Consume the first value from output.
		out := merge(false, c1, c2)
		fmt.Println(<-out)
		return
		// Since we didn't receive the second value from out,
		// one of the output goroutines is hung attemping to send it
		/*
			This is a resource leak: goroutines consume memory and runtime resources, and
			heap references in goroutine stacks keep data from being garbage collected.
			Goroutines are not garbage collected; they must exit on their own
		*/
	}

	/*
		While this fixes the blocked goroutine in this program, this is bad code. The choice
		of buffer size of 1 here depends on knowing the number of values merge will receive
		and the number of values downstream stages will consume. This is fragile: if we pass
		an additional value to gen, or if the downstream stage reads any fewer values,
		we will again have blocked goroutines.
	*/
	fmt.Println("\neasy way resolve stopping short, but have cons")
	in := gen(true, 2, 3)
	c1 := sq(in)
	c2 := sq(in)
	for n := range merge(true, c1, c2) {
		fmt.Println(n)
	}
	return
}
