package http_server

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

func Call(a []byte) {
	const op = "./http-server/client.go"
	var log *slog.Logger
	req, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Error(op, HTTPErr(err))
		os.Exit(1)
	}
	defer req.Close()
	n, err := req.Read(a)
	if err != nil {
		log.Error(op, HTTPErr(err))
	}
	fmt.Println(string(a[:n]))
	return
}
