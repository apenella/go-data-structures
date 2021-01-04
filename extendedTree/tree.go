package gdsexttree

import (
	"fmt"
	"io"

	errors "github.com/apenella/go-common-utils/error"
)

// Graph defines the extended tree structure
type Graph struct {
	Root       []*Node
	NodesIndex map[string]*Node
}

// AddNode method includes a new node on the tree
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
		if n.Parents == nil || len(n.Parents) == 0 {
			g.Root = append(g.Root, n)
		} else {
			for _, parent := range n.Parents {
				parent.AddChild(n)
			}
		}

		g.NodesIndex[n.Name] = n
	}

	return nil
}

// AddRelationship method update the parent-child relationship between two nodes
func (g *Graph) AddRelationship(parent, child *Node) error {
	var exist bool
	var err error
	var p, c *Node

	if g == nil {
		return errors.New("(graph::AddRelationship)", "Graph is null")
	}
	if parent == nil {
		return errors.New("(graph::AddRelationship)", "Parent is null")
	}
	if child == nil {
		return errors.New("(graph::AddRelationship)", "Child is null")
	}
	p, exist = g.NodesIndex[parent.Name]
	if !exist {
		return errors.New("(graph::AddRelationship)", "Parent does not exist")
	}
	c, exist = g.NodesIndex[child.Name]
	if !exist {
		return errors.New("(graph::AddRelationship)", "Child does not exist")
	}

	err = c.AddParent(p)
	if err != nil {
		return errors.New("(graph::AddRelationship)", fmt.Sprintf("Parent can not be added to '%s'", c.Name), err)
	}

	// remove child from root nodes when child node was defined on root nodes
	for i := 0; i < len(g.Root); i++ {
		if g.Root[i].Name == child.Name {
			g.Root[i] = g.Root[len(g.Root)-1]
			g.Root = g.Root[:len(g.Root)-1]
			break
		}
	}

	if hasCyclesRec(p, map[string]int8{}) {
		return errors.New("(graph::AddRelationship)", fmt.Sprintf("Cycle detected adding relationship from '%s' to '%s'", p.Name, c.Name))
	}

	if len(g.Root) < 1 {
		return errors.New("(graph::AddRelationship)", fmt.Sprintf("Relationship from '%s' to '%s' caused an empty list of root nodes", p.Name, c.Name))
	}

	return nil
}

// DrawGraph method prints the graph
func (g *Graph) DrawGraph(w io.Writer) {

	for _, root := range g.Root {
		prefix := "|-> "
		drawGrapRec(w, prefix, root)
	}
}

// drawGraphRec method walks along the tree to draw it
func drawGrapRec(w io.Writer, prefix string, node *Node) {

	fmt.Fprintln(w, prefix, node.Name)
	prefix = "  " + prefix
	for _, child := range node.Childs {
		drawGrapRec(w, prefix, child)
	}
}

// HasCycles returns whether a cyclic dependency exists on the whole graph. It calls hasCyclesRec by each root node
func (g *Graph) HasCycles() bool {
	for _, root := range g.Root {
		if hasCyclesRec(root, map[string]int8{}) {
			return true
		}
	}

	return false
}

// hasCyclesRec returns whether a cyclic dependency
func hasCyclesRec(node *Node, visitedNodes map[string]int8) bool {

	_, exists := visitedNodes[node.Name]
	if exists {
		return true
	}

	visitedNodes[node.Name] = int8(0)

	for _, child := range node.Childs {
		if hasCyclesRec(child, visitedNodes) {
			return true
		}
	}

	return false
}

// Exist return if a node already exists on the graph
func (g *Graph) Exist(n *Node) bool {
	if g == nil || g.NodesIndex == nil {
		return false
	}
	_, exist := g.NodesIndex[n.Name]
	return exist
}

// GetNode method returns the node which matches to the gived name
func (g *Graph) GetNode(nodeName string) (*Node, error) {
	if g == nil {
		return nil, errors.New("(graph::GetNode)", "Graph is nil")
	}

	if g.NodesIndex == nil {
		return nil, errors.New("(graph::GetNode)", "NodesIndex is nil")
	}

	node, exists := g.NodesIndex[nodeName]
	if !exists {
		return nil, errors.New("(graph::GetNode)", "Node '"+nodeName+"' does not exists on the graph")
	}
	return node, nil
}
