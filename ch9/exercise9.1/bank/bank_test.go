package bank

import (
	"sync"
	"testing"
)

func TestWithDraw(t *testing.T) {
	Deposit(50005000)

	var wg sync.WaitGroup
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func(amount int) {
			Withdraw(amount)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if got, want := Balance(), 0; got != want {
		t.Errorf("got: %v, but want: %v", got, want)
	}
}
