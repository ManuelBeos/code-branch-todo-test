package main

import (
	"log"

	"github.com/manuelbeos/code-branch-todo-test/internal/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("recovered from panic: ", r)
		}
	}()

	myServer := server.NewServer()
	if err := myServer.Run(); err != nil {
		log.Fatal(err)
	}
}
