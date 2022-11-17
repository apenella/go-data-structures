package gdstree

import (
	"fmt"
	"io"

	errors "github.com/apenella/go-common-utils/error"
)

// Graph
type Graph struct {
	Root       []*Node
	NodesIndex map[string]*Node
}

// AddNode
func (g *Graph) AddNode(n *Node) error {

	if g == nil {
		return errors.New("(graph::AddNode)", "Adding a node to a nil graph")
	}
	// there is no nodes on the graph
	if g.NodesIndex == nil || len(g.NodesIndex) == 0 {
		g.NodesIndex = map[string]*Node{
			n.Name: n,
		}
		g.Root = []*Node{n}
	} else {
		_, ok := g.NodesIndex[n.Name]
		if ok {
			return errors.New("(graph::AddNode)", "Node '"+n.Name+"' already exists on the graph")
		}
		// add node to the graph
		if n.Parent == nil {
			g.Root = append(g.Root, n)
		} else {
			n.Parent.AddChild(n)
		}
		g.NodesIndex[n.Name] = n
	}

	return nil
}

// AddRelationship
func (g *Graph) AddRelationship(parent, child *Node) error {
	var exist bool

	if g == nil {
		return errors.New("(graph::AddParentToNode)", "Graph is null")
	}
	if parent == nil {
		return errors.New("(graph::AddParentToNode)", "Parent is null")
	}
	if child == nil {
		return errors.New("(graph::AddParentToNode)", "Child is null")
	}
	_, exist = g.NodesIndex[parent.Name]
	if !exist {
		return errors.New("(graph::AddParentToNode)", "Parent does not exist")
	}
	_, exist = g.NodesIndex[child.Name]
	if !exist {
		return errors.New("(graph::AddParentToNode)", "Child does not exist")
	}

	child.AddParent(parent)

	for i := 0; i < len(g.Root); i++ {
		if g.Root[i].Name == child.Name {
			g.Root[i] = g.Root[len(g.Root)-1]
			g.Root = g.Root[:len(g.Root)-1]
			break
		}
	}

	return nil
}

// DrawGraph
func (g *Graph) DrawGraph(w io.Writer) {

	for _, root := range g.Root {
		prefix := "\u251C\u2500\u2500\u2500"
		drawGrapRec(w, prefix, root)
	}
}

// drawGraphRec
func drawGrapRec(w io.Writer, prefix string, node *Node) {

	fmt.Fprintln(w, prefix, node.Name)
	prefix = "\u2502  " + prefix
	for _, child := range node.Children {
		drawGrapRec(w, prefix, child)
	}
}

func (g *Graph) Exist(n *Node) bool {
	if g == nil || g.NodesIndex == nil {
		return false
	}
	_, exist := g.NodesIndex[n.Name]
	return exist
}

func (g *Graph) GetNode(n string) (*Node, error) {
	if g == nil {
		return nil, errors.New("(graph::GetNode)", "Graph is nil")
	}

	if g.NodesIndex == nil {
		return nil, errors.New("(graph::GetNode)", "NodesIndex is nil")
	}

	node, exists := g.NodesIndex[n]
	if !exists {
		return nil, errors.New("(graph::GetNode)", "Node '"+n+"' does not exists on the graph")
	}
	return node, nil
}
