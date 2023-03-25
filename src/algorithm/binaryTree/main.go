package binaryTree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Stack struct{
	items []*TreeNode
}

func NewStack() *Stack{
	stack := &Stack{}
	stack.items = []*TreeNode{}
	return stack
}

func (stack *Stack) push(item *TreeNode){
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() *TreeNode{
	if len(stack.items) == 0{
		return nil
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[0:len(stack.items)-1]
	return item
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil{
		return result
	}
	stack1 := NewStack()
	stack2 := NewStack()

	curNode := root
	// 从左向右遍历，子节点从左向右入栈1
	result = append(result, []int{curNode.Val})
	if curNode.Left != nil{
		stack1.push(curNode.Left)
	}
	if curNode.Right != nil{
		stack1.push(curNode.Right)
	}

	for len(stack1.items) > 0 || len(stack2.items) > 0{
		level := []int{}
		// 栈1不为空就出栈
		for len(stack1.items) > 0{
			curNode = stack1.pop()
			level = append(level, curNode.Val)
			// 从右向左遍历，子节点从右向左入栈2
			if curNode.Right != nil{
				stack2.push(curNode.Right)
			}
			if curNode.Left != nil{
				stack2.push(curNode.Left)
			}
		}
		if len(level) > 0{
			result = append(result, level)
		}

		level = []int{}
		// 栈2不空就出栈
		for len(stack2.items) > 0{
			curNode = stack2.pop()
			level = append(level, curNode.Val)
			// 从左向右遍历，子节点从左向右入栈1
			if curNode.Left != nil{
				stack1.push(curNode.Left)
			}
			if curNode.Right != nil{
				stack1.push(curNode.Right)
			}
		}
		if len(level) > 0{
			result = append(result, level)
		}
	}
	return result
}