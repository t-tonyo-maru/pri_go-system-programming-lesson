package main

import (
	"errors"
	"io"
	"log"
	"strings"

	"github.com/google/shlex"
	"github.com/peterh/liner"
)

func parseCmd(cmdStr string) (cmd string, args []string, err error) {
	l := shlex.NewLexer(strings.NewReader(cmdStr))

	cmd, err = l.Next()
	if err != nil {
		return
	}

	for {
		token, nextErr := l.Next()
		if errors.Is(nextErr, io.EOF) {
			break
		}
		if nextErr != nil {
			err = nextErr
			return
		}
		args = append(args, token)
	}
	return
}

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		cmdStr, err := line.Prompt(" ")
		if err == nil {
			if cmdStr == "" {
				continue
			}
			// ここでコマンドを処理する
			cmd, args, parseErr := parseCmd(cmdStr)
			if parseErr != nil {
				log.Print("Error parsing command:", parseErr)
				continue
			}
			log.Printf("cmd=%s args=%v", cmd, args)
			continue
		}

		if errors.Is(err, io.EOF) {
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
