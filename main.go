package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")
	

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
