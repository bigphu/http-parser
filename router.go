package httpparser

import (
	"fmt"
	"strings"
)

type Router struct {
	routes map[HTTPMethod]*PathTrie
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[HTTPMethod]*PathTrie),
	}
}

func (r *Router) Get(path string, handler func(*Request, *Response)) {
	if _, ok := r.routes[MethodGet]; !ok {
		r.routes[MethodGet] = NewPathTrie()
	}

	r.routes[MethodGet].Insert(path, handler)
}

func (r *Router) Post(path string, handler func(*Request, *Response)) {
	if _, ok := r.routes[MethodPost]; !ok {
		r.routes[MethodPost] = NewPathTrie()
	}

	r.routes[MethodPost].Insert(path, handler)
}

func (r *Router) Invoke(req *Request, res *Response) {
	if trie, ok := r.routes[req.Method]; ok {
		handler := trie.Search(req.Path)
		if handler != nil {
			handler(req, res)
		} else {
			res.WithStatus(404).WithString("404 Not Found: " + req.Path)
		}
		
	} else {
		res.WithStatus(405).WithString("Nuh uh :/")
	}
}

func (r *Router) Routes() string {
	var routes []string
	for method, trie := range r.routes {
		route := fmt.Sprintf("%s: %s", method, trie.String())
		routes = append(routes, route)
	}

	return strings.Join(routes, "\n\n---\n\n")
}