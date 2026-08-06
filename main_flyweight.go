package main

import (
	"fmt"
	"shihua_algo/flyweight"
)

func main() {
	factory := flyweight.NewFlyweightFactory()
	flyweight1 := factory.GetFlyweight("Barry")
	flyweight2 := factory.GetFlyweight("Shirdon")

	fmt.Println(flyweight1.Operation("ok"))
	fmt.Println(flyweight2.Operation("good"))
}

//$ go run main.go
//Barry
//ok
//Shirdon
//good