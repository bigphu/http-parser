package routing

import (
	"fmt"
	pathpkg "http-parser/path"
	"http-parser/request"
	"http-parser/response"
	"strings"
)

type Router struct {
	routes map[request.HTTPMethod]*pathpkg.PathTrie
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[request.HTTPMethod]*pathpkg.PathTrie),
	}
}

func (r *Router) Get(path string, handler func(*request.Request, *response.Response)) {
	if _, ok := r.routes[request.MethodGet]; !ok {
		r.routes[request.MethodGet] = pathpkg.NewPathTrie()
	}

	r.routes[request.MethodGet].Insert(path, handler)
}

func (r *Router) Post(path string, handler func(*request.Request, *response.Response)) {
	if _, ok := r.routes[request.MethodPost]; !ok {
		r.routes[request.MethodPost] = pathpkg.NewPathTrie()
	}

	r.routes[request.MethodPost].Insert(path, handler)
}

func (r *Router) Invoke(req *request.Request, res *response.Response) {
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