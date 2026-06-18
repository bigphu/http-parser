package httpparser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseBuffer(reader *bufio.Reader) (*Request, error) {
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Bỏ '\r' ở cuối dòng
	requestLine = strings.TrimSpace(requestLine)

	rlParts := strings.Split(requestLine, " ")
	if len(rlParts) < 3 {
		return nil, fmt.Errorf("Malformed request line: %s", requestLine)
	}
	
	request := NewRequest().
							WithMethod(HTTPMethod(rlParts[0])).
							WithPath(rlParts[1]).
							WithProtocol(rlParts[2])

	contentLength := 0

	// TODO: Đọc phần còn lại của header
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)
		// "\r\n" -> "": báo hiệu kết thúc phần header, chuyển sang đọc body
		if (line == "") {
			break
		}

		lineParts := strings.SplitN(line, ":", 2)
		if len(lineParts) == 2 {
			key := lineParts[0]
			val := strings.TrimSpace(lineParts[1])

			if key == "Content-Length" {
				contentLength, _ = strconv.Atoi(val)
			}

			request.AddHeaderField(key, val)
		}
	}

	// TODO: Đọc body
	if contentLength > 0 {
		body := make([]byte, contentLength)

		_, err := io.ReadFull(reader, body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read request body: %s", err)
		}

		request.WithBody(body)
	}

	return request, nil
}