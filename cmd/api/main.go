package main

import (
	"fmt"
	server "offgrid/internal/app"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
