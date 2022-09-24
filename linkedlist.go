// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type ListNode struct {
	num  int
	next *ListNode
}

type linklList struct {
	node *ListNode
}

func (list *linklList) display() {
	ptr := list.node

	for ptr.next != nil {
		fmt.Printf("\n item %v", ptr.num)
		ptr = ptr.next
	}
	fmt.Printf("\n last item %v", ptr.num)

}
func (list *linklList) insert(i int) {
	node := &ListNode{num: i, next: nil}
	if list.node == nil {
		list.node = node
		return
	} else {
		ptr := list.node
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = node
		return
	}
}
func main() {
	list := linklList{}
	for i := 1; i < 9; i++ {
		list.insert(i)
	}
	list.display()

}
