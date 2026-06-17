package path

import (
	"strings"
)

type Node struct {
	children map[string]*Node
	handler  func()
}

func NewNode() *Node {
	return &Node{
		children: make(map[string]*Node),
		handler:  nil,
	}
}

func splitPath(path string) []string {
	// Remove leading and trailing slashes
	if len(path) > 0 && path[0] == '/' {
		return splitPath(path[1:])
	}

	if len(path) > 0 && path[len(path)-1] == '/' {
		return splitPath(path[:len(path)-1])
	}

	if path == "" {
		return []string{}
	}

	return strings.Split(path, "/")
}

func (n *Node) Insert(path string, handler func()) {
	parts := splitPath(path)
	curr := n
	for _, part := range parts {
		if _, ok := curr.children[part]; !ok {
			curr.children[part] = NewNode()
		}
		curr = curr.children[part]
	}
	curr.handler = handler
}

func (n *Node) Search(path string) func() {
	parts := splitPath(path)
	curr := n
	for _, part := range parts {
		if _, ok := curr.children[part]; !ok {
			return nil
		}
		curr = curr.children[part]
	}
	return curr.handler
}

func (n *Node) String(prefix string) string {
	if n.handler != nil {
		return prefix + " (handler)"
	}
	var paths []string
	for part, child := range n.children {
		childPath := child.String(prefix + "/" + part)
		paths = append(paths, childPath)
	}
	
	return strings.Join(paths, "\n")
}