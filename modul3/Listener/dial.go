package main

import (
	"bufio"
	"fmt"
	"net"
	)

func main() {
conn, err := net.Dial("tcp", "127.0.0.1:8080")
if err != nil {
fmt.Println(err)
}
fmt.Fprintf(conn, "Hei")
echo, err := bufio.NewReader(conn).ReadString('\n')
fmt.Println(echo)
}
