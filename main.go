package main

import (
	"errors"
	"io"
	"log"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		if cmd, err := line.Prompt(" "); err != nil {
			if cmd == "" {
				continue
			}
			// ここでコマンドを処理する
		} else if errors.Is(err, io.EOF) {
			break
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
			break
		} else {
			log.Print("Error reading line:", err)
		}
	}
}

// =======================================================
// ## 11.5　シェルがコマンドを起動するまで

// =======================================================
// ## 11.4.2 .env ファイル
// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"os/exec"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	filename := flag.String("e", ".env", ".env file name to read")
// 	flag.Parse()
// 	cmdName := flag.Arg(0)
// 	args := flag.Args()[1:]
// 	flag.Args()

// 	cmd := exec.Command(cmdName, args...)

// 	envs := os.Environ()
// 	dotenvs, _ := godotenv.Read(*filename)
// 	for key, value := range dotenvs {
// 		envs = append(envs, key+"="+value)
// 	}
// 	cmd.Env = envs
// 	o, err := cmd.CombinedOutput()
// 	fmt.Println(string(o))
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
