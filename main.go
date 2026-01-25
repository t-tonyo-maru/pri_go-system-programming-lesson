package main

func main() {

}

// =======================================================
// ## 9.2.5 ファイルの属性の取得
// import (
// 	"fmt"
// 	"os"
// )

// /**
//  * コマンド: go run main.go README.md
//  */
// func main() {
// 	if len(os.Args) == 1 {
// 		fmt.Printf("%s [exec file name]", os.Args[0])
// 		os.Exit(1)
// 	}

// 	info, err := os.Stat(os.Args[1])
// 	if err == os.ErrNotExist {
// 		fmt.Printf("file not found: %s\n", os.Args[1])
// 	} else if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("FileInfo")
// 	fmt.Printf("  ファイル名: %v\n", info.Name())
// 	fmt.Printf("  サイズ: %v\n", info.Size())
// 	fmt.Printf("  変更日時: %v\n", info.ModTime())
// 	fmt.Printf("Mode()")
// 	fmt.Printf("  ディレクトリ?: %v\n", info.Mode().IsDir())
// 	fmt.Printf("  読み書き可能な通常ファイル?: %v\n", info.Mode().IsRegular())
// 	fmt.Printf("  Unix のファイルアクセス権限ビット?: %o\n", info.Mode().Perm())
// 	fmt.Printf("  モードのテキスト表現?: %v\n", info.Mode().String())
// }

// =======================================================
// ## 9.2.4 ファイルの削除/移動/リネーム
// import (
// 	"io"
// 	"os"
// )

// // コピー
// func main() {
// 	oldFile, err := os.Create("old_name.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	oldFile.WriteString("old name")
// 	oldFile, err = os.Open("old_name.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	os.Mkdir("other_device", 0755)
// 	newFile, err := os.Create("./other_device/new_file.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer newFile.Close()

// 	_, err = io.Copy(newFile, oldFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	os.Remove("old_name.txt")
// }

// =======================================================
// import "os"

// // リネーム
// func rename() {
// 	// リネーム
// 	os.Rename("old_name.txt", "new_name.txt")
// 	// 移動
// 	os.Rename("olddir/file.txt", "newdir/file.txt")
// }

// func main() {
// 	_, err := os.Create("old_name.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	os.Mkdir("newdir", 0755)
// 	_, err = os.Create("newdir/file.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	rename()
// }

// =======================================================
// import (
// 	"math/rand"
// 	"os"
// )

// func getString() string {
// 	str := ""

// 	strings := []string{
// 		"a", "b", "c", "d", "e", "f", "g",
// 	}

// 	for i := 0; i < 30000; i++ {
// 		str += strings[rand.Intn(len(strings)-1)]
// 	}

// 	return str
// }

// func main() {
// 	file, err := os.Create("server.log")
// 	if err != nil {
// 		panic(err)
// 	}
// 	file.WriteString(getString())

// 	// 先頭 100 バイトで切る
// 	os.Truncate("server.log", 100)
// 	// Truncateメソッドを利用する場合
// 	file.Truncate(100)
// }

// =======================================================
// import "os"

// func create() {
// 	_, err := os.Create("server.log")
// 	if err != nil {
// 		panic(err)
// 	}

// 	os.Mkdir("workdir", 0755)
// }

// func remove() {
// 	os.Remove("server.log")
// 	os.Remove("workdir")
// }

// func main() {
// 	create()
// 	remove()
// }

// =======================================================
// ## 9.2.3 ディレクトリの作成
// import "os"

// func main() {
// 	// フォルダを1階層だけ作成
// 	os.Mkdir("setting", 0755)
// 	// 深いフォルダを1回で作成
// 	os.MkdirAll("setting/myapp/networksettings", 0755)
// }

// =======================================================
// ## 9.2.1 ファイル作成/読み込み
// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// // 新規作成
// func open() {
// 	file, err := os.Create("textfile.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	io.WriteString(file, "New file content\n")
// }

// // 読み込み
// func read() {
// 	file, err := os.Open("textfile.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	fmt.Println("Read file:")
// 	io.Copy(os.Stdout, file)
// }

// // 追記モード
// func append() {
// 	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	io.WriteString(file, "Appended content\n")
// }

// func main() {
// 	open()
// 	append()
// 	read()
// }
