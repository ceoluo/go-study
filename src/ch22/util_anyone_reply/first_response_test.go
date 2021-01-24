package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string  {
	time.Sleep(time.Millisecond*10)
	return fmt.Sprintf("The result is from: %d", i)
}

func FirstResponse() string {
	numberOfTask := 10
	// bufferedChannel 防阻塞
	ch := make(chan string, numberOfTask)
	for i:=0;i<numberOfTask;i++{
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <- ch
}

func TestFirstResponse(t *testing.T)  {
	// 输出当前系统中的协程数
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second*1)
	t.Log("After:", runtime.NumGoroutine())
}