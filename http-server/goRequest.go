package http_server

import (
	"fmt"
	"log/slog"
	"net/http"
)

func req() {
	var l *slog.Logger
	resp, err := http.Get("https://golang.org")
	if err != nil {
		l.Error("ERR", HTTPErr(err))
	}
	fmt.Println(resp)
	defer resp.Body.Close()

}
