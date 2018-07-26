package main

import (
	"github.com/ghjan/go-designpattern/concurency/mutex/counter"
	"time"
)

func main() {
	counter := counter.Counter{}
	for i := 0; i < 10; i++ {
		go func(index int) {
			counter.Lock()
			counter.Value++
			defer counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)

	counter.Lock()
	defer counter.Unlock()
	println(counter.Value)

}
