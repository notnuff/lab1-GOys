package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultPort  = 8795
	defaultRoute = "/time"
)

type TimeServer struct {
	port  int
	route string
}

func (t TimeServer) StartServer() {
	http.HandleFunc(t.route, t.TimeHandler)

	fmt.Printf("Server started at port: %d", t.port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", t.port), nil)
	if err != nil {
		fmt.Printf("Failed to start server at %d", t.port)
		return
	}
}

func (t TimeServer) TimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := getFormattedTime()
	jsonResponse, err := dateToJsonT(currentTime)
	if err != nil {
		fmt.Println("Failed to create json response:", err)
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Failed to write response:", err)
		return
	}
}

func CreateDefaultTimerServer() TimeServer {
	return TimeServer{
		defaultPort,
		defaultRoute,
	}
}

// some useful functions for our server, that can be reused
func getFormattedTime() string {
	t := time.Now().Format(time.RFC3339)
	return t
}

func dateToJsonT(t string) (res []byte, err error) {
	timeMap := map[string]string{"time": t}
	res, err = json.MarshalIndent(timeMap, "", "  ")
	return
}
