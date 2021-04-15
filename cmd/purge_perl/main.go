package purge_perl

import (
	"strings"
	"bufio"
	"os"
	"path/filepath"
)

const wildcard string = `/usr/bin/perl*`
const files string = `/usr/lib/x86_64-linux-gnu/perl-base`

func Main() {
	ws := bufio.NewScanner(strings.NewReader(wildcard))
	for ws.Scan() {
		glob, _ := filepath.Glob(ws.Text())
		for _, f := range glob {
			os.Remove(f)
		}
	}
	s := bufio.NewScanner(strings.NewReader(files))
	for s.Scan() {
		os.RemoveAll(s.Text())
	}
}
