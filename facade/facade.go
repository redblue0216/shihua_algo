package facade

import "fmt"

//外观类
type Facade struct {
	subSystemA SubSystemA
	subSystemB SubSystemB
}

//初始化
func NewFacade() *Facade {
	return &Facade{
		subSystemA: SubSystemA{},
		subSystemB: SubSystemB{},
	}
}

//外观方法A
func (c *Facade) MethodA() {
	c.subSystemB.MethodThree()
	c.subSystemA.MethodOne()
	c.subSystemB.MethodFour()
}

//外观方法B
func (c *Facade) MethodB() {
	c.subSystemB.MethodFour()
	c.subSystemA.MethodTwo()
}

//子系统A
type SubSystemA struct {
}

//初始化子系统B
func NewSubSystemA() *SubSystemA {
	return &SubSystemA{}
}

//子系统B方法
func (c *SubSystemA) MethodOne() {
	fmt.Println("SubSystemB - MethodOne")
}

//子系统B方法
func (c *SubSystemA) MethodTwo() {
	fmt.Println("SubSystemB - MethodTwo")

}

//子系统B
type SubSystemB struct {
}

//初始化子系统A
func NewSubSystemB() *SubSystemB {
	return &SubSystemB{}
}

//子系统A方法
func (c *SubSystemB) MethodThree() {
	fmt.Println("SubSystemA - MethodThree")
}

//子系统A方法
func (c *SubSystemB) MethodFour() {
	fmt.Println("SubSystemA - MethodFour")
}