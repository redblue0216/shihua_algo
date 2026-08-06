package main

import (
	"fmt"
	"shihua_algo/chain_of_responsibility"
)

func main() {
	barry := chain_of_responsibility.NewBaseHandler("Barry", nil, 1)
	shirdon := chain_of_responsibility.NewBaseHandler("Shirdon", barry, 2)
	jack := chain_of_responsibility.NewBaseHandler("Shirdon", shirdon, 3)
	res := shirdon.Handle(2)
	res1 := jack.Handle(3)
	fmt.Println(res)
	fmt.Println(res1)
}

//$ go run main.go
//ConcreteHandler handleID: 2
//Shirdon
//ConcreteHandler handleID: 3
//Barry
//ConcreteHandler handleID: 3
//Shirdon
//3
//4