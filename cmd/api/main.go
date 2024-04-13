package main

import (
	"fmt"

	"damstudy-backend/internal/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	} else {
		fmt.Println("Server started successfully")
	}
}
