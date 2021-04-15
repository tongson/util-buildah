package mkdir

import (
	"os"
	"log"
	"strconv"
)

func stringToOctal(m string) os.FileMode {
	octal, err := strconv.ParseInt(m, 8, 32)
	if err != nil {
		log.Fatal(err)
	}
	return os.FileMode(octal)
}
