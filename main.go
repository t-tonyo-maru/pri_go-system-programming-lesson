package main

func main() {

}

// =======================================================
// ## 13.5.3 Server::Starter 対応のサーバーの実装例

// =======================================================
// ## 13.4.4 シグナルを他のプロセスに送る
// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Printf("usuge: %s [pid]\n", os.Args[0])
// 		return
// 	}

// 	// 第一引数で指定されたプロセス ID を数値に変換
// 	pid, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		panic(err)
// 	}
// 	process, err := os.FindProcess(pid)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// シグナルを送る
// 	process.Signal(os.Kill)
// 	// Kill の場合は次のショートカットも利用可能
// 	process.Kill()
// }

// =======================================================
// ## 13.4.3 シグナルの送信を停止させる

// =======================================================
// ## 13.4.2 シグナルのハンドラをデフォルトに戻す
// import (
// 	"os/signal"
// 	"syscall"
// )

// func main() {
// 	signal.Reset(syscall.SIGINT, syscall.SIGHUP)
// }

// =======================================================
// ## 13.4.1 シグナルを無視する
// import (
// 	"fmt"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	// 最初の10sは Ctrl + C で止まる
// 	fmt.Println("Accept Ctrl + C for 10 second")
// 	time.Sleep(time.Second * 10)

// 	// 可変長引数で任意のカスシグナル設定を可能
// 	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

// 	// 次の 10 sは Ctrl + C を無視する
// 	fmt.Println("Ignore Ctrl + C for 10 second")
// 	time.Sleep(time.Second * 10)
// }

// =======================================================
// ## 13.4 シグナルのハンドラを書く
// import (
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
// )

// func main() {
// 	// サイズが1より大きいチャネルを作成
// 	signals := make(chan os.Signal, 1)
// 	// SIGINT (Ctrl+C) を受け取る
// 	signal.Notify(signals, syscall.SIGINT)

// 	// シグナルがくるまで待つ
// 	fmt.Println("Waiting SIGINT (CTRL+C)")
// 	<-signals
// 	fmt.Println("SIGINT arrived")
// }
// =======================================================
// import (
// 	"context"
// 	"fmt"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	ctx := context.Background()
// 	sigctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
// 	defer cancel()

// 	toctx, cancel2 := context.WithTimeout(ctx, time.Second*5)
// 	defer cancel2()

// 	select {
// 	case <-sigctx.Done():
// 		fmt.Println("signal")
// 	case <-toctx.Done():
// 		fmt.Println("timeout")
// 	}
// }
