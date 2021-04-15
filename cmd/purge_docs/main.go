package purge_docs

import (
	"bufio"
	"os"
	"strings"
)

const files string = `/usr/share/doc
/usr/share/man                                                        
/usr/share/menu
/usr/share/groff
/usr/share/info
/usr/share/lintian
/usr/share/linda
/usr/share/bug
/usr/share/locale
/usr/share/bash-completion
/var/cache/man`

func Main() {
	s := bufio.NewScanner(strings.NewReader(files))
	for s.Scan() {
		os.RemoveAll(s.Text())
	}
}
