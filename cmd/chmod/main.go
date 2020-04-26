package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Chmod(os.Args[1], stringToOctal(os.Args[2])); err != nil {
		log.Fatal(err)
	}
}
