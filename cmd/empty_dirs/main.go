package main

import "os"

func main() {
	os.RemoveAll("/tmp")
	os.MkdirAll("/tmp", 01777)
	os.RemoveAll("/var/tmp")
	os.MkdirAll("/var/tmp", 01777)
	os.RemoveAll("/var/log")
	os.MkdirAll("/var/log", 0755)
	os.RemoveAll("/var/cache")
	os.MkdirAll("/var/cache", 0755)
}
