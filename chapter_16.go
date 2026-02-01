package main

// =======================================================
// ## 15.2.10 自立した複数のシステムで協調動作: アクターモデル
// import (
// 	"fmt"

// 	"github.com/AsynkronIT/goconsole"
// 	"github.com/AsynkronIT/protoactor-go/actor"
// )

// // メッセージの構造体
// type hello struct{ Who string }

// // アクターの構造体
// type helloActor struct{}

// // アクターのメールボックス受信時に呼ばれるメソッド
// func (state *helloActor) Receive(context actor.Context) {
// 	switch msg := context.Message().(type) {
// 	case "hello":
// 		fmt.Printf("Hello %v\n", msg.Who)
// 	}
// }

// func main() {
// 	props := actor.FromInstance(&helloActor{})
// 	pid := actor.Spawn(props)
// 	pid.Tell(&hello{Who: "Roger"})
// 	console.ReadLine()
// }

// =======================================================
// ## 15.2.9 イベントの流れを定義する: ReactiveX
// import (
// 	"fmt"
// 	"os"
// 	"strings"
//
// 	"github.com/reactivex/rxgo/observer"
// )

// func main() {
// 	// observable を作成
// 	emitter := make(chan interface{})
// 	source := observable.Observable(emitter)

// 	// イベントを受け取る observer を作成
// 	watcher := observer.Observer{
// 		NextHandler: func(item interface{}) {
// 			line := item.(string)
// 			if strings.HasPrefix(line, "func ") {
// 				fmt.Println(line)
// 			}
// 		},
// 		ErrHandler: func(err error) {
// 			fmt.Printf("Encounterd error: %v\n", err)
// 		},
// 		DoneHandler: func() {
// 			fmt.Println("Done!")
// 		},
// 	}

// 	// observable と observer を接続
// 	sub := source.Subscribe(watcher)

// 	// observalbe に値を投入
// 	go func() {
// 		content, err := os.ReadFile("reactive.go")
// 		if err != nil {
// 			emitter <- err
// 		} else {
// 			for _, line := range strings.Split(string(content), "\n") {
// 				emitter <- line
// 			}
// 		}
// 		close(emitter)
// 	}()

// 	// 終了まち
// 	<-sub
// }

// =======================================================
// ## 15.2.8 依存関係のあるタスクを表現する Future/Promise
// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type StringFuture struct {
// 	receiver chan string
// 	cache    string
// }

// func NewStringFuture() (*StringFuture, chan string) {
// 	f := &StringFuture{
// 		receiver: make(chan string),
// 	}
// 	return f, f.receiver
// }

// func (f *StringFuture) Get() string {
// 	r, ok := <-f.receiver
// 	if ok {
// 		close(f.receiver)
// 		f.cache = r
// 	}
// 	return f.cache
// }

// func (f *StringFuture) Close() {
// 	close(f.receiver)
// }

// func readFile(path string) *StringFuture {
// 	// ファイルを読み込み、その結果を返す Future を返す
// 	promise, future := NewStringFuture()
// 	go func() {
// 		content, err := os.ReadFile(path)
// 		if err != nil {
// 			fmt.Printf("read error %s\n", err.Error())
// 		} else {
// 			future <- string(content)
// 		}
// 	}()
// 	return promise
// }

// func printFunc(futureSource *StringFuture) chan []string {
// 	// 文字列中の関数一覧を返す Future を返す
// 	promise := make(chan []string)
// 	go func() {
// 		var result []string
// 		// future が解決するまで待って待機
// 		for _, line := range strings.Split(futureSource.Get(), "\n") {
// 			if strings.HasPrefix(line, "func ") {
// 				result = append(result, line)
// 			}
// 		}
// 		// 約束を果たした
// 		promise <- result
// 	}()
// 	return promise
// }

// func countLines(futureSource *StringFuture) chan int {
// 	promise := make(chan int)
// 	go func() {
// 		promise <- len(strings.Split(futureSource.Get(), "\n"))
// 	}()
// 	return promise
// }

// func main() {
// 	futureSource := readFile("future_promise.go")
// 	futureFuncs := printFunc(futureSource)
// 	fmt.Println(strings.Join(<-futureFuncs, "\n"))
// 	fmt.Println(<-countLines(futureSource))
// }

// =======================================================
// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func readFile(path string) chan string {
// 	// ファイルを読み込み、その結果を返す Future を返す
// 	promise := make(chan string)
// 	go func() {
// 		content, err := os.ReadFile(path)
// 		if err != nil {
// 			fmt.Printf("read error %s\n", err.Error())
// 			close(promise)
// 		} else {
// 			// 約束を果たした
// 			promise <- string(content)
// 		}
// 	}()
// 	return promise
// }

// func printFunc(futureSource chan string) chan []string {
// 	// 文字列中の関数一覧を返す Future を返す
// 	promise := make(chan []string)
// 	go func() {
// 		var result []string
// 		// future が解決するまで待って待機
// 		for _, line := range strings.Split(<-futureSource, "\n") {
// 			if strings.HasPrefix(line, "func ") {
// 				result = append(result, line)
// 			}
// 		}
// 		// 約束を果たした
// 		promise <- result
// 	}()
// 	return promise
// }

// func main() {
// 	futureSource := readFile("future_promise.go")
// 	futureFuncs := printFunc(futureSource)
// 	fmt.Println(strings.Join(<-futureFuncs, "\n"))
// }

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
