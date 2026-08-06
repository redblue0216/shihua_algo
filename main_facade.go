package main

import (
	"shihua_algo/facade"
)

func main() {
	fa := facade.NewFacade()
	fa.MethodA()
	fa.MethodB()

	sub := facade.NewSubSystemA()
	sub.MethodOne()
	sub.MethodTwo()
}

//$ go run main.go
//SubSystemA - MethodThree
//SubSystemB - MethodOne
//SubSystemA - MethodFour
//SubSystemA - MethodFour
//SubSystemB - MethodTwo
//SubSystemB - MethodOne
//SubSystemB - MethodTwo