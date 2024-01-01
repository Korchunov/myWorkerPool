package main

import (
	"context"
	"time"
)

const (
	INTERVAL     = time.Second * 5
	COUNT_WORKER = 15
)

func main() {
	ctx := context.Background()
	userTasks := []string{
		"https://workshop.zhashkevych.com/",
		"https://golang-ninja.com/",
		"https://zhashkevych.com/",
		"https://google.com/",
		"https://golang.org/",
	}
	result :=
}
