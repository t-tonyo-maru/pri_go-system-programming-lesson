package main

// =======================================================
// ## 15.2.8 依存関係のあるタスクを表現する Future/Promise

// =======================================================
// ## 15.2.7 決まった数の goroutine でタスクを消化: ワーカープール
// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// // 元金均等
// func calc(id, price int, interestRate float64, year int) {
// 	months := year * 12
// 	interest := 0
// 	for i := 0; i < months; i++ {
// 		balance := price * (months - i) / months
// 		interest += int(float64(balance) * interestRate / 12)
// 	}
// 	fmt.Printf("year=%d total=%d interest=%d id=%d\n", year, price+interest, interest, id)
// }

// // ワーカー
// func worker(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
// 	// タスクがなくなってタスクのチャネルが close されるまで無限ループ
// 	for year := range years {
// 		calc(id, price, interestRate, year)
// 		wg.Done()
// 	}
// }

// func main() {
// 	// 借入金
// 	price := 400000000
// 	// 利子 1.1%
// 	interestRate := 0.011
// 	// タスクは chan に格納
// 	years := make(chan int, 35)
// 	for i := 1; i < 36; i++ {
// 		years <- i
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(35)
// 	// CPU コア数分の goroutine を起動
// 	for i := 0; i < runtime.NumCPU(); i++ {
// 		go worker(i, price, interestRate, years, &wg)
// 	}
// 	// すべてのワーカーが終了する
// 	close(years)
// 	wg.Wait()
// }

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
