package main

import (
	"fmt"
	"shihua_algo/singleton"
)

func main() {
	for i := 0; i < 3; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()
}

//创建单个实例
//已创建单个实例!
//已创建单个实例!