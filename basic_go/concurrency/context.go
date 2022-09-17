package concurrency

import "time"

// source: https://go.dev/blog/context

// Introduction
/*

	In Go servers, each incoming request is handled in its own goroutine.
	Request handlers often start additional goroutines to access backends such as databases
	and RPC services. The set of goroutines working on a request typically needs access to request-specific
	values such as the identity of the end user, authorization tokens, and the request’s deadline.
	When a request is canceled or times out, all the goroutines working on that request should exit
	quickly so the system can reclaim any resources they are using.

	it easy to pass request-scoped values, cancellation signals,
	and deadlines across API boundaries to all the goroutines
	involved in handling a request.
*/

// core context

// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
	// Done returns a channel that is closed when this Context is canceled
	// or times out.
	Done() <-chan struct{}

	// Err indicates why this context was canceled, after the Done channel
	// is closed.
	Err() error

	// Deadline returns the time when this Context will be canceled, if any.
	Deadline() (deadline time.Time, ok bool)

	// Value returns the value associated with key or nil if none.
	Value(key interface{}) interface{}
}
