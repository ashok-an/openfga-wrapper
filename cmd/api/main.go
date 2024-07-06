package main

import (
	"fmt"

	"github.com/ashok-an/openfga-wrapper/internal/server"
)

func main() {

	server := server.NewServer()
	fmt.Println("Starting server at ", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
