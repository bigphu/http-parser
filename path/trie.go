package path

type PathTrie struct {
	root *Node
}

func NewPathTrie() *PathTrie {
	return &PathTrie{
		root: NewNode(),
	}
}

func (pt *PathTrie) Insert(path string, handler func()) {
	pt.root.Insert(path, handler)
}

func (pt *PathTrie) Search(path string) func() {
	return pt.root.Search(path)
}

func (pt *PathTrie) String() string {
	return pt.root.String("")
}

func (pt *PathTrie) Invoke(path string) {
	handler := pt.root.Search(path)
	if handler != nil {
		handler()
	}
}
