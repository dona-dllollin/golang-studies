package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect ke server
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Print("Ketik pesan: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// kirim ke server
	fmt.Fprintf(conn, text)

	// baca balasan dari server
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Balasan dari server: %s", message)
}
