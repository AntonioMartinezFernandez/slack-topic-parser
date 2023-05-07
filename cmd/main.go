package main

import (
	"os"
	"os/exec"
	http_server "slack-topic-parser/internal/http-server"
)

func main() {
	// Clears the screen
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// Start HTTP Server
	http_server.Start()
}
