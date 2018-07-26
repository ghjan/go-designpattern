package single3

import (
	"sync"
	"fmt"
)

/* single3
加锁的代价是很大的，有没有办法继续对我们的代码进行进一步的优化呢？
熟悉java的同学可能早就想到了双重的概念，没错，在go中我们也可以使用双重锁机制来提高效率。
 */
var m *Manager
var lock *sync.Mutex = &sync.Mutex{}

func GetInstance() *Manager {
	if m == nil {
		lock.Lock()
		defer lock.Unlock()
		if m == nil {
			m = &Manager{}
		}
	}

	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
