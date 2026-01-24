package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}

	var conn net.Conn = nil
	var err error
	requests := make([]*http.Request, 0, len(sendMessages))

	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access\n")
	defer conn.Close()

	// リクエストだけ先に送る
	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest(
			"GET",
			"http://localhost:8888?message="+sendMessages[i],
			nil,
		)
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessages[i])
		requests = append(requests, request)
	}

	// レスポンスをまとめて受信
	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}

// =======================================================
// ## 6.9.2 パイプライニングのクライアント実装

// =======================================================
// ## 6.9.1 パイプライニングのサーバー実装
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

// func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
// 	defer conn.Close()
// 	// 順番に取り出す
// 	for sessionResponse := range sessionResponses {
// 		// 選択された仕事が終わるまで待つ
// 		response := <-sessionResponse
// 		response.Write(conn)
// 		close(sessionResponse)
// 	}
// }

// // セッション内のリクエストを取得する
// func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
// 	dump, err := httputil.DumpRequest(request, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(dump))

// 	content := "Hello World\n"
// 	// レスポンスを書き込む
// 	// セッションを維持するために Keep-Alive でないといけない
// 	response := &http.Response{
// 		StatusCode:    200,
// 		ProtoMajor:    1,
// 		ProtoMinor:    1,
// 		ContentLength: int64(len(content)),
// 		Body:          io.NopCloser(strings.NewReader(content)),
// 	}
// 	// 処理が終わったらチャネルに書き込み、
// 	// ブロックされていた writeToConn の処理を再始動する
// 	resultReceiver <- response
// }

// // セッション1つを処理
// func processSession(conn net.Conn) {
// 	fmt.Printf("Accept %v\n", conn.RemoteAddr())
// 	// セッション内のリクエストを順に処理するためのチャネル
// 	sessionRespones := make(chan chan *http.Response, 50)
// 	defer close(sessionRespones)

// 	// レスポンスを直列化してソケットに書き出す専用の goroutine
// 	go writeToConn(sessionRespones, conn)
// 	reader := bufio.NewReader(conn)
// 	for {
// 		// レスポンスを受け取ってセッションのキューに入れる
// 		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
// 		// リクエストを読み込む
// 		request, err := http.ReadRequest(reader)
// 		if err != nil {
// 			neterr, ok := err.(net.Error)
// 			if ok && neterr.Timeout() {
// 				fmt.Println("Timeout")
// 				break
// 			} else if err == io.EOF {
// 				break
// 			}
// 			panic(err)
// 		}
// 		sessionRespone := make(chan *http.Response)
// 		sessionRespones <- sessionRespone
// 		// 非同期でレスポンスを実行
// 		go handleRequest(request, sessionRespone)
// 	}
// }

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
// 		go processSession(conn)
// 	}
// }

// =======================================================
// ## 6.8.2 チャンク形式のクライアントの実装
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strconv"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = request.Write(conn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	reader := bufio.NewReader(conn)
// 	response, err := http.ReadResponse(reader, request)
// 	if err != nil {
// 		panic(err)
// 	}

// 	dump, err := httputil.DumpResponse(response, false)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(dump))

// 	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
// 		panic("wrong transfer encoding")
// 	}

// 	for {
// 		// サイズを取得
// 		sizeStr, err := reader.ReadBytes('\n')
// 		if err == io.EOF {
// 			break
// 		}

// 		// 16進数のサイズをパース。サイズがゼロならクローズ
// 		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
// 		if size == 0 {
// 			break
// 		}
// 		if err != nil {
// 			panic(err)
// 		}

// 		// サイズ数分バッファを確保して読み込み
// 		line := make([]byte, int(size))
// 		io.ReadFull(reader, line)
// 		reader.Discard(2)
// 		fmt.Printf(" %d bytes; %s\n", size, string(line))
// 	}
// }

// =======================================================
// ## 6.8.1 チャンク形式のサーバーの実装
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strings"
// )

// var contents = []string{
// 	"これは、わたしが小さいときに、村の茂兵（もへい）というおじいさんからきいたお話です。",
// 	"むかしは、わたしたちの村のちかくの、中山というところに小さなお城（しろ）があって、中山さまというおとのさまがおられたそうです。",
// 	"その中山から、すこしはなれた山の中に、「ごんぎつね」というきつねがいました。ごんは、ひとりぼっちの小ぎつねで、しだのいっぱいしげった森の中に穴（あな）をほって住んでいました。そして、夜でも昼でも、あたりの村へ出ていって、いたずらばかりしました。畑へ入っていもをほりちらしたり、菜種（なたね）がらの、ほしてあるのへ火をつけたり、百姓家（ひゃくしょうや）のうら手につるしてあるとんがらしをむしり取っていったり、いろんなことをしました。",
// 	"ある秋のことでした。二、三日雨がふりつづいたそのあいだ、ごんは、ほっとして穴（あな）からはい出しました。空はからっと晴れていて、もずの声がキンキンひびいていました。",
// 	"ごんは、村の小川のつつみまで出てきました。あたりのすすきの穂（ほ）には、まだ雨のしずくが光っていました。川はいつもは水が少ないのですが、三日もの雨で、水がどっとましていました。ただのときは水につかることのない、川べりのすすきやはぎのかぶが、黄色くにごった水に横だおしになって、もまれています。ごんは川下の方へと、ぬかるみ道を歩いていきました。",
// 	"ふと見ると、川の中に人がいて、何かやっています。ごんは、見つからないように、そうっと草の深いところへ歩きよって、そこからじっとのぞいてみました。",
// }

