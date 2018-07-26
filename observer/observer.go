package main

import (
	"fmt"
	"time"
)

/*

Golang设计模式之观察者模式
https://blog.csdn.net/Jeanphorn/article/details/78784197
 */

//Event 事件
type Event struct {
	Data string
}

//Observer 观察者接口
type Observer interface {
	//更新事件
	Update(*Event)
}

//Subject 被观察的对象接口
type Subject interface {
	//注册观察者
	Register(Observer)
	//注销观察者
	Unregister(Observer)

	//通知观察者事件
	Notify(*Event)
}

type ConcreteObserver struct {
	Id int
}

func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", co.Id, e.Data)
}

type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreteSubject) Register(ob Observer) {
	cs.Observers[ob] = struct{}{} //类似于一个bool值
}

func (cs *ConcreteSubject) Unregister(ob Observer) {
	delete(cs.Observers, ob)
}

func (cs *ConcreteSubject) Notify(e *Event) {
	for ob := range cs.Observers {
		ob.Update(e)
	}
}

func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer]struct{}),
	}

	co1 := &ConcreteObserver{1}
	co2 := &ConcreteObserver{2}

	cs.Register(co1)
	cs.Register(co2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)

		time.Sleep(time.Duration(1) * time.Second)
	}
}
