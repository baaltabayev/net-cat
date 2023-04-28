package main

import (
	"fmt"
	"net-cat/pkg"
	"os"
	"sync"
)

var mutex sync.Mutex

func main() {
	input := os.Args
	switch len(input) {
	case 1:
		pkg.ServerCaseWithDefPort()
	case 2:
		pkg.ServerCaseWithPort()
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
}
