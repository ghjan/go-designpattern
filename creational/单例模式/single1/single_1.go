package single1

import "fmt"

/*
设计模式-单例模式(Go语言描述)
https://blog.csdn.net/qibin0506/article/details/50733314
保证一个类仅有一个实例，并提供一个访问它的全局访问点。

 */
/* single1
没有考虑并发的实现
现在我们是在并发的情况下去调用的 GetInstance函数，现在恰好第一个goroutine执行到m = &Manager {}这句话之前，
第二个goroutine也来获取实例了，第二个goroutine去判断m是不是nil,因为m = &Manager{}还没有来得及执行，
所以m肯定是nil，现在出现的问题就是if中的语句可能会执行两遍！
 */

var m *Manager

func GetInstance() *Manager {
	if m == nil {
		m = &Manager{}
	}
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
