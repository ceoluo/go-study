package lazy

import (
	"fmt"
	"sync"
)

// 懒汉式单例
type Singleton struct {

}

var (
	singleton *Singleton
	once = &sync.Once{}
)

func GetSingletonInstance() *Singleton {
	// 双重检测, 提升性能
	if singleton == nil {
		once.Do(func() {
			fmt.Printf("create a lazy singeleton instance\n")
			singleton = &Singleton{}
		})
	}
	return singleton
}

