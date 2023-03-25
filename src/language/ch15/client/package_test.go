package client

import (
	"go_learning/src/ch15/series"
	"testing"
)

/*
1 两种导包的方式：
	1】使用GOPATH，但需要关闭go mod模式管理包
	2】使用go mod, 但需要在公共目录下执行go mod init , 初始化包
2 想要试包可被其它包导入使用，其名字必须是大写
 */

func TestPackage(t *testing.T)  {
	fibList := series.GetFibonacciSeries(10)
	t.Log(fibList)
	t.Log(series.Square(8))
}
