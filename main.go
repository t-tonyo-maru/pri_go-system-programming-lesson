package main

// =======================================================
// ## 6.6.2 Keep-Alive 対応の HTTP クライアント
// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strings"
// )

// func main() {
// 	sendMessages := []string{
// 		"ASCII",
// 		"PROGRAMMING",
// 		"PLUS",
// 	}
// 	current := 0
// 	var conn net.Conn = nil
// 	// リトライ用にループを全体で囲う
// 	for {
// 		var err error
// 		// まだコネクションを張っていない / エラーでリトライ
// 		if conn == nil {
// 			// Dial から行って conn を初期化
// 			conn, err = net.Dial("tcp", "localhost:8888")
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println("Access: %d\n", current)
// 		}

// 		// POST で文字列を送るリクエストを作成
// 		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
// 		if err != nil {
// 			panic(err)
// 		}

// 		err = request.Write(conn)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
// 		response, err := http.ReadResponse(bufio.NewReader(conn), request)
// 		if err != nil {
// 			fmt.Println("Retry")
// 			conn = nil
// 			continue
// 		}
// 		// 結果を表示
// 		dump, err := httputil.DumpResponse(response, true)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(dump))

// 		// 全部送信完了していれば終了
// 		current++
// 		if current == len(sendMessages) {
// 			break
// 		}
// 	}
// 	conn.Close()
// }

// =======================================================
// ## 6.6.1 Keep-Alive 対応の HTTP サーバー
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strings"
// 	"time"
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
// 			defer conn.Close()
// 			fmt.Printf("Accept %v\n", conn.RemoteAddr())
// 			// Accept 後のソケットで何度も応答を返すためループ
// 			for {
// 				// タイムアウトを設定
// 				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
// 				// リクエストを読み込む
// 				request, err := http.ReadRequest(bufio.NewReader(conn))
// 				if err != nil {
// 					// タイムアウトもしくはソケットクローズ時は終了
// 					// それ以外はエラーにする
// 					neterr, ok := err.(net.Error) // ダウンキャスト
// 					if ok && neterr.Timeout() {
// 						fmt.Println("Timeout")
// 						break
// 					} else if err == io.EOF {
// 						break
// 					}
// 					panic(err)
// 				}

// 				// リクエストを表示
// 				dump, err := httputil.DumpRequest(request, true)
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println(string(dump))
// 				content := "Hello World\n"

// 				// レスポンスを書き込む
// 				// HTTP/1.1 かつ、ContentLength の設定が必要
// 				response := http.Response{
// 					StatusCode:    200,
// 					ProtoMajor:    1,
// 					ProtoMinor:    1,
// 					ContentLength: int64(len(content)),
// 					Body:          io.NopCloser(strings.NewReader(content)),
// 				}
// 				response.Write(conn)
// 			}
// 		}()
// 	}
// }

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
