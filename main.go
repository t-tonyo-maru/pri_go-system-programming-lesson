package main

import (
	"context"
	"fmt"
	"os"

	"github.com/billziss-gh/cgofuse/fuse"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)

type CloudFileSystem struct {
	fuse.CloudFileSystem
	bucket *blob.Bucket
}

func main() {
	ctx := context.Background()
	if len(os.Args) < 3 {
		fmt.Printf("%s [bucket-path] [mount-point] etc…", os.Args[0])
		os.Exit(1)
	}
	b, err := blob.OpenBucket(ctx, os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer b.Close()
	cf := &CloudFileSystem{bucket: b}
	host := fuse.NewFileSytemHost(cf)
	host.Mount(os.Args[2], os.Args[3])
}

// =======================================================
// ## 10.6 FUSE を使った自作のファイルシステムの作成

// =======================================================
// ## 10.5 select 属のシステムコールによる I/O 多重化
// import (
// 	"fmt"
// 	"syscall"
// )

// func main() {
// 	kq, err := syscall.Kqueue()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 監視対象のファイルディスクリプタを取得
// 	fd, err := syscall.Open("./test", syscall.O_RDONLY, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 監視したいイベントの構造体を作成
// 	ev1 := syscall.Kevent_t{
// 		Ident:  uint64(fd),
// 		Filter: syscall.EVFILT_VNODE,
// 		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
// 		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE,
// 		Data:   0,
// 		Udata:  nil,
// 	}
// 	// イベント待ちの無限ループ
// 	for {
// 		// kevent を作成
// 		events := make([]syscall.Kevent_t, 10)
// 		nev, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// イベントを確認
// 		for i := 0; i < nev; i++ {
// 			fmt.Printf("Event [%d] -> %v\n", i, events[i])
// 		}
// 	}
// }

// =======================================================
// ## 10.3 ファイルのメモリへのマッピング
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"

// 	"github.com/edsrzf/mmap-go"
// )

// func main() {
// 	// テストデータの読み込み
// 	var testData = []byte("0123456789")
// 	var testPath = filepath.Join(os.TempDir(), "testdata")
// 	err := ioutil.WriteFile(testPath, testData, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// メモリにマッピング
// 	// m は []byte のエイリアスなので添字アクセス可能
// 	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	m, err := mmap.Map(f, mmap.RDWR, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer m.Unmap()

// 	// メモリ上のデータを修正して書き込む
// 	m[9] = 'X'
// 	m.Flush()

// 	// 読み込んでみる
// 	fileData, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("original: %s\n", testData)
// 	fmt.Printf("mmap:     %s\n", m)
// 	fmt.Printf("file:     %s\n", fileData)
// }

// =======================================================
// ## 10.2.2 LockFileEx() による Windows でのファイルロック

// =======================================================
// ## 10.2.1 syscall.Flock() による POSIX 系 OS でのファイルロック
// import (
// 	"io/fs"
// 	"syscall"
// )

// type lockType int16

// const (
// 	readLock  lockType = syscall.LOCK_SH
// 	writeLock lockType = syscall.LOCK_EX
// )

// func lock(f fs.File, lt lockType) (err error) {
// 	for {
// 		err = syscall.Flock(int(f.Fd()), int(lt))
// 		if err != syscall.EINTR {
// 			break
// 		}
// 	}
// 	return err
// }

// func unlock(f fs.File) error {
// 	return lock(f, syscall.LOCK_UN)
// }

// func main() {

// }

// =======================================================
// ## 10.1 ファイルの変更監視
// import (
// 	"log"

// 	"gopkg.in/fsnotify/fsnotify.v1"
// )

// func main() {
// 	counter := 0
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer watcher.Close()

// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case event := <-watcher.Events:
// 				log.Println("event: ", event)
// 				if event.Op&fsnotify.Create == fsnotify.Create {
// 					log.Println("created file: ", event.Name)
// 					counter++
// 				} else if event.Op&fsnotify.Write == fsnotify.Write {
// 					log.Println("modified file: ", event.Name)
// 					counter++
// 				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
// 					log.Println("removed file: ", event.Name)
// 					counter++
// 				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
// 					log.Println("renamed file: ", event.Name)
// 					counter++
// 				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
// 					log.Println("chmod file: ", event.Name)
// 					counter++
// 				}
// 			case err := <-watcher.Errors:
// 				log.Println("error: ", err)
// 			}
// 			if counter > 3 {
// 				done <- true
// 			}
// 		}
// 	}()

// 	err = watcher.Add(".")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	<-done
// }
