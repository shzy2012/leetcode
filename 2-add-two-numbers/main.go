package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0} //link head
	link := head              //临时link，用于将新的node添加到head
	carry := 0                //进位
	for l1 != nil || l2 != nil {

		tmp1Val := 0
		tmp2Val := 0
		if l1 != nil {
			tmp1Val = l1.Val
		}
		if l2 != nil {
			tmp2Val = l2.Val
		}

		sum := carry + tmp1Val + tmp2Val
		carry = sum / 10
		link.Next = &ListNode{Val: sum % 10} //求余，并加入到link
		link = link.Next                     //添加到link中

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		link.Next = &ListNode{Val: 1}
	}

	return head.Next
}

func linkToArray(head *ListNode) []int {

	res := []int{}
	tmpLink := head
	for {

		res = append(res, tmpLink.Val)
		if tmpLink.Next == nil {
			return res
		}
		tmpLink = tmpLink.Next
	}
}

func main() {
	l1 := &ListNode{}
	l1.Val = 2
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{}
	l2.Val = 5
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 4}

	fmt.Printf("%v\n", linkToArray(addTwoNumbers(l1, l2)))
}
