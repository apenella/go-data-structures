package gdsexttree

import (
	"errors"
)

// Node is the extended tree graph node
type Node struct {
	Name    string
	Item    interface{}
	Parents []*Node
	Childs  []*Node
}

// AddParent method update node's parents list adding a new one. It also update parent's childs list
func (n *Node) AddParent(parent *Node) error {
	if n == nil {
		return errors.New("(graph::AddParent) Adding parent to a nil node")
	}

	if parent == nil {
		return errors.New("(graph::AddParent) Adding nil parent to node")
	}

	if n.Parents == nil || len(n.Parents) == 0 {
		n.Parents = []*Node{}
	}

	if !n.HasParent(parent) {
		n.Parents = append(n.Parents, parent)
	}

	// node node a parent childe
	err := parent.AddChild(n)
	if err != nil {
		return errors.New("(graph:AddParent) -> " + err.Error())
	}

	return nil
}

// AddChild method update node's childs list adding a new one
func (n *Node) AddChild(child *Node) error {
	if n == nil {
		return errors.New("(graph::AddChild) Adding child to a nil node")
	}
	if n.Childs == nil || len(n.Childs) == 0 {
		n.Childs = []*Node{}
	}

	if !n.HasChild(child) {
		n.Childs = append(n.Childs, child)
	}

	return nil
}

// HasChild method validate whether a child node already exists in node's child list. Two nodes are equal when they have the same node name
func (n *Node) HasChild(child *Node) bool {
	hasChild := false
	for _, c := range n.Childs {
		if c.Name == child.Name {
			return true
		}
	}

	return hasChild
}

// HasParent method validate whether a parent node already exists in node's parent list. Two nodes are equal when they have the same node name
func (n *Node) HasParent(parent *Node) bool {
	hasParent := false
	for _, p := range n.Parents {
		if p.Name == parent.Name {
			return true
		}
	}

	return hasParent
}