package completed

import "fmt"

func main() {
	ll1 := createLinkedList([]int{1, 2, 3, 4, 5})
	ll2 := createLinkedList([]int{0, 1, 2})

	rotateRight(ll1, 2)
	rotateRight(ll2, 4)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func rotateRight(head *ListNode, k int) *ListNode {
	//var count int

	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode

	for i := 0; i < k; i++ {
		temp := head
		for temp.Next != nil {
			prev = temp
			temp = temp.Next
		}
		prev.Next = nil
		temp.Next = head
		head = temp
	}

	fmt.Println(prev.Val)
	fmt.Println(head.Val)

	fmt.Println("Module Here", 2000000000%3)

	printLinedList(head)
	return head
}

//func createLinkedList(nums []int) *ListNode {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	head := &ListNode{Val: nums[0]}
//	current := head
//	for i := 1; i < len(nums); i++ {
//		current.Next = &ListNode{Val: nums[i]}
//		current = current.Next
//	}
//	return head
//}
//
//func printLinedList(head *ListNode) {
//	temp := head
//	for temp != nil {
//		fmt.Printf("%d: ", temp.Val)
//		temp = temp.Next
//	}
//	fmt.Println()
//}

//func lenList(head *ListNode) int {
//	temp := head
//	var count int
//	for temp != nil {
//		temp = temp.Next
//		count++
//	}
//	return count
//}
