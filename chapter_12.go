package main

// =======================================================
// ## 12.6.2 外部プロセスに対して自分が擬似端末だと詐称する

// =======================================================
// ## 12.6.1 プロセスの出力に色づけをする
// import (
// 	"fmt"
// 	"io"
// 	"os"

// 	"github.com/mattn/go-colorable"
// 	"github.com/mattn/go-isatty"
// )

// var data = "\033[34m\033[47m\033[4mB\033[31me\n\033[24m\033[30mOS\033[49m\033[m\n"

// func main() {
// 	var stdOut io.Writer
// 	if isatty.IsTerminal(os.Stdout.Fd()) {
// 		stdOut = colorable.NewColorableStdout()
// 	} else {
// 		stdOut = colorable.NewNonColorable(os.Stdout)
// 	}
// 	fmt.Fprintln(stdOut, data)
// }

// =======================================================
// ## 12.5.2 リアルタイムな入出力
// import (
// 	"bufio"
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	count := exec.Command("./count")
// 	stdout, _ := count.StdoutPipe()
// 	go func() {
// 		scanner := bufio.NewScanner(stdout)
// 		for scanner.Scan() {
// 			fmt.Printf("(stdout) %s\n", scanner.Text())
// 		}
// 	}()
// 	err := count.Run()
// 	if err != nil {
// 		panic(err)
// 	}
// }
// =======================================================
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// =======================================================
// ## 12.5.1 exec.Cmd によるプロセスの起動
// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	if len(os.Args) == 1 {
// 		return
// 	}

// 	cmd := exec.Command(os.Args[1], os.Args[2:]...)
// 	err := cmd.Run()
// 	if err != nil {
// 		panic(err)
// 	}
// 	state := cmd.ProcessState
// 	fmt.Printf("%s\n", state.String())
// 	fmt.Printf("  Pid: %d\n", state.Pid())
// 	fmt.Printf("  System: %v\n", state.SystemTime())
// 	fmt.Printf("  User: %v\n", state.UserTime())
// }

// =======================================================
// ## 12.3 自分以外のプロセスの名前や資源情報の取得
// import (
// 	"fmt"
// 	"os"

// 	"github.com/shirou/gopsutil/process"
// )

// func main() {
// 	p, _ := process.NewProcess(int32(os.Getppid()))
// 	name, _ := p.Name()
// 	cmd, _ := p.Cmdline()
// 	fmt.Printf("parent pid: %d name: '%s' cmd: '%s'\n", p.Pid, name, cmd)
// }

// =======================================================
// ## 12.2.2 終了コードと$・ERRORLEVEL
// import "os"

// func main() {
// 	os.Exit(1)
// }

// =======================================================
// ## 12.1.6 作業フォルダ
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	wd, _ := os.Getwd()
// 	fmt.Println(wd)
// }

// =======================================================
// ## 12.1.5 実行ユーザーIDと実行グループID
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Printf("ユーザーID: %d\n", os.Getuid())
// 	fmt.Printf("グループID: %d\n", os.Getgid())
// 	fmt.Printf("実行ユーザーID: %d\n", os.Geteuid())
// 	fmt.Printf("実行グループID: %d\n", os.Getegid())
// }

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
