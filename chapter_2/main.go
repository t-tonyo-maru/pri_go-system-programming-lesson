package main

func main() {

}

// =======================================================
// ## 2.4.6 io.Writer のデコレータ

// =======================================================
// ## 2.4.5 インターネット接続
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
// 	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
// 	io.Copy(os.Stdout, conn)
// }
// =======================================================
// import (
// 	"io"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "hhtp.ResponseWriter Sample")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// =======================================================
// ## 2.4.4 書かれた内容を記憶しておくバッファ(2)
// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var builder strings.Builder
// 	builder.Write([]byte("strings.Builder example\n"))
// 	fmt.Println(builder.String())
// }

// =======================================================
// ## 2.4.3 書かれた内容を記憶しておくバッファ(1)
// import (
// 	"bytes"
// 	"fmt"
// )

// func main() {
// 	var buffer bytes.Buffer
// 	buffer.Write([]byte("bytes.Buffer example\n"))
// 	fmt.Println(buffer.String())
// }

// =======================================================
// ## 2.4.2 画面出力
// import "os"
// func main() {
// 	os.Stdout.Write([]byte("os.Stdout example\n"))
// }

// =======================================================
// ## 2.4.1 ファイル出力

// func main() {
// 	file, err := os.Create("text.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	file.Write([]byte("os.File example\n"))
// 	file.Close()
// }

// =======================================================
// ## 2.2 Go 言語のインターフェース

// type Talker interface {
// 	Talk()
// }

// type Greeter struct {
// 	name string
// }

// func (g Greeter) Talk() {
// 	fmt.Printf("Hello, my name is %s\n", g.name)
// }

// func main() {
// 	var talker Talker
// 	talker = Greeter{"wozzo"}
// 	talker.Talk()
// }
