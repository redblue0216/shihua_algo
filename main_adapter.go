package main

import (
	"shihua_algo/adapter"
)

func main() {
	//创建客户端
	adapter := adapter.Adapter{}
	adapter.Execute()
}

//$ go run main.go
//最终执行的方法