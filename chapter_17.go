package main

// =======================================================
// ## 16.4.1 メモリアリーナ

// =======================================================
// ## 16.3.4 sync.Pool による、アロケート回数の削減
// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	var count int
// 	pool := sync.Pool{
// 		New: func() interface{} {
// 			count++
// 			return fmt.Sprintf("created: %d", count)
// 		},
// 	}

// 	// GC を呼ぶと追加された要素が消える
// 	pool.Put("removed 1")
// 	pool.Put("removed 2")
// 	runtime.GC()
// 	fmt.Println(pool.Get())
// }
// =======================================================
// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// Pool を作成。New で新規作成時のコードを実装
// 	var count int
// 	pool := sync.Pool{
// 		New: func() interface{} {
// 			count++
// 			return fmt.Sprintf("created: %d", count)
// 		},
// 	}

// 	// 追加した要素から受け取れる
// 	pool.Put("manualy added: 1")
// 	pool.Put("manualy added: 2")
// 	fmt.Println(pool.Get())
// 	fmt.Println(pool.Get())
// 	fmt.Println(pool.Get())
// }

// =======================================================
// ## 16.3.2 スライスのメモリ確保とパフォーマンス改善のヒント
// import "fmt"

// func main() {
// 	// 長さ 1, 確保された要素 2 のスライスを作成
// 	s := make([]int, 1, 2)
// 	fmt.Println(&s[0], len(s), cap(s))

// 	// 1 要素追加(確保された範囲内)
// 	s = append(s, 1)
// 	fmt.Println(&s[0], len(s), cap(s))

// 	// さらに要素を追加(新しく配列を確保され直す)
// 	s = append(s, 2)
// 	fmt.Println(&s[0], len(s), cap(s))
// }

// =======================================================
// ## 16.3.1 スライスの作成方法
// import "fmt"

// func main() {
// 	// 既存の配列を参照するスライス
// 	a := [4]int{1, 2, 3, 4}
// 	b := a[:]
// 	fmt.Println(&b[0], len(b), cap(b))

// 	// 既存の配列の一部を参照するスライス
// 	c := a[1:3]
// 	fmt.Println(&c[0], len(c), cap(c))

// 	// 何も参照していないスライス
// 	var d []int
// 	fmt.Println(len(d), cap(d))

// 	// 初期の配列とリンクされているスライス
// 	e := []int{1, 2, 3, 4}
// 	fmt.Println(&e[0], len(e), cap(e))

// 	// サイズをもったスライスを定義
// 	f := make([]int, 4)
// 	fmt.Println(&f[0], len(f), cap(f))

// 	// サイズと容量をもったスライスを定義
// 	g := make([]int, 4, 8)
// 	fmt.Println(&g[0], len(g), cap(g))
// }

// =======================================================
// ## 16.2 Go 言語の配列
// func main() {
// 	// リストの宣言
// 	var a [4]int

// 	// リストを生成
// 	b := [4]int{}

// 	// リストを生成（初期値つき）
// 	c := [4]int{0, 1, 2, 3}

// 	// リストを生成（初期値つき/要素数は自動設定）
// 	d := [...]int{0, 1, 2, 3}
// }

// =======================================================
// ## 16.1.5 ユーザーコードでメモリを使う
// func main() {
// 	// 固定長の配列を定義
// 	a := [4]int{1, 2, 3, 4}

// 	// サイズをもったスライスを定義
// 	b := make([]int, 4)

// 	// サイズとキャパシティをもったスライスを定義
// 	c := make([]int, 4, 16)

// 	// マップを定義
// 	d := make(map[string]int)

// 	// キャパシティをもったマップを定義
// 	e := make(map[string]int, 100)

// 	// バッファなしのチャネル
// 	f := make(chan string)

// 	// バッファありのチャネル
// 	g := make(chan string, 10)
// }
// =======================================================
// func main() {
// 	// プリミティブのインスタンスを定義
// 	var a int = 10

// 	// 構造体のインスタンスを new して作成
// 	// 変数にはポインタを保存
// 	var b *Struct = new(Struct)

// 	// 構造体を {} でメンバーの初期化を与えて初期化
// 	// 変数にはインスタンスを保存
// 	var c Struct = Struct{"param"}

// 	// 構造体を {} でメンバーに初期値を与えて保存
// 	var d *Struct = &Struct{"param"}
// }
