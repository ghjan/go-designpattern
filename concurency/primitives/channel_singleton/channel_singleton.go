package channel_singleton

type singleton struct{}

var instance singleton

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int) //传递channel的channel
var quitCh chan bool = make(chan bool)
/*
The init() function in any package will get executed on program execution,
so we don't need to worry about executing this function specifically from our code.
 */
func init() {
	var count int
	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count //得到一个 chan int类型的管道
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh

}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
