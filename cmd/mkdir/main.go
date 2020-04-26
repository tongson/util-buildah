package main

import "os"

func main() {
	if len(os.Args) == 2 {
		os.MkdirAll(os.Args[1], 0755)
	} else {
		os.MkdirAll(os.Args[1], stringToOctal(os.Args[2]))
	}
}
