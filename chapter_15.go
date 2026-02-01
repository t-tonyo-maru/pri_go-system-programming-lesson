package main

// =======================================================
// ## 14.7.5 sync.Map
// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// 初期化
// 	smap := &sync.Map{}

// 	// なんでも入れられる
// 	smap.Store("hello", "world")
// 	smap.Store(1, 2)

// 	// 削除
// 	smap.Delete("test")

// 	// 取り出し方法
// 	value, ok := smap.Load("hello")
// 	fmt.Printf("key=%v value=%v exists=%v\n", "hello", value, ok)
// }

// =======================================================
// ## 14.7.4 sync.Once
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var mutex sync.Mutex
// 	cond := sync.NewCond(&mutex)

// 	for _, name := range []string{"A", "B", "C"} {
// 		go func(name string) {
// 			// ロックしてから Wait メソッドを呼ぶ
// 			mutex.Lock()
// 			defer mutex.Unlock()
// 			// BroadCast() が呼ばれるまで待機
// 			cond.Wait()
// 			// 呼ばれた
// 			fmt.Println(name)
// 		}(name)
// 	}

// 	fmt.Println("よーい")
// 	time.Sleep(time.Second)
// 	fmt.Println("どん！")
// 	cond.Broadcast()
// 	time.Sleep(time.Second)
// }

// =======================================================
// ## 14.7.3 sync.Once
// import (
// 	"fmt"
// 	"sync"
// )

// func initialize() {
// 	fmt.Println("初期化処理")
// }

// var once sync.Once

// func main() {
// 	// 3回呼び出しても一度しか呼ばれない
// 	once.Do(initialize)
// 	once.Do(initialize)
// 	once.Do(initialize)
// }

// =======================================================
// ## 14.7.2 sync.WaitGroup
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	// ジョブ数をあらかじめ登録
// 	wg.Add(2)

// 	go func() {
// 		// 非同期で仕事をする
// 		fmt.Println("仕事1")
// 		// 数秒待機
// 		time.Sleep(time.Second * 2)
// 		// Done で完了を通知
// 		wg.Done()
// 	}()

// 	go func() {
// 		// 自動機で字ごとをする(2)
// 		fmt.Println("仕事2")
// 		// 数秒待機
// 		time.Sleep(time.Second * 5)
// 		// Done で完了を通知
// 		wg.Done()
// 	}()

// 	// すべての処理が終了するまで待つ
// 	wg.Wait()
// 	fmt.Println("終了")
// }

// =======================================================
// ## 14.7.1 sync.Mutex / sync.RWMutex
// import (
// 	"fmt"
// 	"sync"
// )

// var id int

// func generateId(mutex *sync.Mutex) int {
// 	// Lock() / Unlock() をペアで呼び出してブロックする
// 	mutex.Lock()
// 	id++
// 	result := id
// 	mutex.Unlock()
// 	return result
// }

// func main() {
// 	// sync.Mutex 構造体の変数宣言
// 	// 次の宣言をしてもポインタ型になるだけで正常に動作する
// 	// mutex := new(sync.Mutex)
// 	var mutex sync.Mutex
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			fmt.Printf("id: %d\n", generateId(&mutex))
// 		}()
// 	}
// }

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
