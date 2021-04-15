package chmod

import (
	"log"
	"os"
)

func Main() {
	if err := os.Chmod(os.Args[1], stringToOctal(os.Args[2])); err != nil {
		log.Fatal(err)
	}
}
