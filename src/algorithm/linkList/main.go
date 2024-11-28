package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func newLinkList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := new(Node)
	tail := head
	for _, v := range arr {
		tail.next = &Node{value: v}
		tail = tail.next
	}
	return head.next
}

func findCrossNode(list1, list2 *Node) *Node {
	n1, n2 := 0, 0
	cur1, cur2 := list1, list2
	for cur1 != nil {
		n1++
		cur1 = cur1.next
	}

	for cur2 != nil {
		n2++
		cur2 = cur2.next
	}

	findEqualNode := func(h1, h2 *Node) *Node {
		for h1 != nil && h2 != nil {
			if h1 == h2 {
				return h1
			}
			h1 = h1.next
			h2 = h2.next
		}
		return nil
	}

	stepN := func(h *Node, n int) *Node {
		for ; h != nil && n > 0; n-- {
			h = h.next
		}
		return h
	}

	if n1 > n2 {
		diff := n1 - n2
		list1 = stepN(list1, diff)
	} else {
		diff := n2 - n1
		list2 = stepN(list2, diff)
	}
	return findEqualNode(list1, list2)
}

func main() {
	l1 := newLinkList([]int{0, 1, 2, 3, 4, 5})
	l2 := newLinkList([]int{1, 2, 3})
	cur2 := l2
	for cur2.next != nil {
		cur2 = cur2.next
	}
	//     0
	//      1,2,3,4,5
	// 1,2,3
	cur2.next = l1.next
	crosNode := findCrossNode(l1, l2)
	fmt.Printf("Cros: %d", crosNode.value)
}
