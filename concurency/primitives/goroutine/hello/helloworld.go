package main

import (
	"time"
	"sync"
	"fmt"
	"strings"
)

func main() {
	helloWorld1()

	helloworld2()

	helloGoroutines()

	demoUpperSync()
	demoUpperASync()
}

func helloWorld1() {
	go func() {
		println("Hello World!")
	}()
	time.Sleep(time.Second)
}

func helloworld2() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		println("Hello World!")
		wait.Done()
	}()

	wait.Wait()
}

func helloGoroutines() {
	var wait sync.WaitGroup
	count := 5
	wait.Add(count)
	for i := 0; i < count; i++ {
		go func(index int) {
			fmt.Printf("Hello Goroutine:%d!\n", index)
			wait.Done()
		}(i)

	}
	wait.Wait()
}

//同步版本
func demoUpperSync() {
	toUpperSync("Hello Callbacks!", func(v string) {
		fmt.Printf("Callback:%s\n", v)
	})
}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}

//异步版本
func demoUpperASync() {
	//使用WaitGroup
	var wait sync.WaitGroup
	wait.Add(1)
	toUpperASync("Hello Callbacks(Async)!", func(v string) {
		fmt.Printf("Callback:%s\n", v)
		wait.Done()
	})
	println("Waiting async response...")
	wait.Wait()
}
func toUpperASync(word string, f func(string)) {
	//使用goroutine就成为异步函数了
	go func() {
		f(strings.ToUpper(word))
	}()
}
