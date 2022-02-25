package main

import (
	"github.com/elton/project-layout/app/myapp/internal/pkg/server"
	"github.com/elton/project-layout/app/myapp/internal/service"
)

func main() {
	greeting, err := service.Hello()
	if err != nil {
		panic(err)
	}
	println(greeting)
	server.Start()
}
