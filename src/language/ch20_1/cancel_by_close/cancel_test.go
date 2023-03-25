package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool  {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func close_1(cancelChan chan struct{})  {
	cancelChan <- struct {}{}
}

func close_2(cancelChan chan struct{})  {
	// close是广播机制的，当关闭时所有接收ch的地方都会收到
	close(cancelChan)
}

func TestCancel(t *testing.T)  {
	cancelChan := make(chan struct{}, 0)
	for i:=0;i<5;i++ {
		go func(i int, ch chan struct{}) {
			for{
				if isCancelled(ch){
					break
				}
				time.Sleep(time.Millisecond*10)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//close_1(cancelChan)
	close_2(cancelChan)
	time.Sleep(time.Second)
}
