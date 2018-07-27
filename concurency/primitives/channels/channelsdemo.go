package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	println("----channelDemo1")
	channelDemo1()
	println("----channelDemo2")
	channelDemo2()
	println("----bufferedChannelDemo1")
	bufferedChannelDemo1()
	bufferedChannelDemo2()
	directionChannel1()
	selectDemo1()
	time.Sleep(time.Second * 3)
	rangeOverChannelDemo()
}

//channelDemo1
// default unbuffered channel
// send block & receive block
func channelDemo1() {
	channel := make(chan string)
	go func() {
		channel <- "hello channel"
	}()
	message := <-channel
	fmt.Println(message)
}

//channelDemo2 次序没有保证
func channelDemo2() {
	channel := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		channel <- "hello channel 2"
		println("Finishing goroutine in channelDemo2")
		wg.Done()
	}()
	time.Sleep(time.Microsecond * 300)
	message := <-channel
	fmt.Println(message)
	wg.Wait()
}

//bufferedChannelDemo1 次序基本上是固定的
func bufferedChannelDemo1() {
	channel := make(chan string, 1) //第一个string消息不会block
	go func() {
		channel <- "Hello buffered channel"
		println("Finishing goroutine in bufferedChannelDemo1")
	}()
	time.Sleep(time.Microsecond * 200)

	message := <-channel
	fmt.Println(message)
}

func bufferedChannelDemo2() {
	channel := make(chan string, 1) //第一个string消息不会block
	go func() {
		channel <- "Hello buffered channel 1"
		channel <- "Hello buffered channel 2"

		println("Finishing goroutine in bufferedChannelDemo2")
	}()
	time.Sleep(time.Microsecond * 200)

	message := <-channel
	fmt.Println(message)
}

func directionChannel1() {
	channel := make(chan string, 1) //第一个string消息不会block
	go func(ch chan<- string) { // receiver only chan
		ch <- "directionChannel1"
		//msg:= <- ch  //error
		println("Finishing goroutine in directionChannel1")
	}(channel)
	time.Sleep(time.Microsecond * 200)

	message := <-channel
	fmt.Println(message)
}

func selectDemo1() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)

	go receiver(helloCh, goodbyeCh, quitCh)

	go sendString(helloCh, "hello selectDemo1!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "goodbye!")
	<-quitCh
}

func sendString(ch chan<- string, s string) {
	ch <- s
}
func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			fmt.Println(msg)
		case msg := <-goodbyeCh:
			fmt.Println(msg)
		case <-time.After(time.Second * 2):
			println("(receiver)Nothing received in 2 seconds!Exiting")
			quitCh <- true
			return
		}
	}
}

/* rangeOverChannelDemo
until the concurrent Goroutine closes this channel. At that moment, the range finishes and the app can exit.
Range is very useful in taking data from a channel,
and it's commonly used in fan-in patterns where many different Goroutines send data to the same channel.
 */
func rangeOverChannelDemo() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("rangeOverChannelDemo:%d\n", v)
	}
}
