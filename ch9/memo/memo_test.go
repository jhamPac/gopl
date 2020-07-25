package memo

import (
	"fmt"
	"sync"
	"testing"
)

func TestCancel(t *testing.T) {
	cancelled := fmt.Errorf("cancelled")

	finish := make(chan string)
	f := func(key string, done <-chan struct{}) (interface{}, error) {
		if done == nil {
			res := <-finish
			return res, nil
		}
		<-done
		return nil, cancelled
	}

	m := New(Func(f))
	key := "key"

	done := make(chan struct{})
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		
	}
}
