package graph

import (
	"errors"
)

// Node
type Node struct {
	Name   string
	Parent *Node
	Item   interface{}
	Childs []*Node
}

// AddChild
func (n *Node) AddParent(parent *Node) error {
	if n == nil {
		return errors.New("(graph::AddParent) Adding parent to a nil node")
	}

	n.Parent = parent
	err := parent.AddChild(n)
	if err != nil {
		return errors.New("(graph:AddParent) -> " + err.Error())
	}

	return nil
}

// AddChild
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

func (n *Node) HasChild(child *Node) bool {
	hasChild := false
	for _, c := range n.Childs {
		if c.Name == child.Name {
			return true
		}
	}

	return hasChild
}
