package main

import "shihua_algo/state"

func main() {
	context := state.NewContext()
	context.Off()
	context.On()
	context.On()
	context.Off()
}

//$ go run main.go
//Context准备好了
//已经关闭～
//将状态从关闭切换到打开～
//已经打开了～
//将状态从打开切换到关闭～