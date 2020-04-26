package main

import "os"

func main() {
	os.RemoveAll(os.Args[1])
}
