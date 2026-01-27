package main

func main() {

}

// =======================================================
// ## 10.3 ファイルのメモリへのマッピング

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
