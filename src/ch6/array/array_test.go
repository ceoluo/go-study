package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr1[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTavel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	// 第一种遍历方式
	t.Log("第一种遍历方式")
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	// 第二种遍历方式
	t.Log("第二种遍历方式")
	for idx, v := range arr2 {
		t.Log(idx, v)
	}

	// 不关心索引
	t.Log("不关心索引")
	for _, v := range arr2 {
		t.Log(v)
	}
}

func TestMutiDimArry(t *testing.T) {
	var mutiDimArray = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	t.Log(mutiDimArray)

	mutiDimArray2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	t.Log(mutiDimArray2)
}

func TestArraySection(t *testing.T) {
	var mutiDimArray = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mutidimarray3 := mutiDimArray[:2]
	t.Log(mutidimarray3)
}
