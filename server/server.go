package server

import (
	"fmt"
	"net/http"
)

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
