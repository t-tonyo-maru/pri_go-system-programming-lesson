package main

// =======================================================
// ## 3.7 io.Reader / io.Writer でストリームを自由に操る
// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// )

// func main() {
// 	var buffer bytes.Buffer
// 	reader := bytes.NewBufferString("Example of io.TeeReader\n")
// 	teeReader := io.TeeReader(reader, &buffer)
// 	// データを読み捨てる
// 	_, _ = io.ReadAll(teeReader)

// 	// しかし、バッファには残っている
// 	fmt.Println(buffer.String())
// }

// =======================================================
// import (
// 	"bytes"
// 	"io"
// 	"os"
// )

// func main() {
// 	header := bytes.NewBufferString("----- HEADER -----\n")
// 	content := bytes.NewBufferString("Example of io.MultiReader\n")
// 	footer := bytes.NewBufferString("----- FOOTER -----\n")

// 	reader := io.MultiReader(header, content, footer)
// 	io.Copy(os.Stdout, reader)
// }

// =======================================================
// ## 3.6.3 その他の形式の決まったフォーマットの文字列の解析
// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"strings"
// )

// var csvSource = `
// 13101, 100, 10000001, 東京都, 千代田区
// 13102, 200, 20000002, 東京都, 千代田区
// 13103, 300, 30000003, 東京都, 千代田区
// `

// func main() {
// 	reader := strings.NewReader(csvSource)
// 	csvReader := csv.NewReader(reader)
// 	for {
// 		line, err := csvReader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Println(line[0], line[1:5])
// 	}
// }
// =======================================================
// ## 3.6.2 データ型を指定して解析
// import (
// 	"fmt"
// 	"strings"
// )

// var source = "123 1.234 1.0e4 test"

// func main() {
// 	reader := strings.NewReader(source)
// 	var i int
// 	var f, g float64
// 	var s string
// 	fmt.Fscan(reader, &i, &f, &g, &s)
// 	fmt.Printf("i=%#v f=%#v g=%#v s=%#v", i, f, g, s)
// }

// =======================================================
// ## 3.6.1 改行/単語で区切る
// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"strings"
// )

// var source = `1行目
// 2行目
// 3行目`

// func main() {
// 	reader := bufio.NewReader(strings.NewReader(source))
// 	for {
// 		line, err := reader.ReadString('\n')
// 		fmt.Printf("%v\n", line)
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// =======================================================
// ## 3.5.4 PNG 画像に秘密のテキストを入れてみる
// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"hash/crc32"
// 	"io"
// 	"os"
// )

// func dumpChunk(chunk io.Reader) {
// 	var length int32
// 	binary.Read(chunk, binary.BigEndian, &length)
// 	buffer := make([]byte, 4)
// 	chunk.Read(buffer)
// 	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
// }

// func readChunks(file *os.File) []io.Reader {
// 	// チャンクを格納する配列
// 	var chunks []io.Reader

// 	// 最初の 8 バイトを飛ばす
// 	file.Seek(8, 0)
// 	var offset int64 = 8

// 	for {
// 		var length int32
// 		err := binary.Read(file, binary.BigEndian, &length)
// 		if err == io.EOF {
// 			break
// 		}
// 		chunks = append(chunks,
// 			io.NewSectionReader(file, offset, int64(length)+12))
// 		// 次のチャンクの先頭に移動
// 		// 現在位置は長さを読み終わった箇所なので
// 		// チャンク名（4バイト）+データ長+CRC(4バイト)先に移動
// 		offset, _ = file.Seek(int64(length+8), 1)
// 	}
// 	return chunks
// }

// func textChunk(text string) io.Reader {
// 	byteText := []byte(text)
// 	crc := crc32.NewIEEE()
// 	var buffer bytes.Buffer
// 	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))

// 	// CRC 計算とバッファへの書き込みを同時に行う MultiWriter
// 	writer := io.MultiWriter(&buffer, crc)
// 	io.WriteString(writer, "teXt") // 2バイト目の5ビット目を立てる
// 	writer.Write(byteText)
// 	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
// 	return &buffer
// }

// func main() {
// 	file, err := os.Open("PNG_transparency_demonstration_1.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	newFile, err := os.Create("PNG_transparency_demonstration_secret.png")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer newFile.Close()
// 	chunks := readChunks(file)
// 	// シグニチャ書き込み
// 	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
// 	// 先頭に必要な IHDR チャンクを書き込み
// 	io.Copy(newFile, chunks[0])
// 	// テキストチャンクを追加
// 	io.Copy(newFile, textChunk("Lamda Note++"))
// 	// 残りのチャンクを追加
// 	for _, chunk := range chunks[1:] {
// 		io.Copy(newFile, chunk)
// 	}
// }

// =======================================================
// ## 3.5.2 PNG ファイルを分析してみる
// import (
// 	"encoding/binary"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func dumpChunk(chunk io.Reader) {
// 	var length int32
// 	binary.Read(chunk, binary.BigEndian, &length)
// 	buffer := make([]byte, 4)
// 	chunk.Read(buffer)
// 	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
// }

// func readChunks(file *os.File) []io.Reader {
// 	// チャンクを格納する配列
// 	var chunks []io.Reader

// 	// 最初の 8 バイトを飛ばす
// 	file.Seek(8, 0)
// 	var offset int64 = 8

// 	for {
// 		var length int32
// 		err := binary.Read(file, binary.BigEndian, &length)
// 		if err == io.EOF {
// 			break
// 		}
// 		chunks = append(chunks,
// 			io.NewSectionReader(file, offset, int64(length)+12))
// 		// 次のチャンクの先頭に移動
// 		// 現在位置は長さを読み終わった箇所なので
// 		// チャンク名（4バイト）+データ長+CRC(4バイト)先に移動
// 		offset, _ = file.Seek(int64(length+8), 1)
// 	}
// 	return chunks
// }

// func main() {
// 	file, err := os.Open("PNG_transparency_demonstration_1.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	chunks := readChunks(file)
// 	for _, chunk := range chunks {
// 		dumpChunk(chunk)
// 	}
// }

// =======================================================
// ## 3.5.2 エンディアン変換
// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// )

// func main() {
// 	data := []byte{0x0, 0x0, 0x27, 0x10}
// 	var i int32
// 	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
// 	fmt.Printf("data: %d\n", i)
// }

// =======================================================
// ## 3.5.1 必要な部位を切り出す io.LimitReader / io.SectionReader
// import (
// 	"io"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := strings.NewReader("Example of io.SectionReader\n")
// 	sectionReader := io.NewSectionReader(reader, 14, 7)
// 	io.Copy(os.Stdout, sectionReader)
// }

// =======================================================
// ## 3.4.2 ファイル入力
// import (
// 	"bufio"
// 	"io"
// 	"net"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "example.com:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
// 	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
// 	defer res.Body.Close()
// 	io.Copy(os.Stdout, res.Body)
// }

// =======================================================
// import (
// 	"io"
// 	"net"
// 	"os"
// )
// func main() {
// 	conn, err := net.Dial("tcp", "example.com:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
// 	io.Copy(os.Stdout, conn)
// }

// =======================================================
// ## 3.4.2 ファイル入力
// import (
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("file.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	io.Copy(os.Stdout, file)
// }

// =======================================================
// ## 3.3.2 入出力関連インタフェースのキャスト
// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	for {
// 		buffer := make([]byte, 5)
// 		size, err := os.Stdin.Read(buffer)
// 		if err == io.EOF {
// 			fmt.Println("EOF")
// 			break
// 		}
// 		fmt.Printf("size=%d inputs=%s\n", size, string(buffer))
// 	}
// }
