package removeduplicates

import "fmt"

// Node is a linked list representation
type Node struct {
	value int
	next  *Node
}

// CountElements returns a map with each element occurrence
func CountElements(head *Node) map[int]int {
	m := make(map[int]int)

	l := head
	for l != nil {
		m[head.value]++

		l = l.next
	}

	return m
}

func RemoveElement(head *Node, element int) *Node {
	if head.value == element {
		return head.next
	}

	l := head
	for l != nil {

		if l.next != nil && l.next.value == element {
			l.next = l.next.next
		} else {
			l = l.next
		}
	}
	return head
}

// RemoveDuplicates returns a linked list without duplicated values
func RemoveDuplicates(head *Node) *Node {
	if head == nil {
		return head
	}

	m := CountElements(head)

	l2 := head
	for l2 != nil {
		if m[l2.value] > 1 {
			head = RemoveElement(head, l2.value)
			m[l2.value]--
		}

		l2 = l2.next
	}

	fmt.Println(head)

	return head
}
