package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicRecover(t *testing.T)  {
	// 延迟函数，最后执行，用于错误处理等
	defer func() {
		// recover 恢复错误，处理错误
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
			//os.Exit(1)
		}
		// 在不确定时，干脆Let it crash！
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}
