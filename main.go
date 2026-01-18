package main

// =======================================================
// ## 4.1 goroutine
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
