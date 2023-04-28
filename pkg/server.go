package pkg

import (
	"fmt"
	"log"
	"net"
	"os"
)

// константа значении порта и типа
const (
	CONN_PORT = ":8989"
	CONN_TYPE = "tcp"
)

// функция запуска сервера по умолчанию
func ServerCaseWithDefPort() {
	ls, err := net.Listen(CONN_TYPE, CONN_PORT)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer ls.Close()
	log.Printf("Listening on the port :8989")
	go Broadcaster(&mutex)
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			conn.Close()
			continue
		}
		go HandleConnection(conn, &mutex)
	}
}

// функция запуска сервера собственным портом
func ServerCaseWithPort() {
	arg := os.Args[1]
	port := fmt.Sprintf(":%s", arg)
	ls, err := net.Listen(CONN_TYPE, port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer ls.Close()
	fmt.Printf("Listening on the port %s\n", port)
	go Broadcaster(&mutex)
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			conn.Close()
			continue
		}
		go HandleConnection(conn, &mutex)
	}
}
