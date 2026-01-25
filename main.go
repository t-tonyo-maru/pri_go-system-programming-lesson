package main

func main() {

}

// =======================================================
// ## 8.2.4 データグラム方の Unix ドメインソケット
// import (
// 	"fmt"
// 	"net"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
// 	// エラーチェックは削除（存在しなかったらしかなったで問題ないので不要）
// 	os.Remove(path)
// 	fmt.Println("Server is running at " + path)

// 	conn, err := net.ListenPacket("unixgram", path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	buffer := make([]byte, 1500)
// 	for {
// 		length, remoteAddress, err := conn.ReadFrom(buffer)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("Revieved from %v: %v\n", remoteAddress, string(buffer[:length]))
// 		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }

// =======================================================
// ## 8.2.3 Unix ドメインソケット版の HTTP クライアントを作る
// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	conn, err := net.Dial("unix", filepath.Join(os.TempDir(), "unixdomainsocket-sample"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	request, err := http.NewRequest("get", "http;//localhost:8888", nil)
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
// ## 8.2.2 Unix ドメインソケット版の HTTP サーバーを作る
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func main() {
// 	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")
// 	os.Remove(path)

// 	listener, err := net.Listen("unix", path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server is running at " + path)

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

// 			// レスポンスを書き込む
// 			response := http.Response{
// 				StatusCode: 200,
// 				ProtoMajor: 1,
// 				ProtoMinor: 0,
// 				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
// 			}
// 			response.Write(conn)
// 			conn.Close()
// 		}()
// 	}
// }
