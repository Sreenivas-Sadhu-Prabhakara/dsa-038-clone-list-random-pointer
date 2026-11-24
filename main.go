package main

import "fmt"

// Node has a Next pointer and a Random pointer to any node in the list (or nil).
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Clone deep-copies the list in O(1) extra space by interleaving copy nodes
// with originals, wiring Randoms through the interleaving, then unzipping.
func Clone(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. A -> A' -> B -> B' ...
	for p := head; p != nil; p = p.Next.Next {
		p.Next = &Node{Val: p.Val, Next: p.Next}
	}
	// 2. The copy of p.Random sits right after p.Random.
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	// 3. Unzip, restoring the original list.
	dummy := &Node{}
	tail := dummy
	for p := head; p != nil; p = p.Next {
		tail.Next = p.Next
		tail = tail.Next
		p.Next = p.Next.Next
	}
	return dummy.Next
}

func main() {
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	a.Next, a.Random = b, b
	b.Random = b
	c := Clone(a)
	fmt.Println(c.Val, c.Random.Val, c.Next.Random.Val)
}
