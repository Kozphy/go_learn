package concurrency

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func All_concurrency() {
	fmt.Println("Concurrency=========================")
	fmt.Println("Goroutines")
	/*
		A goroutine is a lightweight thread managed by the Go runtime.
		go f(x, y, z)

		The evaluation of f, x, y, and z happens in the current goroutine and
		the execution of f happens in the new goroutine.

		Goroutines run in the same address space, so access to shared memory must be synchronized.
	*/
	go say("world")
	say("hello")

	fmt.Println()
	fmt.Println("Channels")
	/*
		Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	*/
	sc := []int{7, 2, 8, -9, 4, 0}
	// Like maps and slices, channels must be created before use:
	ch := make(chan int)
	go sum(sc[:len(sc)/2], ch)
	go sum(sc[len(sc)/2:], ch)
	// Receive from channel, and assign value to x, y
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	fmt.Println()
	fmt.Println("Buffered Channels")
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	// overfill, deadlock
	// ch2 <- 3
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	// ok
	ch2 <- 3
	fmt.Println()
	fmt.Println("Range and Close")
	/*
		A sender can close a channel to indicate that no more values will be sent.

		Using ok-syntax can test whether a channel has been closed.
		v, ok := <- ch
		if there are no more values to receive and the channel is closed, ok is "false".

		Note:  Channels aren't like files; you don't usually need to close them.
		Closing is only necessary when the receiver must be told there are no more values coming
		, such as to terminate a "range" loop.
	*/
	c_f := make(chan int, 10)
	go fibonacci(cap(c_f), c_f)
	// The loop receives values from channel repeatedly until it is closed.
	for i := range c_f {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("Select")
	/*
		The select statement lets a goroutine wait on multiple communication operations.

		A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	*/
	c_s := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c_s)
		}
		quit <- 0
	}()
	fibonacci2(c_s, quit)

	fmt.Println()
	fmt.Println("Default Selection")
	default_selection()

	fmt.Println()

}

func default_selection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
