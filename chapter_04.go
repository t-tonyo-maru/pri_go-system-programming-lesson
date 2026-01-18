package main

// =======================================================
// ## 4.2.5 コンテキスト
// import (
// 	"context"
// 	"fmt"
// )

// func main() {
// 	fmt.Println("start sub()")
// 	// 終了を受け取るための終了関数付きコンテキスト
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func() {
// 		fmt.Println("sub() is finished")
// 		// 終了を通知
// 		cancel()
// 	}()
// 	// 終了を待つ
// 	<-ctx.Done()
// 	fmt.Println("all tasks are finished")
// }

// =======================================================
// ## 4.2.3 for 文
// import (
// 	"fmt"
// 	"math"
// )

// func primeNumber() chan int {
// 	result := make(chan int)
// 	go func() {
// 		result <- 2
// 		for i := 3; i < 100000; i += 2 {
// 			l := int(math.Sqrt(float64(i)))
// 			found := false
// 			for j := 3; j < l+1; j += 2 {
// 				if i%j == 0 {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				result <- i
// 			}
// 		}
// 		close(result)
// 	}()

// 	return result
// }

// func main() {
// 	pn := primeNumber()
// 	// ココがポイント
// 	for n := range pn {
// 		fmt.Println(n)
// 	}
// }
// =======================================================
// ## 4.2.1 チャネルの使用方法
// import "fmt"

// func main() {
// 	fmt.Println("start sub()")
// 	// 終了を受け取るためのチャネル
// 	done := make(chan bool)
// 	go func() {
// 		fmt.Println("sub() is finished")
// 		// 終了を通知
// 		done <- true
// 	}()
// 	// 終了を待つ
// 	<-done
// 	fmt.Println("all tasks are finidhed.")
// }

// =======================================================
// ## 4.2 チャネル
// func main() {
// 	// バッファなし
// 	// tasks := make(chan string)
// 	// バッファあり
// 	tasks := make(chan string, 10)

// 	// データ送信
// 	tasks <- "cmake…"
// 	tasks <- "cmake…2"
// 	// データ受け取り
// 	task := <-tasks
// 	task, ok := <-tasks
// 	<-wait
// }

// =======================================================
// ## 4.1 goroutine
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("start sub()")

// 	go func() {
// 		fmt.Println("sub() is rnning")
// 		time.Sleep(time.Second)
// 		fmt.Println("sub() is finished")
// 	}()

// 	time.Sleep(2 * time.Second)
// }
// =======================================================
// go というキーワードをつけて実行すれば goroutine が作られて並列実行される
// import (
// 	"fmt"
// 	"time"
// )

// func sub() {
// 	fmt.Println("sub() is running")
// 	time.Sleep(time.Second)
// 	fmt.Println("sub() is finished")
// }

// func main() {
// 	fmt.Println("start sub()")
// 	go sub()
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("end sub()")
// }
