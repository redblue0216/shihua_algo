package main

import "shihua_algo/command"

func main() {
	//创建接收者
	receiver := command.NewReceiver()
	cc := command.NewInvoker()
	//创建具体命令对象，如有需要可将其关联至接收者
	cmd1 := command.NewCommand1("commandA", receiver)
	cmd2 := command.NewCommand2("commandB", receiver)
	cc.SetCommand(cmd1)
	cc.SetCommand(cmd2)
	//执行命令
	cc.ExecuteCommand()
}

//operation1: commandA
//operation2: commandB commandB commandB