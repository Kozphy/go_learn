package concurrency

import (
	"context"
	"fmt"
	"time"
)

// source: https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

// Createing a Context and Using data within a context
/*
	When using contexts, it’s important to know that the values stored in a specific context.Context
	are immutable, meaning they can’t be changed.

	When you called the context.WithValue, you passed in the parent context and you
	also received a context back. You received a context back because the context.WithValue
	function didn’t modify the context you provided. Instead, it wrapped your parent context
	inside another one with the new value.

	context.TODO(): you use this when you’re not sure which context to use.

	context.Background: The context.Background function creates an empty context like context.TODO does,
	but it’s designed to be used where you intend to start a known context.

	ctx.Done(): This method returns a channel that is closed when the context is done,
	and any functions watching for it to be closed will know they should consider their
	execution context completed and should stop any processing related to the context.

*/

/*
	Note:
	Contexts can be a powerful tool with all the values they can hold, but
	a balance needs to be struck between data being stored in a context (hard to read and maintain)
	and data being passed to a function as parameters.

	A good rule of thumb is that any data required for a function to run should be passed as parameters
*/

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

func Execute1_digitoc() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}

// Determining if a Context is Done
/*
	For the code execution in this example, the for loop will continue forever until the
	ctx.Done channel is closed because the only return statement is inside that case statement.
	Even though the case <- ctx.Done doesn’t assign values to any variables, it will still
	be triggered when ctx.Done is closed because the channel still has a value that can be read
	even if it’s ignored. If the ctx.Done channel isn’t closed, the select statement will wait until it is,
	or if resultsCh has a value that can be read. If resultsCh can be read, then that case statement’s code
	block will be executed instead.
*/

// func Execute2_digitoc() {
// 	ctx := context.Background()
// 	resultsCh := make(chan *WorkResult)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// The context is over, stop processing results
// 			return
// 		case result := <-resultsCh:
// 			// Process the results received
// 		}
// 	}
// }

// Cancelling a Context
/*
	Note:
	If you’ve run Go programs before and looked at the logging output, you may have
	seen the context canceled error in the past. When using the Go http package, this
	is a common error to see when a client disconnects from the server before the
	server handles the full response.

*/

func doSomething_cancel(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doAnother3(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	cancelCtx()

	/*
		In many cases the time.Sleep would not be required, but it’s needed since
		the example code finishes executing so quickly. If time.Sleep isn’t included,
		the program may end before you see the rest of the program’s output on the screen.
	*/
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")

}

func doAnother3(ctx context.Context, printCh chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func Execute_cancel_digitoc() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething_cancel(ctx)
}

// Giving a context a deadline
/*
	Using context.WithDeadline with a context allows you to set a deadline for
	when the context needs to be finished, and it will automatically end when that
	deadline passes.

	The context can also be canceled manually by calling the cancel function the
	same as you would for context.WithCancel.

	When a context is canceled from a deadline, the cancel function is still required
	to be called in order to clean up any resources that were used, so this is more of
	a safety measure.

	If the deadline ends up at or over 3 seconds, you could even see the error change
	back to the context canceled error because the deadline is no longer being exceeded.


*/
func doSomething_deadline(ctx context.Context) {
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	// not necessarily required
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother3(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("doSomething: finished\n")
}
func Execute_deadline_digitoc() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething_deadline(ctx)
}

// Giving a Context a Time Limit
/*
	by using the context.WithTimeout function you only need to provide a time.Duration
	value for how long you want the context to last.

*/

func doSomething_timeout(ctx context.Context) {
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancelCtx()
	printCh := make(chan int)
	go doAnother3(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("doSomething: finished\n")
}

func Execute_timeout_digitoc() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething_timeout(ctx)
}
