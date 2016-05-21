/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    head := new(ListNode)
	curp := head

	x := 0
	y := 0
	sum := 0
	flag := 0
	for l1 != nil || l2 != nil {

		x = 0
		if l1 != nil {
			x = l1.Val
		}

		y = 0
		if l2 != nil {
			y = l2.Val
		}

		sum = x + y + flag

		flag = sum / 10
		sum = sum % 10

		tmp := new(ListNode)
		tmp.Val = sum
		curp.Next = tmp
		curp = curp.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if flag > 0 {
		tmp := new(ListNode)
		tmp.Val = flag
		curp.Next = tmp
	}

	return head.Next
}