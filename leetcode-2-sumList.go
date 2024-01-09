package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，
并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func initList(valList []int) *ListNode {
	var head, l *ListNode
	//head := l
	for _, val := range valList {

		tmp := new(ListNode)
		tmp.Val = val
		tmp.Next = nil

		if l == nil {
			head = tmp
			l = tmp
		} else {
			l.Next = tmp
			l = tmp
		}
	}
	return head
}

func showList(node *ListNode) {
	for l := node; l != nil; l = l.Next {
		fmt.Println(l.Val)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, tempResult *ListNode
	flag := 0
	temp1, temp2 := l1, l2
	for ; temp1 != nil && temp2 != nil; temp1, temp2 = temp1.Next, temp2.Next {
		var val ListNode
		val.Val = temp1.Val + temp2.Val
		if flag == 1 {
			val.Val += 1
		}
		if val.Val >= 10 {
			flag = 1
			val.Val -= 10
		} else {
			flag = 0
		}

		if result == nil {
			result = &val
			tempResult = result
		} else {
			tempResult.Next = &val
			tempResult = tempResult.Next
		}
	}
	for ; temp1 != nil; temp1 = temp1.Next {
		var val ListNode
		val.Val = temp1.Val
		if flag == 1 {
			val.Val += 1
		}
		if val.Val >= 10 {
			flag = 1
			val.Val -= 10
		} else {
			flag = 0
		}
		tempResult.Next = &val
		tempResult = tempResult.Next
	}
	for ; temp2 != nil; temp2 = temp2.Next {
		var val ListNode
		val.Val = temp2.Val
		if flag == 1 {
			val.Val += 1
		}
		if val.Val >= 10 {
			flag = 1
			val.Val -= 10
		} else {
			flag = 0
		}
		tempResult.Next = &val
		tempResult = tempResult.Next
	}

	if flag == 1 {
		var val ListNode
		val.Val = 1
		tempResult.Next = &val
	}
	return result
}

// func main() {
// 	firstList := initList([]int{9, 8})
// 	secondList := initList([]int{1})

// 	showList(firstList)
// 	showList(secondList)

// 	result := addTwoNumbers(firstList, secondList)
// 	fmt.Println("show result:")
// 	showList(result)
// }
