package main

import (
	"errors"
	"fmt"
)

type Heap struct {
	// 存储数据
	m []int
	// 堆的大小
	size int
}

func (heap *Heap) insert(node int) error {
	// 插入容器最后，向上堆化
	heap.m = append(heap.m, node)
	heap.size++

	// 当前节点与父节点比较
	i := heap.size
	for {
		if i/2 > 0 && heap.m[i/2] > heap.m[i] {
			heap.m[i/2], heap.m[i] = heap.m[i], heap.m[i/2]
			i = i / 2
		} else {
			break
		}
	}
	return nil
}

func (heap *Heap) delete() error {
	if heap.size == 0 {
		return errors.New("empty")
	}
	// 最后一个元素覆盖根
	heap.m[1] = heap.m[heap.size]
	heapFy(heap.m, 1, heap.size)
	return nil
}

func (heap *Heap) topN(n int) []int {
	if n >= len(heap.m) {
		return heap.m
	}
	// 从n+1个元素开始，与堆顶比较
	for i := n + 1; i < len(heap.m); i++ {
		if heap.m[1] < heap.m[i] {
			heap.m[1], heap.m[i] = heap.m[i], heap.m[1]
			heapFy(heap.m, 1, n)
		}
	}
	return heap.m[1 : n+1]
}

func heapFy(nums []int, i, n int) {
	for {
		minPos := i
		if 2*i <= n && nums[2*i] < nums[minPos] {
			minPos = 2 * i
		}
		if 2*i+1 <= n && nums[2*i+1] < nums[minPos] {
			minPos = 2*i + 1
		}
		if minPos == i {
			break
		}
		nums[minPos], nums[i] = nums[i], nums[minPos]
		i = minPos
	}
}

func BuildHeap(nodes []int, n int) *Heap {
	m := append(make([]int, 1), nodes...)
	// 从n/2到1个节点堆化
	for i := n / 2; i > 0; i-- {
		heapFy(m, i, n)
	}
	return &Heap{
		m:    m,
		size: n,
	}
}

func findKthLargest(m []int, k int) []int {
	// m := append(make([]int, 1), nums...)
	for i := k + 1; i < len(m); i++ {
		if m[1] < m[i] {
			m[1], m[i] = m[i], m[1]
			heapFy(m, 1, k)
		}
	}
	return m
}

func BuildHeapNew(nodes []int, n int) []int {
	m := append(make([]int, 1), nodes...)
	// 从n/2到1个节点堆化
	for i := n / 2; i > 0; i-- {
		heapFy(m, i, n)
	}
	return m
}

func main() {
	m := []int{2, 1}
	// heap := BuildHeap(m)
	// fmt.Printf("builded: %v", heap.m)

	// _ = heap.insert(0)
	// _ = heap.insert(-1)
	// _ = heap.insert(-2)
	// _ = heap.insert(-3)
	// _ = heap.insert(-4)
	// _ = heap.insert(-5)
	// _ = heap.insert(-6)
	// _ = heap.insert(-7)
	// _ = heap.insert(-8)
	// _ = heap.insert(-9)
	// fmt.Printf("\ninserted: %v", heap.m)
	//
	// _ = heap.delete()
	// fmt.Printf("\ndeleted: %v", heap.m)

	// heap := BuildHeap(m, 3)
	// fmt.Printf("builded: %v", heap.m)
	// topN := heap.topN(3)
	// fmt.Printf("\ntopN: %v", topN)

	// heap := &Heap{
	// 	m: m,
	// }
	// topN := heap.topN(1)
	m = BuildHeapNew(m, len(m))
	topN := findKthLargest(m, 2)
	fmt.Printf("\ntopN: %v", topN)
}
