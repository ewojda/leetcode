package main

//https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultRoot := &ListNode{}

	//Add digits
	result := resultRoot
	for {
		if l1 != nil {
			result.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result.Val += l2.Val
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			result.Next = &ListNode{}
			result = result.Next
		} else {
			break
		}
	}

	//Handle carry over
	result = resultRoot
	for result != nil {
		if result.Val >= 10 {
			if result.Next == nil {
				result.Next = &ListNode{}
			}
			result.Next.Val++
			result.Val %= 10
		}
		result = result.Next
	}
	return resultRoot
}

func main() {}
