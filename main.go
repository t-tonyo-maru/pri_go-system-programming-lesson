package main

// =======================================================
// ## 14.7.1 sync.Mutex / sync.RWMutex

// =======================================================
// ## 14.2.1 goroutine と情報共有
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	tasks := []string{
// 		"cmake ..",
// 		"cmake . --build Relase",
// 		"cpack",
// 	}
// 	for _, task := range tasks {
// 		go func() {
// 			fmt.Println(task)
// 		}()
// 	}
// 	time.Sleep(time.Second)
// }
// =======================================================
// import (
// 	"fmt"
// 	"time"
// )

// func sub1(c int) {
// 	fmt.Println("share by arguments", c*c)
// }

// func main() {
// 	// 引数渡し
// 	go sub1(10)

// 	// クロージャのキャプチャ渡し
// 	c := 20
// 	go func() {
// 		fmt.Println("share by capture", c*c)
// 	}()
// 	time.Sleep(time.Second)
// }
