package cancel_by_contex


import (
	"context"
	"fmt"
"testing"
"time"
)

func isCancelled(ctx context.Context) bool  {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}


func TestCancel(t *testing.T)  {
	/*
		通过根context创建子context，当根context被取消时，其它子context也会被取消
	 */
	ctx, cancel := context.WithCancel(context.Background())
	for i:=0;i<5;i++ {
		go func(i int, ctx context.Context) {
			for{
				if isCancelled(ctx){
					break
				}
				time.Sleep(time.Millisecond*10)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second)
}

