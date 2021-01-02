package gdsexttree

import (
	"bytes"
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

// TestLoadImage tests
func TestAddNode(t *testing.T) {

	node := &Node{
		Name:    "node",
		Parents: nil,
		Item:    "string item",
		Childs:  nil,
	}

	node2 := &Node{
		Name: "node2",
		Parents: []*Node{
			node,
		},
		Item:   "string item",
		Childs: nil,
	}

	node3 := &Node{
		Name:    "node3",
		Parents: nil,
		Item:    "string item",
		Childs:  nil,
	}

	tests := []struct {
		desc  string
		graph *Graph
		node  *Node
		err   error
		res   *Graph
	}{
		{
			desc:  "Add Node to nil graph",
			graph: nil,
			node:  node,
			err:   errors.New("(graph::AddNode)", "Adding a node to a nil graph"),
			res:   nil,
		},
		{
			desc:  "Add Node to an empty graph",
			graph: &Graph{},
			node:  node,
			err:   nil,
			res: &Graph{
				Root: []*Node{
					node,
				},
				NodesIndex: map[string]*Node{
					"node": node,
				},
			},
		},
		{
			desc: "Add Node to an empty graph with and empty index",
			graph: &Graph{
				NodesIndex: map[string]*Node{},
			},
			node: node,
			err:  nil,
			res: &Graph{
				Root: []*Node{
					node,
				},
				NodesIndex: map[string]*Node{
					"node": node,
				},
			},
		},
		{
			desc: "Add an existing node to a graph",
			graph: &Graph{
				Root: []*Node{
					{
						Name: "root",
					},
				},
				NodesIndex: map[string]*Node{
					"root": {
						Name: "root",
					},
				},
			},
			node: &Node{
				Name: "root",
			},
			err: errors.New("(graph::AddNode)", "Node 'root' already exists on the graph"),
			res: nil,
		},
		{
			desc: "Add Node as a child to another node",
			graph: &Graph{
				Root: []*Node{
					node,
				},
				NodesIndex: map[string]*Node{
					"node": node,
				},
			},
			node: node2,
			err:  nil,
			res: &Graph{
				Root: []*Node{
					{
						Name:    "node",
						Parents: nil,
						Item:    "string item",
						Childs: []*Node{
							{
								Name: "node2",
								Parents: []*Node{
									node,
								},
								Item:   "string item",
								Childs: nil,
							},
						},
					},
				},
				NodesIndex: map[string]*Node{
					"node":  node,
					"node2": node2,
				},
			},
		},
		{
			desc: "Add Node without parent to a graph with one element",
			graph: &Graph{
				Root: []*Node{
					node,
				},
				NodesIndex: map[string]*Node{
					"node": node,
				},
			},
			node: node3,
			err:  nil,
			res: &Graph{
				Root: []*Node{
					{
						Name:    "node",
						Parents: nil,
						Item:    "string item",
						Childs:  nil,
					},
					{
						Name:    "node3",
						Parents: nil,
						Item:    "string item",
						Childs:  nil,
					},
				},
				NodesIndex: map[string]*Node{
					"node":  node,
					"node2": node3,
				},
			},
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.graph.AddNode(test.node)
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, len(test.graph.Root), len(test.res.Root), "Root size not equal")
			assert.Equal(t, len(test.graph.NodesIndex), len(test.res.NodesIndex), "Root size not equal")
		}
	}
}

