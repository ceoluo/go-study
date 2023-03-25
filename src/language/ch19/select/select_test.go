package _select


import (
	"fmt"
"testing"
"time"
)

func service() string  {
	fmt.Println("service")
	time.Sleep(time.Second*2)
	return "Done"
}

func otherTask()  {
	fmt.Println("Working on something else...")
	time.Sleep(time.Second*1)
	fmt.Println("Other task is done.")
}

func TestService(t *testing.T)  {
	ret := service()
	otherTask()
	t.Log(ret)
}

func AsyncService() chan string {
	// 普通channel, 会阻塞
	//retCh := make(chan string)
	// buffer channel 半阻塞
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T)  {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

func TestSelect(t *testing.T)  {
	// 利用多路选择实现超时机制
	select {
	case res := <- AsyncService():
		t.Log(res)
	case <-time.After(time.Second):
		t.Error("time out!")
	}
}
