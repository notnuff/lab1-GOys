package server

import (
	"fmt"
	"time"
)

func GetTime() {
	currentTime := time.Now()
	fmt.Println("time:", currentTime)
}

const (
	PORT  string = ":8795"
	ROUTE string = "/time"
)

const WHADAFQ = "idk what i`m doing (just a joke)"
