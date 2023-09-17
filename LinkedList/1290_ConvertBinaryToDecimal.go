package ll1290

import (
	"fmt"
	"math"
)

type ListNode struct {
	data int
	next *ListNode
}

func (head *ListNode) Add(num int) {

	//create new node
	newNode := &ListNode{
		data: num,
	}

	//check head is empty if empty set newnode to head
	if head.data == 0 && head.next == nil {
		head.data = num
		return
	}

	curr := head
	//loop through the list
	for curr.next != nil {
		curr = curr.next
	}
	// add new node to the last
	curr.next = newNode
}

func PrintList(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func GetDecimalValue(head *ListNode) int {
	arr := []int{}
	curr := head
	sum := 0

	for curr != nil {
		//fmt.Printf("%d ", curr.data)
		arr = append(arr, curr.data)
		curr = curr.next
	}

	rev_arr := []int{}
	for j := len(arr) - 1; j >= 0; j-- {
		rev_arr = append(rev_arr, arr[j])
	}

	//fmt.Println(rev_arr)

	for i, _ := range arr {
		//fmt.Println(i)
		sum += int(math.Pow(float64(2), float64(i))) * rev_arr[i]

	}

	return sum
}

// func Reverse(head *ListNode) {
// 	curr := head
// 	prev := &ListNode{}
// 	for curr != nil {
// 		next := curr.next
// 		curr.next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	head = prev
// }
