package purge_dpkg

import (
	"strings"
	"bufio"
	"os"
	"path/filepath"
)

const files string = `/XXXXX
/usr/bin/dpkg
/usr/bin/dpkg-deb
/usr/bin/dpkg-divert
/usr/bin/dpkg-maintscript-helper
/usr/bin/dpkg-query
/usr/bin/dpkg-split
/usr/bin/dpkg-statoverride
/usr/bin/dpkg-trigger
/usr/bin/update-alternatives
/usr/share/dpkg
/etc/dpkg
/usr/lib/dpkg
/XXXXX
/usr/bin/apt
/usr/bin/apt-cache
/usr/bin/apt-cdrom
/usr/bin/apt-config
/usr/bin/apt-get
/usr/bin/apt-key
/usr/bin/apt-mark
/usr/lib/apt
/usr/bin/debsig-verify
/sbin/start-stop-daemon
/etc/apt
/XXXXX
/usr/bin/deb-systemd-helper
/usr/bin/deb-systemd-invoke
/usr/sbin/invoke-rc.d
/usr/sbin/service
/usr/sbin/update-rc.d
/XXXXX
/usr/bin/gpgv
/XXXXX
/bin/run-parts
/bin/tempfile
/bin/which
/sbin/installkernel
/usr/bin/ischroot
/usr/bin/savelog
/usr/sbin/add-shell
/usr/sbin/remove-shell
/usr/share/debianutils/shells
/XXXXX
/etc/apt/apt.conf.d/01autoremove
/etc/cron.daily/apt-compat
/etc/kernel/postinst.d/apt-auto-removal
/etc/logrotate.d/apt
/lib/systemd/system/apt-daily-upgrade.service
/lib/systemd/system/apt-daily-upgrade.timer
/lib/systemd/system/apt-daily.service
/lib/systemd/system/apt-daily.timer
/usr/bin/apt
/usr/bin/apt-cache
/usr/bin/apt-cdrom
/usr/bin/apt-config
/usr/bin/apt-get
/usr/bin/apt-key
/usr/bin/apt-mark
/usr/lib/apt/apt-helper
/usr/lib/apt/apt.systemd.daily
/usr/lib/apt/methods/cdrom
/usr/lib/apt/methods/copy
/usr/lib/apt/methods/file
/usr/lib/apt/methods/ftp
/usr/lib/apt/methods/gpgv
/usr/lib/apt/methods/http
/usr/lib/apt/methods/https
/usr/lib/apt/methods/mirror
/usr/lib/apt/methods/mirror+copy
/usr/lib/apt/methods/mirror+file
/usr/lib/apt/methods/mirror+ftp
/usr/lib/apt/methods/mirror+http
/usr/lib/apt/methods/mirror+https
/usr/lib/apt/methods/rred
/usr/lib/apt/methods/rsh
/usr/lib/apt/methods/ssh
/usr/lib/apt/methods/store
/usr/lib/apt/planners/dump
/usr/lib/apt/solvers/dump
/usr/lib/dpkg/methods/apt/desc.apt
/usr/lib/dpkg/methods/apt/install
/usr/lib/dpkg/methods/apt/names
/usr/lib/dpkg/methods/apt/setup
/usr/lib/dpkg/methods/apt/update
/usr/lib/s390x-linux-gnu/libapt-private.so.0.0
/usr/lib/s390x-linux-gnu/libapt-private.so.0.0.0
/usr/share/bash-completion/completions/apt
/XXXXX
/etc/apt/apt.conf.d/70debconf
/etc/debconf.conf
/usr/bin/debconf
/usr/bin/debconf-apt-progress
/usr/bin/debconf-communicate
/usr/bin/debconf-copydb
/usr/bin/debconf-escape
/usr/bin/debconf-set-selections
/usr/bin/debconf-show
/usr/sbin/dpkg-preconfigure
/usr/sbin/dpkg-reconfigure
/usr/share/debconf
/usr/share/perl5/Debconf
/usr/share/perl5/Debian
/usr/share/pixmaps/debian-logo.pngG
/XXXXX`

const wildcard string = `/usr/lib/x86_64-linux-gnu/libapt-*
`

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
