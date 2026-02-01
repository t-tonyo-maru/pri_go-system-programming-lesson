package main

func main() {

}

// =======================================================
// ## 15.2.4 開始した順で処理する: チャネルのチャネル
// import (
// 	"net"
// 	"net/http"
// )

// // 終了した順に書き出し
// func writeToConn(responses chan *http.Response, conn net.Conn) {
// 	defer conn.Close()
// 	// 順番に取り出す
// 	for response := range responses {
// 		response.Write(conn)
// 	}
// }

// // 開始した順に書き出し
// func writeToConn2(sessionResponses chan chan *http.Response, conn net.Conn) {
// 	defer conn.Close()
// 	// 順番に取り出す
// 	for sessionResponse := range sessionResponses {
// 		// 選択された仕事が終わるまで待つ
// 		response := <-sessionResponse
// 		response.Write(conn)
// 	}
// }
