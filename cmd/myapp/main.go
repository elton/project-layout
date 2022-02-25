package main

import "github.com/elton/project-layout/internal/app/myapp/hello"

func main() {
	greeting, err := hello.Hello()
	if err != nil {
		panic(err)
	}
	println(greeting)
}
