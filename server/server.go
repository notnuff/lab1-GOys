package server

import (
	"fmt"
	"time"
)

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}

func GetTime() {
	currentTime := time.Now()
	fmt.Println("time:", currentTime)
}

const (
	PORT  string = ":8795"
	ROUTE string = "/time"
)
