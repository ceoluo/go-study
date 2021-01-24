package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10;i++ {
			ch <- i
		}
		// 关闭通道，在接收端可通过通道放回的第二个值判断通道是否关闭，从而不用一直等待
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for {
			if data, ok := <- ch;ok{
				fmt.Println(data)
			}else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChan(t *testing.T)  {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}