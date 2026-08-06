package main

import (
	"shihua_algo/strategy"
)

func main() {
	strategyB := strategy.NewStrategyB()
	context := strategy.NewContext()
	context.SetStrategy(strategyB)
	strategyA := strategy.NewStrategyA()
	context.SetStrategy(strategyA)
	context.Execute()
}

//$ go run main.go
//执行策略 A