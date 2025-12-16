package main

import "github.com/rickferrdev/api-basic-v2/internal/server"

func main() {
	if err := server.Server().Run(":8080"); err != nil {
		panic(err)
	}
}
