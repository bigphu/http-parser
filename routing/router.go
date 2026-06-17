package routing

import (
	"http-parser/path"
)

type Router struct {
	pathTrie *path.PathTrie
	
}