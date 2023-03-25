package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type singleton struct {
	
}

var singletonInstance *singleton
var once sync.Once

func getSingletonObj() *singleton {
	/*
		应用go内置的只运行一次的once实现单例模式
	 */
	once.Do(func() {
		fmt.Println("Create obj")
		singletonInstance = new(singleton)
	})
	return singletonInstance
}

func TestGetSingletonObj(t *testing.T)  {
	var wg sync.WaitGroup
	for i:=0;i<10;i++ {
		wg.Add(1)
		go func() {
			obj := getSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}