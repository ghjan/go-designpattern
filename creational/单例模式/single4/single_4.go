package single4

import (
	"sync"
	"fmt"
)

/*
在go中我们还有更优雅的方式去实现。单例的目的是啥？保证实例化的代码只执行一次，
在go中就中这么一种机制来保证代码只执行一次，而且不需要我们手工去加锁解锁。
对，就是我们的sync.Once，它有一个Do方法，在它中的函数go会只保证仅仅调用一次！
 */
var m *Manager
var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
