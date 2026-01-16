package main

// =======================================================
// ## 3.4.2 ファイル入力
// import (
// 	"bufio"
// 	"io"
// 	"net"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "example.com:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
// 	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
// 	defer res.Body.Close()
// 	io.Copy(os.Stdout, res.Body)
// }

// =======================================================
// import (
// 	"io"
// 	"net"
// 	"os"
// )
// func main() {
// 	conn, err := net.Dial("tcp", "example.com:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
// 	io.Copy(os.Stdout, conn)
// }

// =======================================================
// ## 3.4.2 ファイル入力
// import (
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("file.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	io.Copy(os.Stdout, file)
// }

// =======================================================
// ## 3.3.2 入出力関連インタフェースのキャスト
// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	for {
// 		buffer := make([]byte, 5)
// 		size, err := os.Stdin.Read(buffer)
// 		if err == io.EOF {
// 			fmt.Println("EOF")
// 			break
// 		}
// 		fmt.Printf("size=%d inputs=%s\n", size, string(buffer))
// 	}
// }
