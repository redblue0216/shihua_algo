package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var instance *singleton

//获取实例
func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("创建单个实例")
			instance = new(singleton)
		} else {
			fmt.Println("已创建单个实例!")
		}
	} else {
		fmt.Println("已创建单个实例!")
	}

	return instance
}