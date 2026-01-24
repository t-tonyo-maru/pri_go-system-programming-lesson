package main

// =======================================================
// ## 7.2.2 クライアント側の実装例
// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	conn, err := net.Dial("udp4", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	fmt.Println("Sending to server")

// 	_, err = conn.Write([]byte("Hello from Client"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Receiving from server")
// 	buffer := make([]byte, 1500)
// 	length, err := conn.Read(buffer)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Received: %s\n", string(buffer[:length]))
// }

// =======================================================
// ## 7.2.1 サーバー側の実装例
// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	fmt.Println("Server is running at localhost:8888")
// 	conn, err := net.ListenPacket("udp", "localhost:8888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	buffer := make([]byte, 1500)
// 	for {
// 		length, remoteAddress, err := conn.ReadFrom(buffer)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("Reveived from %v: %v\n", remoteAddress, string(buffer[:length]))

// 		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)

// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
