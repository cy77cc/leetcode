package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l1
		l1 = l1.Next
	} else {
		res = l2
		l2 = l2.Next
	}
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val < l2.Val {
			res
		}
	}
}

func main() {

}
