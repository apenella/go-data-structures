package graph

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAddParent
func TestAddParent(t *testing.T) {
	tests := []struct {
		desc   string
		node   *Node
		parent *Node
		err    error
		res    *Node
	}{
		{
			desc: "Add parent node",
			node: &Node{
				Name:   "node",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			parent: &Node{
				Name:   "parent",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			err: nil,
			res: &Node{
				Name:   "node",
				Childs: nil,
				Item:   nil,
				Parent: &Node{
					Name:   "parent",
					Parent: nil,
					Childs: []*Node{
						{
							Name: "node",
						},
					},
					Item: nil,
				},
			},
		},

		{
			desc: "Add parent to nil node",
			node: nil,
			parent: &Node{
				Name:   "parent",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			err: errors.New("(graph::AddParent) Adding parent to a nil node"),
			res: nil,
		},

		{
			desc:   "Add nil parent to node",
			parent: nil,
			node: &Node{
				Name:   "node",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			err: errors.New("(graph:AddParent) -> (graph::AddChild) Adding child to a nil node"),
			res: nil,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.node.AddParent(test.parent)
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, test.res.Name, test.node.Name, "Name not equal")
			assert.Equal(t, test.res.Parent.Name, test.node.Parent.Name, "Parent name not equal")
			assert.Equal(t, len(test.res.Parent.Childs), len(test.node.Parent.Childs), "Parent childs length not equal")
		}
	}
}

// TestAddChild
func TestAddChild(t *testing.T) {
	tests := []struct {
		desc   string
		node   *Node
		parent *Node
		err    error
		res    *Node
	}{
		{
			desc: "Add child to node",
			node: &Node{
				Name:   "node",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			parent: &Node{
				Name:   "parent",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			err: nil,
			res: &Node{
				Name: "parent",
				Childs: []*Node{
					&Node{
						Name:   "node",
						Parent: nil,
						Childs: nil,
						Item:   nil,
					},
				},
				Item:   nil,
				Parent: nil,
			},
		},
		{
			desc: "Add second child to node",
			node: &Node{
				Name:   "node2",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			parent: &Node{
				Name:   "parent",
				Parent: nil,
				Childs: []*Node{
					{
						Name:   "node",
						Parent: nil,
						Childs: nil,
						Item:   nil,
					},
				},
				Item: nil,
			},
			err: nil,
			res: &Node{
				Name: "parent",
				Childs: []*Node{
					{
						Name:   "node",
						Parent: nil,
						Childs: nil,
						Item:   nil,
					},
					{
						Name:   "node2",
						Parent: nil,
						Childs: nil,
						Item:   nil,
					},
				},
				Item:   nil,
				Parent: nil,
			},
		},

		{
			desc:   "Add child to nil parent",
			parent: nil,
			node: &Node{
				Name:   "node",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			err: errors.New("(graph::AddChild) Adding child to a nil node"),
			res: nil,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.parent.AddChild(test.node)
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, test.res, test.parent, "Nodes not equal")
		}
	}
}

func TestHasChild(t *testing.T) {

	tests := []struct {
		desc   string
		node   *Node
		parent *Node
		err    error
		res    bool
	}{
		{
			desc: "Node is not a child",
			node: &Node{
				Name:   "node",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			parent: &Node{
				Name:   "parent",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			err: nil,
			res: false,
		},

		{
			desc: "Node is not a child",
			node: &Node{
				Name:   "node",
				Parent: nil,
				Childs: nil,
				Item:   nil,
			},
			parent: &Node{
				Name:   "parent",
				Parent: nil,
				Childs: []*Node{
					{
						Name:   "node",
						Parent: nil,
						Childs: nil,
						Item:   nil,
					},
				},
				Item: nil,
			},
			err: nil,
			res: true,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		has := test.parent.HasChild(test.node)
		assert.Equal(t, test.res, has, "Nodes not equal")

	}
}
