package main

import "testing"

func TestClone(t *testing.T) {
	// 1 -> 2 -> 3, with randoms: 1->3, 2->1, 3->3
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n1.Next, n2.Next = n2, n3
	n1.Random, n2.Random, n3.Random = n3, n1, n3

	c1 := Clone(n1)
	c2 := c1.Next
	c3 := c2.Next

	// Structure and values.
	if c1.Val != 1 || c2.Val != 2 || c3.Val != 3 || c3.Next != nil {
		t.Fatalf("clone structure wrong")
	}
	// Random topology mirrors the original but points at clone nodes.
	if c1.Random != c3 || c2.Random != c1 || c3.Random != c3 {
		t.Errorf("clone random pointers wrong")
	}
	// Clones are new nodes.
	if c1 == n1 || c2 == n2 || c3 == n3 {
		t.Errorf("clone shares nodes with the original")
	}
	// Original list restored intact.
	if n1.Next != n2 || n2.Next != n3 || n3.Next != nil {
		t.Errorf("original Next chain damaged")
	}
	if n1.Random != n3 || n2.Random != n1 || n3.Random != n3 {
		t.Errorf("original Random pointers damaged")
	}

	if Clone(nil) != nil {
		t.Errorf("Clone(nil) should be nil")
	}
}
