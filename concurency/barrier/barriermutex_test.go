package barrier

import (
	"testing"
	"fmt"
	"sync"
)

func TestNewBarrier(t *testing.T) {
	fmt.Println("Hello World!")
	barrier := NewBarrier(3)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("A")
			barrier.BarrierWait()
			fmt.Println("B")
			barrier.BarrierWait()
			fmt.Println("C")
		}()
	}

	wg.Wait()
}
