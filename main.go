package main

import (
	"lab1-GOys/server"
)

func main() {
	timeServer := server.CreateDefaultTimerServer()
	timeServer.StartServer()
}
