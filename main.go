package main

func main() {

}

// =======================================================
// ## 9.5.2 パスを分割する
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	dir, name := filepath.Split(os.Getenv("GOPATH"))
// 	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
// }

// =======================================================
// ## 9.5.1 ディレクトリのパスとファイル名を連結する
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
// }

// =======================================================
// ## 9.2.11 ディレクトリ情報の取得
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	dir, err := os.Open("/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, fileInfo := range fileInfos {
// 		if fileInfo.IsDir() {
// 			fmt.Println("[Dir] %s\n", fileInfo.Name())
// 		} else {
// 			fmt.Println("[File] %s\n", fileInfo.Name())
// 		}
// 	}
// }

// =======================================================
// ## 9.2.10 リンク
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// ハードリンク
// 	os.Link("oldfile.txt", "newfile.txt")
// 	// シンボリックリンク
// 	os.Symlink("oldfile.txt", "newfile-symlink.txt")
// 	// シンボリックリンクのリンク先を取得
// 	link, err := os.Readlink("newfile-symlink.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(link)
// }

// =======================================================
// ## 9.2.9 ファイルの属性の設定
// import (
// 	"os"
// 	"time"
// )

// func main() {
// 	// ファイルのモードを変更
// 	os.Chmod("setting.txt", 0644)
// 	// ファイルのオーナーを変更
// 	os.Chown("setting.txt", os.Getuid(), os.Getgid())
// 	// ファイルの最終アクセス日時と変更日時を更新
// 	os.Chtimes("setting.txt", time.Now(), time.Now())
// }

// =======================================================
// ## 9.2.8 ファイルの同一性チェック
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fileInfo1, err := os.Stat("README.md")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fileInfo2, err := os.Stat("README.md")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if os.SameFile(fileInfo1, fileInfo2) {
// 		fmt.Println("同じファイルです。")
// 	} else {
// 		fmt.Println("異なるファイルです。")
// 	}
// }

// =======================================================
// ## 9.2.6 ファイルの存在チェック
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	info, err := os.Stat("README.md")
// 	if err == os.ErrNotExist {
// 		// ファイルが存在しない
// 		fmt.Println("ファイルが存在しません")
// 	} else if err != nil {
// 		// それ以外のエラー
// 		fmt.Println("不明なエラーが発生しました")
// 	} else {
// 		// 正常ケース
// 		fmt.Println("ファイルは存在します。", info)
// 	}
// }

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
