package gdstree

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
)

// Node
type Node struct {
	Name     string
	Parent   *Node
	Item     interface{}
	Children []*Node
}

// AddChild
func (n *Node) AddParent(parent *Node) error {
	if n == nil {
		return errors.New("(graph::AddParent)", "Adding parent to a nil node")
	}

	if parent == nil {
		return errors.New("(graph::AddParent)", "Parent is nil")
	}

	if n.Parent != nil {
		return errors.New("(graph::AddParent)", fmt.Sprintf("Node '%s' is already defined", n.Name))
	}

	n.Parent = parent
	err := parent.AddChild(n)
	if err != nil {
		return errors.New("(graph:AddParent)", fmt.Sprintf("Child could not be add to '%s'", parent.Name), err)
	}

	return nil
}

// AddChild
func (n *Node) AddChild(child *Node) error {
	if n == nil {
		return errors.New("(graph::AddChild)", "Adding child to a nil node")
	}

	if child == nil {
		return errors.New("(graph::AddChild)", "Child is nil")
	}

	if n.Children == nil || len(n.Children) == 0 {
		n.Children = []*Node{}
	}

	if !n.HasChild(child) {
		n.Children = append(n.Children, child)
	}

	return nil
}

func (n *Node) HasChild(child *Node) bool {
	hasChild := false
	for _, c := range n.Children {
		if c.Name == child.Name {
			return true
		}
	}

	return hasChild
}
