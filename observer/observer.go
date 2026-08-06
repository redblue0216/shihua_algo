package observer

import (
	"fmt"
)

// 要观察和通知的事件
type Event struct {
	Id string
}

// 观察者接口
type Observer interface {
	Update(event Event)
}

// 具体观察者
type ConcreteObserver struct {
	name string
}

//创建一个新的具体观察者对象
func NewObserver(name string) Observer {
	return &ConcreteObserver{name}
}

// 具体观察者的方法
func (o *ConcreteObserver) Update(event Event) {
	fmt.Printf("ConcreteObserver '%s' received event '%s'\n", o.name, event.Id)
}

// 抽象主题
type Subject struct {
	ObserverCollection []Observer
}

// 注册一个新的具体观察者
func (e *Subject) Register(obs Observer) {
	e.ObserverCollection = append(e.ObserverCollection, obs)
}

// 取消注册具体观察者
func (e *Subject) Unregister(obs Observer) {
	for i := 0; i < len(e.ObserverCollection); i++ {
		if obs == e.ObserverCollection[i] {
			e.ObserverCollection = append(e.ObserverCollection[:i], e.ObserverCollection[i+1:]...)
		}
	}
}

// 通知所有具体观察者集合
func (e *Subject) NotifyObservers(event Event) {
	for i := 0; i < len(e.ObserverCollection); i++ {
		e.ObserverCollection[i].Update(event)
	}
}