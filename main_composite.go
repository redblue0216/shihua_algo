package main

import "shihua_algo/composite"

func main() {
	composite_instance := composite.NewComposite()
	leaf1 := composite.NewLeaf(99)
	composite_instance.Add(leaf1)
	leaf2 := composite.NewLeaf(100)
	composite_instance.Add(leaf2)
	leaf3 := composite.NewComposite()
	composite_instance.Add(leaf3)
	composite_instance.Execute()
}

//$ go run main.go
//99  100