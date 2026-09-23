//go:build !solution

package linkedlist

func Prepend(head *Node, value int) *Node {
	return &Node{Value: value, Next: head}
}

func Length(head *Node) int {
	length := 0
	for n := head; n != nil; n = n.Next {
		length++
	}

	return length
}

func Find(head *Node, value int) *Node {
	for n := head; n != nil; n = n.Next {
		if n.Value == value {
			return n
		}
	}

	return nil
}

func Reverse(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
