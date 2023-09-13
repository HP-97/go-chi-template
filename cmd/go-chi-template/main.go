package main

import (
	"fmt"

	"example.com/go_chantest/internal/config"
)

func main() {
	config.InitConfig()
	fmt.Printf("hello world!\n")
}
