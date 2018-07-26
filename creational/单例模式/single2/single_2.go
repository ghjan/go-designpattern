package single2

import (
	"sync"
	"fmt"
)

/* single2
代码做了简单的修改了，引入了锁的机制，在GetInstance函数中，每次调用我们都会上一把锁，保证只有一个goroutine执行它，
这个时候并发的问题就解决了。不过现在不管什么情况下都会上一把锁，而且加锁的代价是很大的，
有没有办法继续对我们的代码进行进一步的优化呢？
 */
var m *Manager
var lock = &sync.Mutex{}

func GetInstance() *Manager {
	lock.Lock()
	defer lock.Unlock()
	if m == nil {
		m = &Manager{}
	}
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
