package main

import "shihua_algo/template_method"

func main() {
	concreteClassA := template_method.NewAbstractClass(&template_method.ConcreteClassA{})
	concreteClassA.TemplateMethod()
	concreteClassB := template_method.NewAbstractClass(&template_method.ConcreteClassB{})
	concreteClassB.TemplateMethod()
}

//$ go run main.go
//ConcreteClassA Step1
//ConcreteClassA Step2
//ConcreteClassA Step3
//ConcreteClassB Step1
//ConcreteClassB Step2
//ConcreteClassB Step3