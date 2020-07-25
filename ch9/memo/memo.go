package memo

import "fmt"

// Func is a memoization function
type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	done     <-chan struct{}
	response chan<- result
}

// Memo represents a memoized cache
type Memo struct {
	requests, cancels chan request
}

// New returns a memoiztaion of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{make(chan request), make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, done, response}
	memo.requests <- req
	fmt.Println("get: waiting for response")
	res := <-response
	fmt.Println("get: checking if cancelled")

	select {
	case <-done:
		fmt.Println("get: queueing cancellation request")
		memo.cancels <- req
	default:
		// not cancelled so continue
	}
}
