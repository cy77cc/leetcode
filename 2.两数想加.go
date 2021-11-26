package main

type ListNode struct {
	Val int
	Next *ListNode
}

func getNum(L *ListNode, index int) (int, bool) {
	tmp := L
	for i := 0; i < index; i++ {
		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			return 0, false
		}
	}
	return tmp.Val, true
}

func append(L *ListNode, value int) {
	if L.Val == -1 && L.Next == nil {
		L.Val = value
		return
	}
	tmp := L
	for ; tmp.Next != nil; {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{value, nil}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	L := &ListNode{-1, nil}
	falg1, falg2 := true, true
	var num1, num2 int
	var nextval, result int
	for i := 0; ; i++ {
		num1, falg1 = getNum(l1, i)
		num2, falg2 = getNum(l2, i)
		if !falg1 && !falg2 {
			break
		}
		result = num1 + num2 + nextval
		if result >= 10 {
			result -= 10
			nextval = 1
		} else {
			nextval = 0
		}
		append(L, result)
	}
	if nextval == 1 {
		append(L, nextval)
	}
	return L
}

func main() {

}