// TestDrawGraph
func TestDrawGraph(t *testing.T) {
	tests := []struct {
		desc  string
		graph *Graph
		err   error
		res   string
	}{
		{
			desc: "Print one root",
			res: `|->  root
  |->  level1-sibling1
    |->  level2-sibling1
  |->  level1-sibling2
    |->  level2-sibling1
    |->  level2-sibling2
`,
			graph: &Graph{
				Root: []*Node{
					{
						Name: "root",
						Childs: []*Node{
							{
								Name: "level1-sibling1",
								Childs: []*Node{
									{
										Name:   "level2-sibling1",
										Childs: nil,
									},
								},
							},
							{
								Name: "level1-sibling2",
								Childs: []*Node{
									{
										Name:   "level2-sibling1",
										Childs: nil,
									},
									{
										Name:   "level2-sibling2",
										Childs: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	var w bytes.Buffer
	for _, test := range tests {
		t.Log(test.desc)

		test.graph.DrawGraph(&w)
		assert.Equal(t, test.res, w.String(), "Output not equal")
	}
}

// TestAddParentToNode
func TestAddRelationship(t *testing.T) {

	tests := []struct {
		desc   string
		graph  *Graph
		parent *Node
		node   *Node
		err    error
		res    *Graph
	}{
		{
			desc:   "Add parent to node into a nil graph",
			graph:  nil,
			parent: nil,
			node:   nil,
			err:    errors.New("(graph::AddParentToNode)", "Graph is null"),
			res:    nil,
		},
		{
			desc:   "Add node to a nil parent",
			graph:  &Graph{},
			parent: nil,
			node:   nil,
			err:    errors.New("(graph::AddParentToNode)", "Parent is null"),
			res:    nil,
		},
		{
			desc: "Add nil node to a parent",
			graph: &Graph{
				Root: []*Node{
					{
						Name: "root",
					},
				},
				NodesIndex: map[string]*Node{
					"root": {
						Name: "root",
					},
				},
			},
			parent: &Node{
				Name: "root",
			},
			node: nil,
			err:  errors.New("(graph::AddParentToNode)", "Child is null"),
			res:  nil,
		},
		{
			desc: "Add parent to orphan node",
			graph: &Graph{
				Root: []*Node{
					{
						Name: "root",
					},
					{
						Name:    "orphan",
						Parents: nil,
					},
				},
				NodesIndex: map[string]*Node{
					"root": {
						Name: "root",
					},
					"orphan": {
						Name:    "orphan",
						Parents: nil,
					},
				},
			},
			parent: &Node{
				Name: "root",
			},
			node: &Node{
				Name:    "orphan",
				Parents: nil,
			},
			err: nil,
			res: &Graph{
				Root: []*Node{
					{
						Name: "root",
					},
				},
				NodesIndex: map[string]*Node{
					"root": {
						Name: "root",
					},
					"orphan": {
						Name:    "orphan",
						Parents: nil,
					},
				},
			},
		},
		{
			desc: "Add parent unexistent parent to a node",
			graph: &Graph{
				Root: []*Node{
					{
						Name: "root",
					},
					{
						Name: "orphan",
					},
				},
				NodesIndex: map[string]*Node{
					"root": {
						Name: "root",
					},
					"orphan": {
						Name: "orphan",
					},
				},
			},
			parent: &Node{
				Name: "unexistent",
			},
			node: &Node{
				Name: "orphan",
			},
			err: errors.New("(graph::AddParentToNode)", "Parent does not exist"),
			res: nil,
		},
		{
			desc: "Add parent parent to an unexistent node",
			graph: &Graph{
				Root: []*Node{
					{
						Name: "root",
					},
				},
				NodesIndex: map[string]*Node{
					"root": {
						Name: "root",
					},
				},
			},
			parent: &Node{
				Name: "root",
			},
			node: &Node{
				Name: "unexistent",
			},
			err: errors.New("(graph::AddParentToNode)", "Child does not exist"),
			res: nil,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.graph.AddRelationship(test.parent, test.node)
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, len(test.graph.Root), len(test.res.Root), "Root size not equal")
			assert.Equal(t, len(test.graph.NodesIndex), len(test.res.NodesIndex), "Root size not equal")
		}
	}
}

func TestExist(t *testing.T) {
	tests := []struct {
		desc  string
		graph *Graph
		node  *Node
		res   bool
	}{
		{
			desc: "Search an existing node",
			graph: &Graph{
				NodesIndex: map[string]*Node{
					"node": {
						Name: "node",
					},
				},
			},
			node: &Node{
				Name: "node",
			},
			res: true,
		},
		{
			desc: "Search an unexisting node",
			graph: &Graph{
				NodesIndex: map[string]*Node{
					"node": {
						Name: "node",
					},
				},
			},
			node: &Node{
				Name: "node2",
			},
			res: false,
		},
		{
			desc:  "Search on a nil graph",
			graph: nil,
			node: &Node{
				Name: "node",
			},
			res: false,
		},
		{
			desc: "Search on a nil nodesindex",
			graph: &Graph{
				NodesIndex: nil,
			},
			node: &Node{
				Name: "node",
			},
			res: false,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		exists := test.graph.Exist(test.node)
		assert.Equal(t, exists, test.res, "Unexpected return of existence")

	}
}

func TestGetNode(t *testing.T) {
	tests := []struct {
		desc  string
		graph *Graph
		node  string
		res   *Node
		err   error
	}{
		{
			desc: "Get an existing node",
			graph: &Graph{
				NodesIndex: map[string]*Node{
					"node": {
						Name: "node",
					},
				},
			},
			node: "node",
			res: &Node{
				Name: "node",
			},
			err: nil,
		},
		{
			desc: "Get an unexisting node",
			graph: &Graph{
				NodesIndex: map[string]*Node{
					"node": {
						Name: "node",
					},
				},
			},
			node: "node2",
			res:  nil,
			err:  errors.New("(graph::GetNode)", "Node 'node2' does not exists on the graph"),
		},
		{
			desc:  "Search on a nil graph",
			graph: nil,
			node:  "node",
			res:   nil,
			err:   errors.New("(graph::GetNode)", "Graph is nil"),
		},
		{
			desc: "Search on a nil nodesindex",
			graph: &Graph{
				NodesIndex: nil,
			},
			node: "node",
			res:  nil,
			err:  errors.New("(graph::GetNode)", "NodesIndex is nil"),
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		node, err := test.graph.GetNode(test.node)
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, test.res, node, "Unexpected node")
		}
	}
}
