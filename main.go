package main

func main() {

}

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
