package request

import (
	"fmt"
	"strings"
)

type HTTPMethod string

const (
  MethodGet  HTTPMethod = "GET"
  MethodPost HTTPMethod = "POST"
)

type Request struct {
	Method HTTPMethod
	Path string
	Protocol string

	Header map[string][]string
	Body []byte
}

func NewRequest() *Request {
	return &Request{
		Method: "",
		Path: "",
		Protocol: "",
		Header: make(map[string][]string),
		Body: make([]byte, 0),
	}
}

func (r *Request) WithMethod(method HTTPMethod) *Request {
	r.Method = method
	return r
}

func (r *Request) WithPath(path string) *Request {
	r.Path = path
	return r
}

func (r *Request) WithProtocol(protocol string) *Request {
	r.Protocol = protocol
	return r
}

func (r *Request) WithHeader(header map[string][]string) *Request {
	r.Header = header
	return r
}

func (r *Request) AddHeaderField(key string, value string) *Request {
	if _, ok := r.Header[key]; !ok {
		r.Header[key] = []string{value}
	} else {
		r.Header[key] = append(r.Header[key], value)
	}

	return r
}

func (r *Request) WithBody(body []byte) *Request {
	r.Body = body
	return r
}

func (r *Request) String() string {
	reqLine := fmt.Sprintf("%v %s %s", r.Method, r.Path, r.Protocol)

	header := []string{}
	for key, values := range r.Header {
		header = append(header, fmt.Sprintf("%-20s: %s", key, strings.Join(values, ", ")))
	}

	body := string(r.Body)

	return fmt.Sprintf(
		"--- Request Line ---\n%s" +
		"\n\n--- Header ---\n%s" +
		"\n\n--- Body ---\n%s",
		reqLine, strings.Join(header, "\n"), body)
}