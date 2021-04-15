package main

import (
	chmod "example.com/chmod"
	mkdir "example.com/mkdir"
	purge_dirs "example.com/purge_dirs"
	purge_docs "example.com/purge_docs"
	purge_dpkg "example.com/purge_dpkg"
	purge_perl "example.com/purge_perl"
	purge_sh "example.com/purge_sh"
	purge_userland "example.com/purge_userland"
	rm "example.com/rm"
	"fmt"
	"os"
)

var commands = map[string]func(){
	"chmod":          chmod.Main,
	"mkdir":          mkdir.Main,
	"rm":             rm.Main,
	"purge-dirs":     purge_dirs.Main,
	"purge_dirs":     purge_dirs.Main,
	"purge-docs":     purge_docs.Main,
	"purge_docs":     purge_docs.Main,
	"purge-dpkg":     purge_dpkg.Main,
	"purge_dpkg":     purge_dpkg.Main,
	"purge-perl":     purge_perl.Main,
	"purge_perl":     purge_perl.Main,
	"purge-sh":       purge_sh.Main,
	"purge_sh":       purge_sh.Main,
	"purge-userland": purge_userland.Main,
	"purge_userland": purge_userland.Main,
}

func main() {
	cmd, ok := commands[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown command %q\n", os.Args[1])
		os.Exit(1)
	}
	os.Args = append(os.Args[:1], os.Args[1+1:]...)
	cmd()
}
