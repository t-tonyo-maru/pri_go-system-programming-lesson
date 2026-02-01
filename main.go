package main

// =======================================================
// ## 15.2.7 決まった数の goroutine でタスクを消化: ワーカープール

// =======================================================
// ## 15.2.6 並列 for ループ
// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	tasks := []string{
// 		"cmake ...",
// 		"cmake, . --build Release",
// 		"cpack",
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(len(tasks))
// 	for _, task := range tasks {
// 		go func(task string) {
// 			// ジョブを実行
// 			// このサンプルでは出力だけしている
// 			fmt.Println(task)
// 			wg.Done()
// 		}(task)
// 	}
// 	wg.Wait()
// }

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
