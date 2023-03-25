package groutine

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T)  {
	n := 10
	for i:=1;i<n;i++{
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}