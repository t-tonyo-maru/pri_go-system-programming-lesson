package main

// =======================================================
// ## 7.3.2 クライアント側の実装例
// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	fmt.Println("Listen tick server at 224.0.0.1:9999")
// 	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
// 	if err != nil {
// 		panic(err)
// 	}
// 	listener, err := net.ListenMulticastUDP("udp", nil, address)
// 	defer listener.Close()

// 	buffer := make([]byte, 1500)

// 	for {
// 		length, remoteAddress, err := listener.ReadFromUDP(buffer)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("Server %v\n", remoteAddress)
// 		fmt.Printf("Now    %s\n", string(buffer[:length]))
// 	}
// }

// =======================================================
// ## 7.3.1 サーバー側の実装例
// import (
// 	"fmt"
// 	"net"
// 	"time"
// )

// const interval = 10 * time.Second

// func main() {
// 	fmt.Println("Start tick server at 224.0.0.1:9999")
// 	conn, err := net.Dial("udp", "224.0.0.1:9999")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	start := time.Now()
// 	wait := start.Truncate(interval).Add(interval).Sub(start)
// 	time.Sleep(wait)
// 	ticker := time.Tick(interval)
// 	for now := range ticker {
// 		conn.Write([]byte(now.String()))
// 		fmt.Println("Tick: ", now.String())
// 	}
// }

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
