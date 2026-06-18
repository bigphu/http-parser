package main

import (
	"fmt"

	"http-parser/request"
	"http-parser/response"
	"http-parser/routing"
	"http-parser/server"
)

func main() {	
	router := routing.NewRouter()
	sv := server.NewServer("localhost", 8080, router)

	router.Get("/test/auth/v1", func(req *request.Request, res *response.Response) {
		fmt.Println("Handler for /test/auth/v1 called")
		res.WithStatus(200).WithString("Auth v1 successful")
	})
	
	router.Get("/user/u1", func(req *request.Request, res *response.Response) {
		fmt.Println("Handler for /user/u1 called")
		res.WithStatus(200).WithString("Hello, User 1!")
	})
	
	router.Post("/post/new", func(req *request.Request, res *response.Response) {
		fmt.Println("Handler for /post/new called")
		res.WithStatus(201).WithString("New post created successfully!")
	})

	router.Get("/settings", func(req *request.Request, res *response.Response) {
		fmt.Println("Handler for /settings called")
		res.WithStatus(200).WithString("Settings page data")
	})
	
	router.Get("/user/create", func(req *request.Request, res *response.Response) {
		fmt.Println("Handler for /user/create called")
		res.WithStatus(200).WithString("User creation form")
	})

	router.Get("/////test/auth/v2", func(req *request.Request, res *response.Response) {
		fmt.Println("Handler for /test/auth/v2 called")
		res.WithStatus(200).WithString("Auth v2 successful")
	})
	
	sv.Start()

	fmt.Println(router.Routes())
}

// vvv DONE vvv
// TODO: thêm một struct Request để đại diện cho request nhận được từ client. Struct này sẽ có 
// các trường như Method, Path, Headers, Body, ... và có thể có các method để parse request 
// từ raw bytes nhận được từ connection.	

// vvv DONE vvv
// TODO: thêm struct Router để register các route và handler tương ứng. Router sẽ sử dụng cây trie 
// để lưu trữ các route và tìm kiếm handler tương ứng khi nhận được request. Router sẽ có method 
// RegisterRoute(path string, handler func()) để đăng ký route và handler, và method 
// FindHandler(path string) func() để tìm kiếm handler tương ứng với đường dẫn của request. 
// 
// Mỗi method HTTP (GET, POST, PUT, DELETE) sẽ có một cây trie riêng để lưu trữ các route 
// và handler tương ứng.
// Router sẽ có một map để lưu trữ các cây trie cho từng method HTTP, ví dụ như map[string]*PathTrie, 
// trong đó key là method HTTP và value là cây trie chứa các route và handler tương ứng với method đó.

// vvv DONE vvv
// TODO: strategy pattern để xủ lý các loại request khác nhau, ví dụ như GET, POST, PUT, DELETE. 
// Mỗi loại request sẽ có một struct riêng implement interface RequestHandler với method HandleRequest. 
// Khi nhận được request, server sẽ xác định loại request và gọi method HandleRequest tương ứng để xử lý.


// vvv DONE vvv
// TODO: thêm Node struce để tạo một cây trie để thực hiện routing cho các request. 
// Mỗi node sẽ lưu trữ một phần của đường dẫn và có thể có nhiều node con để đại diện 
// cho các phần tiếp theo của đường dẫn. Khi nhận được request, server sẽ duyệt cây trie để 
// tìm ra handler tương ứng với đường dẫn của request.

// vvv DONE vvv
// TODO: thêm một response struct sử dụng builder design pattern để xây dựng response. Response struct 
// sẽ có các method để thêm header, body, status code, ... và cuối cùng sẽ có một method Build() 
// để tạo ra response hoàn chỉnh.