package main

func main() {

}

// =======================================================
// ## 12.1.5 実行ユーザーIDと実行グループID

// =======================================================
// ## 12.1.4 ユーザーIDとグループID
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Printf("ユーザーID: %d\n", os.Getuid())
// 	fmt.Printf("グループID: %d\n", os.Getgid())
// 	groups, _ := os.Getgroups()
// 	fmt.Printf("サブグループID: %v\n", groups)
// }

// =======================================================
// ## 12.1.3 プロセスグループとセッショングループ
// import (
// 	"fmt"
// 	"os"
// 	"syscall"
// )

// func main() {
// 	sid, _ := syscall.Getsid(os.Getpid())
// 	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", syscall.Getpgrp(), sid)
// }

// =======================================================
// ## 12.1.2 プロセス ID
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Printf("プロセスID: %d\n", os.Getegid())
// 	fmt.Printf("親プロセスID: %d\n", os.Getppid())
// }

// =======================================================
// ## 12.1.1 実行ファイル名
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	path, _ := os.Executable()
// 	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
// 	fmt.Printf("実行ファイルパス: %s\n", path)
// }
