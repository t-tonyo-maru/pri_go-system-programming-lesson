package main

// =======================================================
// ## 18.4.5 時刻のフォーマット
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	fmt.Println(now.Format(time.RFC822))

// 	fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
// }

// =======================================================
// ## 18.4.3 チャネルを使ったタイマー
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("waiting 5 seconds")
// 	for now := range time.Tick(5 * time.Second) {
// 		fmt.Println("now: ", now)
// 	}
// }

// =======================================================
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("waiting 5 seconds")
// 	after := time.After(5 * time.Second)
// 	<-after
// 	fmt.Println("done")
// }

// =======================================================
// ## 18.4.2 スリープ
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("waiting 5 seconds")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("done")
// }

// =======================================================
// ## 18.4.1 時間と時刻
// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	// 5s
// 	fmt.Println(5 * time.Second)
// 	// 10ms
// 	fmt.Println(10 * time.Millisecond)
// 	// 10m30s
// 	t, _ := time.ParseDuration("10m30s")
// 	fmt.Println(t)

// 	// 現在時刻
// 	fmt.Println(time.Now())
// 	// 指定日時を作成
// 	fmt.Println(time.Date(2017, time.August, 26, 11, 50, 30, 0, time.Local))
// 	// フォーマットを指定してパース
// 	fmt.Println(time.Parse(time.Kitchen, "11:30AM"))
// 	// Epoch タイムから作成
// 	fmt.Println(time.Unix(1503673200, 0))

// 	// 3h後の時間
// 	fmt.Println(time.Now().Add(3 * time.Hour))
// 	// ファイル変更日時が何日前かを知る
// 	fileInfo, _ := os.Stat(".vimrc")
// 	fmt.Println("%v 前", time.Now().Sub(fileInfo.ModTime()))
// 	// 時間を1時間ごとに丸める
// 	fmt.Println(time.Now().Round(time.Hour))
// }

// =======================================================
// ## 18.1 OS のタイマー / カウンターの仕組み
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	t := time.Now()
// 	fmt.Println(t.String())
// }
