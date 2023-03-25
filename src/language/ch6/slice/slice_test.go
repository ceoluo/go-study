package slice

import "testing"

/*
	切片：
		数据结构：(指针ptr, 元素的个数len, 内部数组的容量cap)
	切片共享存储结构
*/

func TestSliceInit(t *testing.T) {
	// 与数组的区别就是：没有声明长度
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 2)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

	t.Log(s2[0])

}

func TestSliceGrowing(t *testing.T) {
	var s = []int{}
	// 当len超过cap时，cap成倍增长
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemery(t *testing.T) {
	//切片共享存储结构
	year := []string{"", "Jan", "Fab", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(year, len(year), cap(year))
	Q1 := year[3:6]
	// 因为是共享存储空间，所以这里的切片Q1与原始切片共享数据和剩余的cap
	t.Log(Q1, len(Q1), cap(Q1))

	summer := year[5:8]
	t.Log(summer)

	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"

	t.Log(summer)
	t.Log(year)
}

/*
	数组和切片的对比：
		1、切片可伸缩
		2、数组可比较
*/
func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	t.Log(s1)
	t.Log(s2)
	// invalid operation: s1 == s2 (slice can only be compared to nil)
	//t.Log(s1==s2)

}
