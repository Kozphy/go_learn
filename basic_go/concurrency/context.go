package concurrency

import (
	"time"
)

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

// Core context
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

/*
	The Done method returns a channel that acts as a cancellation signal to functions
	running on behalf of the Context: when the channel is closed, the functions should
	abandon their work and return.

	A Context does not have a "Cancel" method for the same reason the "Done" channel is
	receive-only: the function receiving a cancellation signal is usually not the
	one that sends the signal.

	The Deadline method allows functions to determine whether they should start work
	at all; if too little time is left, it may not be worthwhile. Code may also use
	a deadline to set timeouts for I/O operations.

	Value allows a Context to carry request-scoped data. That data must be safe for
	simultaneous use by multiple goroutines.
*/

// Derived contexts
/*	func Background() Context
	Background returns an empty Context. It is never canceled, has no deadline,
	and has no values. Background is typically used in main, init, and tests,
	and as the top-level Context for incoming requests.
*/

/*	WithCancel and WithTimeout
	WithCancel is also useful for canceling redundant requests when using multiple replicas.
	WithTimeout is useful for setting a deadline on requests to backend servers:

	func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
*/

// WithValue provides a way to associate request-scoped values with a Context:
/*
	WithValue returns a copy of parent whose Value method returns val for key.

	func WithValue(parent Context, key interface{}, val interface{}) Context
*/

// Example: Google web search
/*
	server that handles URLs like /search?q=golang&timeout=1s by forwarding
	the query “golang” to the Google Web Search API and rendering the results.

	The code is split across three packages:
	- server provides the main function and the handler for /search.
	- userip provides functions for extracting a user IP address from a request and associating it with a Context.
	- google provides the Search function for sending a query to Google.
*/
