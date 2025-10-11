package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// listen di TCP port 9000
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Server jalan di port 9000...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accept:", err)
			continue
		}

		// handle tiap koneksi di goroutine biar bisa multi-client
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	fmt.Printf("Pesan dari client: %s", message)

	// balas pesan ke client
	conn.Write([]byte("Halo dari server!\n"))
}