// // クライアントは gzip を受け入れ可能化？
// func isGZipAcceptable(request *http.Request) bool {
// 	return strings.Index(
// 		strings.Join(request.Header["Accept-Encoding"], ","),
// 		"gzip",
// 	) != -1
// }

// // 1 セッションの処理をする
// func processSession(conn net.Conn) {
// 	fmt.Printf("Accpet %v\n", conn.RemoteAddr())
// 	defer conn.Close()

// 	for {
// 		// リクエストを読み込む
// 		request, err := http.ReadRequest(bufio.NewReader(conn))
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			panic(err)
// 		}

// 		dump, err := httputil.DumpRequest(request, true)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(dump))

// 		// レスポンスを書き込む
// 		fmt.Fprintf(conn, strings.Join([]string{
// 			"HTTP/1.1 200 OK",
// 			"Content-Type:  text/plain; charset=utf-8",
// 			"Transfer-Encoding: chunked",
// 			"",
// 			"",
// 		}, "\r\n"))
// 		for _, content := range contents {
// 			bytes := []byte(content)
// 			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
// 		}
// 		fmt.Fprintf(conn, "0\r\n\r\n")
// 	}
// }

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
// 		go processSession(conn)
// 	}
// }

// =======================================================
// ## 6.7.2 gzip 圧縮に対応したサーバー
// import (
// 	"bufio"
// 	"bytes"
// 	"compress/gzip"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"strings"
// 	"time"
// )

// // クライアントは gzip を受け入れ可能化？
// func isGZipAcceptable(request *http.Request) bool {
// 	return strings.Index(
// 		strings.Join(request.Header["Accept-Encoding"], ","),
// 		"gzip",
// 	) != -1
// }

// // 1 セッションの処理をする
// func processSession(conn net.Conn) {
// 	fmt.Printf("Accpet %v\n", conn.RemoteAddr())
// 	defer conn.Close()

// 	for {
// 		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
// 		// リクエストを読み込む
// 		request, err := http.ReadRequest(bufio.NewReader(conn))
// 		if err != nil {
// 			neterr, ok := err.(net.Error)
// 			if ok && neterr.Timeout() {
// 				fmt.Println("Timeout")
// 				break
// 			} else if err == io.EOF {
// 				break
// 			}
// 		}
// 		dump, err := httputil.DumpRequest(request, true)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(dump))

// 		response := http.Response{
// 			StatusCode: 200,
// 			ProtoMajor: 1,
// 			ProtoMinor: 1,
// 			Header:     make(http.Header),
// 		}
// 		if isGZipAcceptable(request) {
// 			content := "Hello World (gzipped)\n"
// 			// コンテンツを gzip 化して転送
// 			var buffer bytes.Buffer
// 			writer := gzip.NewWriter(&buffer)
// 			io.WriteString(writer, content)
// 			writer.Close()
// 			response.Body = io.NopCloser(&buffer)
// 			response.ContentLength = int64(buffer.Len())
// 			response.Header.Set("Content-Encoding", "gzip")
// 		} else {
// 			content := "Hello World\n"
// 			response.Body = io.NopCloser(strings.NewReader(content))
// 			response.ContentLength = int64(len(content))
// 		}
// 		response.Write(conn)
// 	}
// }

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
// 		go processSession(conn)
// 	}
// }

// =======================================================
// ## 6.7.1 gzip 圧縮に対応したクライアント
// import (
// 	"bufio"
// 	"compress/gzip"
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/http/httputil"
// 	"os"
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
// 		request.Header.Set("Accept-Encoding", "gzip")

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

// 		defer response.Body.Close()

// 		if response.Header.Get("Content-Encoding") == "gzip" {
// 			reader, err := gzip.NewReader(response.Body)
// 			if err != nil {
// 				panic(err)
// 			}
// 			io.Copy(os.Stdout, reader)
// 			reader.Close()
// 		} else {
// 			io.Copy(os.Stdout, response.Body)
// 		}

// 		// 全部送信完了していれば終了
// 		current++
// 		if current == len(sendMessages) {
// 			break
// 		}
// 	}
// 	conn.Close()
// }

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
