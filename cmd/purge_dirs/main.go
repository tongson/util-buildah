package purge_dirs

import (
	"os"
	"syscall"
)

func Main() {
	syscall.Umask(0)
	os.RemoveAll("/tmp")
	os.MkdirAll("/tmp", 01777)
	os.RemoveAll("/var/tmp")
	os.MkdirAll("/var/tmp", 01777)
	os.RemoveAll("/var/log")
	os.MkdirAll("/var/log", 0755)
	os.RemoveAll("/var/cache")
	os.MkdirAll("/var/cache", 0755)
}
