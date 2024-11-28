package main

func bubbleSort(nums []int) {
	flag := true
	for i := 0; i < len(nums) && flag; i++ {
		flag = false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
	}
}

func selectSort(nums []int) {
	k := 0
	for i := k; i < len(nums); i++ {
		min := k
		for j := k; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
		k++
	}
}

func insertSort(nums []int) {
	for j := 1; j < len(nums); j++ {
		tmp := nums[j]
		var k int
		for k = j - 1; k >= 0; k-- {
			if nums[k] > tmp {
				nums[k+1] = nums[k]
				continue
			}
			break
		}
		nums[k+1] = tmp
	}
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	p, q := start, end-1
	for p < q {
		for nums[q] > pivot && p != q {
			q--
		}
		nums[p] = nums[q]

		for nums[p] < pivot && p != q {
			p++
		}
		nums[q] = nums[p]
	}
	nums[p] = pivot
	return p
}

func partition2(nums []int, start, end int) int {
	pivot := nums[start]
	p, q := start, end-1
	for p < q {
		for nums[q] >= pivot && p != q {
			q--
		}

		for nums[p] <= pivot && p != q {
			p++
		}
		nums[p], nums[q] = nums[q], nums[p]
	}
	nums[p] = pivot
	return p
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	// part := partition(nums, start, end)
	part := partition2(nums, start, end)

	quickSort(nums, start, part)
	quickSort(nums, part+1, end)
	return
}

func merge(nums []int, start1, end1, start2, end2 int) {
	p, q := start1, start2
	if start1 >= len(nums) || start2 >= len(nums) {
		return
	}
	var merged []int
	for p < end1 && q < end2 {
		if nums[p] <= nums[q] {
			merged = append(merged, nums[p])
			p++
		} else {
			merged = append(merged, nums[q])
			q++
		}
	}
	for ; p < end1; p++ {
		merged = append(merged, nums[p])
	}
	for ; q < end2; q++ {
		merged = append(merged, nums[q])
	}

	for i := 0; i < len(merged); i++ {
		nums[start1+i] = merged[i]
	}
	return
}

func mergeSortLoop(nums []int) {
	// 归并排序-非递归写法
	k := 1
	var i int
	for k < len(nums) {
		for i = 0; i < len(nums)-2*k; i += 2 * k {
			merge(nums, i, i+k, i+k, i+2*k)
		}
		if i+k < len(nums) {
			merge(nums, i, i+k, i+k, len(nums))
		}
		k *= 2
	}
}

func mergeSortReverse(nums []int, left, right int) {
	// 归并排序，递归写法
	if left < right {
		mid := left + ((right - left) >> 1)
		mergeSortReverse(nums, left, mid)
		mergeSortReverse(nums, mid+1, right)
		merge(nums, left, mid, mid, right)
	}
}

func search(nums []int, target int) bool {
	low, heigh := 0, len(nums)-1
	for low < heigh {
		mid := (heigh + low) >> 1
		if target == nums[mid] {
			return true
		}
		if nums[low] >= nums[heigh] {
			if target < nums[mid] && target > nums[heigh] {
				heigh = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target < nums[mid] {
				heigh = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tmpMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := tmpMap[v]; ok {
			tmpMap[v]++
		} else {
			tmpMap[v] = 1
		}
	}

	longestCount := 0
	for k, v := range tmpMap {
		if _, ok := tmpMap[k-1]; !ok {
			curCount := v
			curNum := k
			for {
				if _, ok2 := tmpMap[curNum+1]; !ok2 {
					break
				}
				curCount += tmpMap[curNum+1]
				curNum++
			}
			if longestCount < curCount {
				longestCount = curCount
			}
		}
	}

	return longestCount
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	result := []int{-1, -1}
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			result[0] = binarySearch(nums, left, mid-1, target, true)
			result[1] = binarySearch(nums, mid+1, right, target, false)
			return result
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func binarySearch(nums []int, left int, right int, target int, flag bool) int {
	for left <= right {
		mid := (left + right) >> 1
		if (mid == 0) || (mid == left && mid == right) {
			if nums[mid] == target {
				return mid
			} else {
				return -1
			}
		}

		if flag {
			if nums[mid] == target && nums[mid-1] != nums[mid] {
				return mid
			} else if nums[mid] != target {
				return binarySearch(nums, mid+1, right, target, false)
			} else {
				return binarySearch(nums, left, mid-1, target, true)
			}
		} else {
			if nums[mid] == target && nums[mid+1] != nums[mid] {
				return mid
			} else if nums[mid] != target {
				return binarySearch(nums, left, mid-1, target, true)
			} else {
				return binarySearch(nums, mid+1, right, target, false)
			}
		}
	}
	return -1
}

type heap struct {
	arr []int
}

func (h *heap) push(v int) {
	h.arr = append(h.arr, v)
	min := 0
	for i := 0; i < len(h.arr); i++ {
		if h.arr[min] > h.arr[i] {
			min = i
		}
	}
	h.arr[0], h.arr[min] = h.arr[min], h.arr[0]
}

func (h *heap) pop() int {
	v := h.arr[0]
	h.arr = h.arr[1:len(h.arr)]
	return v
}

func nthUglyNumber(n int) int {
	h := heap{[]int{1}}
	factor := []int{2, 3, 5}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := h.pop()
		if i == n {
			return x
		}

		for _, f := range factor {
			next := x * f
			if _, ok := seen[next]; !ok {
				seen[next] = struct{}{}
				h.push(next)
			}
		}
	}
}

type Item struct {
	number int
	sub    []byte
}

func bytestoint(b []byte) int {
	res := int(b[0] - '0')
	for i := 1; i < len(b); i++ {
		res = res*10 + int(b[i]-'0')
	}
	return res
}

func decodeString(s string) string {
	stack := []*Item{}
	t := []byte{}
	sub := []byte{}
	for i := 0; i < len(s); {
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			t = append(t, s[i])
			i++
		}

		if len(t) > 0 && i == len(s) {
			break
		}

		if len(t) > 0 && s[i] == '[' {
			newsub := make([]byte, len(sub), len(sub))
			copy(newsub, sub)
			stack = append(stack, &Item{bytestoint(t), newsub})
			t = t[:0]
			sub = sub[:0]
		} else if s[i] >= 'a' && s[i] <= 'z' {
			sub = append(sub, s[i])
		} else if s[i] == ']' {
			item := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := sub
			for j := 0; j < item.number-1; j++ {
				sub = append(sub, tmp...)
			}
			sub = append(item.sub, sub...)
		}
		i++
	}
	return string(sub)
}

func rotate(nums []int, k int) {
	if k == 0 || len(nums) == k {
		return
	}
	// 反转数组
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}

	mod := k % len(nums)
	// 反转[mod,len)
	i, j := mod, len(nums)-1
	for i <= (i+j)/2 {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	// 前k位反转
	for i := 0; i < mod/2; i++ {
		nums[i], nums[mod-1-i] = nums[mod-i-1], nums[i]
	}
}

func main() {
	a := []int{5, 7, 7, 8, 8, 10}
	// merge(a, 0, 2, 2,4)
	// mergeSortLoop(a)
	// mergeSortReverse(a, 0, len(a))
	quickSort(a, 0, len(a))
	// bubbleSort(a)
	// selectSort(a)
	// insertSort(a)
	// f := search(a, 0)
	// f := longestConsecutive(a)
	// f := searchRange(a, 8)

	// f := nthUglyNumber(10)
	// f := decodeString("3[a]2[bc]")
	// f := []int{-1, -100, 3, 99}
	// rotate(f, 2)
	// fmt.Printf("%v", f)
}
