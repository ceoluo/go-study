package share_memery

import (
	"sync"
	"testing"
	"time"
)

func TestCounterThreadUnsafe(t *testing.T)  {
	/*
		协程不安全的
	 */
	counter := 0
	
	for i:=0;i<5000;i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second*1)
	// 输出结果与预计不一样
	t.Logf("counter = %d", counter)
}


func TestCounterThreadSafe(t *testing.T)  {
	/*
		协程安全的
	*/
	var mut sync.Mutex
	counter := 0
	for i:=0;i<5000;i++ {
		go func() {
			// 处理忘记解锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second*1)
	t.Logf("counter = %d", counter)
}

func TestCounterWaiteGroup(t *testing.T)  {
	/*
		协程安全的, 等待所有协程全部执行完成才退出程序
	*/
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i:=0;i<5000;i++ {
		wg.Add(1)
		go func() {
			// 处理忘记解锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	//time.Sleep(time.Second*1)
	// 等待所以协程结束
	wg.Wait()
	t.Logf("counter = %d", counter)
}
