package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// Remove first arg
	a := os.Args
	a = append(a[:0], a[1:]...)

	// Make sure message has a forward slash
	cmd := "/" + strings.TrimPrefix(strings.Join(a, " "), "/")
	log.Printf("cmd is \"%s\"\n", cmd)
}
