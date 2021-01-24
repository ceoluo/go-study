package benchmark

import (
	"bytes"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T)  {
	assertions := assert.New(t)
	elems := []string{"1","2","3","4","5"}
	ret := ""
	for _,elem := range elems{
		ret += elem
	}
	assertions.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T)  {
	assertions := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1","2","3","4","5"}
	for _,elem := range elems{
		buf.WriteString(elem)
	}
	assertions.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1","2","3","4","5"}
	b.ResetTimer()
	for i := 0;i<b.N;i++{
		ret := ""
		for _,elem := range elems{
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B)  {
	var buf bytes.Buffer
	b.ResetTimer()
	elems := []string{"1","2","3","4","5"}
	for i := 0;i<b.N;i++ {
		for _,elem := range elems{
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
