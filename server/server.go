package server

import (
	"bufio"
	"fmt"
	"http-parser/request"
	"http-parser/response"
	"http-parser/routing"
	"net"
	"sync"
)

type Server struct {
	host string
	port int
	router *routing.Router
	listener net.Listener
	wg sync.WaitGroup
}

func NewServer(host string, port int, router *routing.Router) *Server {
	if host == "" {
		host = "localhost"
	}

	if port == 0 {
		port = 8080
	}

	if router == nil {
		router = routing.NewRouter()
	}

	return &Server{
		host: host,
		port: port,
		router: router,
		listener: nil,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer s.Stop()

	fmt.Printf("Server is listening on %s:%d\n", s.host, s.port)
	s.listener = listener

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *Server) Stop() {
	if s.listener != nil {
		s.listener.Close()
	}

	s.wg.Wait()
	fmt.Println("Server stopped successfully.")
}

func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()
	
	reader := bufio.NewReader(conn)

	for {
		req, err := request.ParseBuffer(reader)
		if err != nil {
			break;
		}

		fmt.Println(req.String())

		res := response.NewResponse()
		s.router.Invoke(req, res)


		conn.Write(res.Build())
	}

}