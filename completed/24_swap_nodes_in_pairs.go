package completed

import "fmt"

/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without
modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
*/

func main() {

	ll1 := createLinkedList([]int{1, 2, 3, 4})
	swapPairs(ll1)
	ll2 := createLinkedList([]int{})
	swapPairs(ll2)
	ll3 := createLinkedList([]int{1})
	swapPairs(ll3)
	ll4 := createLinkedList([]int{1, 2, 3})
	swapPairs(ll4)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		//printLinedList(head)
		return head
	}

	idx := 0

	next := head.Next
	curr := head

	if lenList(head) >= 2 {
		head = head.Next
	}

	for next != nil {
		if idx%2 == 0 {
			curr.Next = next.Next
			next.Next = curr
			next = curr.Next

		} else {
			if next.Next != nil && idx >= 1 {
				curr.Next = next.Next
			}
			curr = next
			next = next.Next
		}
		idx++
	}

	//printLinedList(head)

	return head
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func printLinedList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d: ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func lenList(head *ListNode) int {
	temp := head
	var count int
	for temp != nil {
		temp = temp.Next
		count++
	}
	return count
}
