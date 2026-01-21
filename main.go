package main

// =======================================================
// ## 6.5.2 TCP ソケットを使った HTTP クライアント
// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}

// 	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	request.Write(conn)
// 	response, err := http.ReadResponse(bufio.NewReader(conn), request)
// 	if err != nil {
// 		panic(err)
// 	}

// 	dump, err := httputil.DumpResponse(response, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(dump))
// }

// =======================================================
// ## 6.5.1 TCP ソケットを使った HTTP サーバー
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strings"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Server is running at localhost:8888")

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			panic(err)
// 		}
// 		go func() {
// 			fmt.Printf("Accept %v\n", conn.RemoteAddr())
// 			// リクエストを読み込む
// 			request, err := http.ReadRequest(bufio.NewReader(conn))
// 			if err != nil {
// 				panic(err)
// 			}
// 			dump, err := httputil.DumpRequest(request, true)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(string(dump))
// 			response := http.Response{
// 				StatusCode: 200,
// 				ProtoMajor: 1,
// 				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
// 			}
// 			response.Write(conn)
// 			conn.Close()
// 		}()
// 	}
// }
