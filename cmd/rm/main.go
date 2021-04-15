package rm

import "os"

func Main() {
	os.RemoveAll(os.Args[1])
}
