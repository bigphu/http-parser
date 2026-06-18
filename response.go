package httpparser

import (
	"fmt"
	"slices"
	"strings"
)

type Response struct {
	Protocol string
	StatusCode int
	Desc string

	Header map[string][]string
	Body []byte
}

func NewResponse() *Response {
	return &Response{
		Protocol: "HTTP/1.1",
		StatusCode: 200,
		Desc: "OK",
		Header: make(map[string][]string),
		Body: make([]byte, 0),
	}
}

var httpStatusCodes = map[int]string {
	100: "Continue",
	101: "Switching Protocols",
	103: "Early Hints",
	200: "OK",
	201: "Created",
	204: "No Content",
	206: "Partial Content",
	301: "Moved Permanently",
	302: "Found",
	304: "Not Modified",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	408: "Request Timeout",
	429: "Too Many Requests",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
}

func (r *Response) WithProtocol(protocol string) *Response {
	r.Protocol = protocol
	return r
}

func (r *Response) WithStatus(code int) *Response {
	desc, ok := httpStatusCodes[code]
	if !ok {
		fmt.Printf("Unknown status code: %d", code)
	} else {
		r.StatusCode = code
		r.Desc = desc
	}
	return r
}

func (r *Response) AddHeaderField(key string, value string) *Response {
	if _, ok := r.Header[key]; !ok {
		r.Header[key] = []string{value}
	} else {
		r.Header[key] = append(r.Header[key], value)
	}

	return r
}

func (r *Response) WithString(body string) *Response {
	r.AddHeaderField("Content-Type", "text/plain; charset=utf-8")
	r.AddHeaderField("Content-Length", fmt.Sprint(len([]byte(body))))
	r.Body = []byte(body)
	return r
}

func (r *Response) WithBytes(body []byte) *Response {
	r.AddHeaderField("Content-Type", "application/octet-stream")
	r.AddHeaderField("Content-Length", fmt.Sprintf("%d", len(body)))
	r.Body = []byte(body)	
	return r
}

func (r *Response) String() string {
	staLine := fmt.Sprintf("%s %v %s", r.Protocol, r.StatusCode, r.Desc)

	header := []string{}
	for key, values := range r.Header {
		header = append(header, fmt.Sprintf("%s: %s", key, strings.Join(values, ", ")))
	}

	body := string(r.Body)

	return fmt.Sprintf(
		"--- Response Line ---\n%s" +
		"\n\n--- Header ---\n%s" +
		"\n\n--- Body ---\n%s",
		staLine, strings.Join(header, "\n"), body)
}

func (r *Response) Build() []byte {
	staLine := []byte(fmt.Sprintf("%s %v %s\r\n", r.Protocol, r.StatusCode, r.Desc))

	header := []string{}
	for key, values := range r.Header {
		header = append(header, fmt.Sprintf("%s: %s", key, strings.Join(values, ", ")))
	}
	header = append(header, "\r\n")

	hBytes := []byte(strings.Join(header, "\r\n"))
	body := r.Body

	return slices.Concat(staLine, hBytes, body)
}