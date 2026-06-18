package httpparser

type PathTrie struct {
	root *Node
}

func NewPathTrie() *PathTrie {
	return &PathTrie{
		root: NewNode(),
	}
}

func (pt *PathTrie) Insert(path string, handler func(*Request, *Response)) {
	pt.root.Insert(path, handler)
}

func (pt *PathTrie) Search(path string) func(*Request, *Response) {
	return pt.root.Search(path)
}

func (pt *PathTrie) String() string {
	return pt.root.String("")
}

// func (pt *PathTrie) Invoke(path string, req *request.Request, res *response.Response) {
// 	handler := pt.root.Search(path)
// 	if handler != nil {
// 		handler(req, res)
// 	} else {
// 		fmt.Printf("No handler found for path: %s", path)
// 		res.WithStatus(404).WithString("404 Not Found: " + path)
// 	}
// }
