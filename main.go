package main

import server "github.com/davidleitw/chronos/pkg/server"

func main() {
	server := server.NewServer()
	_ = server.Run()
}
