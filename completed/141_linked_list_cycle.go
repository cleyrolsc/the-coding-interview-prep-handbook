package completed

import "fmt"

//
//import "fmt"

func main() {
	ll1 := createLinkedList([]int{3, 2, 0, -4})
	ll2 := createLinkedList([]int{1, 2})
	ll3 := createLinkedList([]int{1})
	hasCycle(ll1)
	hasCycle(ll2)
	hasCycle(ll3)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
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
