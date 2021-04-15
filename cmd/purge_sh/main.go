package purge_sh

import (
	"bufio"
	"os"
	"strings"
)

const files string = `/XXXXX
/bin/dash
/bin/sh
/XXXXX`

func Main() {
	s := bufio.NewScanner(strings.NewReader(files))
	for s.Scan() {
		os.Remove(s.Text())
	}
}
