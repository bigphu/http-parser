package main

import (
	"bytes"
	"fmt"
	"io"
	"net"

	"http-parser/path"
)

func main() {
	// server, err := net.Listen("tcp", "localhost:8080")
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// 	return
	// }
	// defer server.Close()

	// fmt.Println("Server is listening on localhost:8080")

	// for {
	// 	conn, err := server.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error accepting connection:", err)
	// 		continue
	// 	}

	// 	go handleConnection(conn)
	// }

	pathTrie := path.NewPathTrie()

	pathTrie.Insert("/test/auth/v1", func() {
		fmt.Println("Handler for /test/auth/v1 called")
	})
	
	pathTrie.Insert("/user/u1", func() {
		fmt.Println("Handler for /user/u1 called")
	})
	
	pathTrie.Insert("/test/auth/v3", func() {
		fmt.Println("Handler for /test/auth/v3 called")
	})

	pathTrie.Insert("/settings", func() {
		fmt.Println("Handler for /settings called")
	})
	
	pathTrie.Insert("/user/create", func() {
		fmt.Println("Handler for /user/create called")
	})

	pathTrie.Insert("/test/auth/v2", func() {
		fmt.Println("Handler for /test/auth/v2 called")
	})

	fmt.Println(pathTrie.String())

	pathTrie.Invoke("/test/auth/v1")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	buf := make([]byte, 4096)

	_, rerr := conn.Read(buf)
	if rerr != nil && rerr != io.EOF {
		fmt.Println("Error reading request:", rerr)
		return
	}

	parseRequest(buf)
	
	body := `{"status":"ok"}`
	response := fmt.Sprintf(
								"HTTP/1.1 200 OK\r\n" +
            		"Content-Type: application/json\r\n" +
            		"Content-Length: %d\r\n" +
            		"\r\n" +
								"%s",
								len(body), body)

	_, werr := conn.Write([]byte(response))
	if werr != nil {
		fmt.Println("Error writing response:", werr)
		return
	}
	
	// fmt.Printf("---\n\nResponse: \n%s\n", response)
}

func parseRequest(request []byte) {
	header, body, _ := bytes.Cut(request, []byte("\r\n\r\n"))
	parseHeader(header)
	parseBody(body)
}

func parseHeader(header []byte) {
	requestLine, headerBlock, _ := bytes.Cut(header, []byte("\r\n"))

	requestLineParts := bytes.Split(requestLine, []byte(" "))
	if len(requestLineParts) != 3 {
		fmt.Println("Invalid request line:", string(requestLine))
		return
	}
	method := string(requestLineParts[0])
	path := string(requestLineParts[1])
	protocol := string(requestLineParts[2])

	fmt.Printf("---\n\nMethod: %s, Path: %s, Protocol: %s\n", method, path, protocol)

	headerBlockLines := bytes.Split(headerBlock, []byte("\r\n"))
	
	headers := make(map[string]string)
	for _, line := range headerBlockLines {
		parts := bytes.Split(line, []byte(": "))
		headers[string(parts[0])] = string(parts[1])
	}

	fmt.Println("---\n\nHeaders:")
	for key, value := range headers {
		fmt.Printf("  %-20s: %s\n", key, value)
	}
}

func parseBody(body []byte) {
	fmt.Println("---\n\nParsing body:\n\n", string(body))
}

// TODO: thêm struct Router để register các route và handler tương ứng. Router sẽ sử dụng cây trie 
// để lưu trữ các route và tìm kiếm handler tương ứng khi nhận được request. 

// TODO: strategy pattern để xủ lý các loại request khác nhau, ví dụ như GET, POST, PUT, DELETE. 
// Mỗi loại request sẽ có một struct riêng implement interface RequestHandler với method HandleRequest. 
// Khi nhận được request, server sẽ xác định loại request và gọi method HandleRequest tương ứng để xử lý.

// vvv DONE vvv
// TODO: thêm Node struce để tạo một cây trie để thực hiện routing cho các request. 
// Mỗi node sẽ lưu trữ một phần của đường dẫn và có thể có nhiều node con để đại diện 
// cho các phần tiếp theo của đường dẫn. Khi nhận được request, server sẽ duyệt cây trie để 
// tìm ra handler tương ứng với đường dẫn của request.

// TODO: thêm một response struct sử dụng builder design pattern để xây dựng response. Response struct 
// sẽ có các method để thêm header, body, status code, ... và cuối cùng sẽ có một method Build() 
// để tạo ra response hoàn chỉnh.