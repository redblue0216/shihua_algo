package main

import "shihua_algo/bridge"

func main() {
	concreteImplementor := bridge.NewConcreteImplementor()

	refinedAbstraction := bridge.
		NewRefinedAbstraction(concreteImplementor)
	refinedAbstraction.Execute("Hello Bridge~")
}

//$ go run main.go
//打印信息：[Hello Bridge~]