package http_server

import (
	"fmt"
	"log/slog"
	"net"
)

func server() {
	const op = "./http-server/server.go"
	var log *slog.Logger
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Error(op, HTTPErr(err))
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		log.Error(op, HTTPErr(err))
	}
	wr, err := conn.Write([]byte("message"))
	fmt.Println(wr)
	return
}
