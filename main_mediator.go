package main

import "shihua_algo/mediator"

func main() {
	mediator_instance := mediator.NewMediator()
	mediator_instance.ConcreteColleague2.Talk()
}

//$ go run main.go
//通过中介者谈话
//具体同事1：ConcreteColleague1回复中...
//具体同事2：ConcreteColleague2回复中...