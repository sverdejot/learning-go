package synch

import (
	"testing"
	"sync"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		
		assertValue(t, counter, 3)
	})
	t.Run("incrementing the counter concurrently", func(t *testing.T) {
		counter := NewCounter()
		loops := 1000	
		
		var wg sync.WaitGroup
		wg.Add(loops)

		for i := 0; i < loops; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertValue(t, counter, loops)
	})
}

func assertValue(t testing.TB, counter *Counter, value int) {
	t.Helper()
	if counter.Value() != value {
		t.Errorf("counter value should be %d and its %d", value, counter.value)
	}
}
