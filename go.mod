module util-buildah

go 1.16

require (
	example.com/chmod v0.0.0-0
	example.com/mkdir v0.0.0-0
	example.com/purge_dirs v0.0.0-0
	example.com/purge_docs v0.0.0-0
	example.com/purge_dpkg v0.0.0-0
	example.com/purge_perl v0.0.0-0
	example.com/purge_sh v0.0.0-0
	example.com/purge_userland v0.0.0-0
	example.com/rm v0.0.0-0
)

replace example.com/chmod => ./cmd/chmod

replace example.com/mkdir => ./cmd/mkdir

replace example.com/rm => ./cmd/rm

replace example.com/purge_docs => ./cmd/purge_docs

replace example.com/purge_dirs => ./cmd/purge_dirs

replace example.com/purge_dpkg => ./cmd/purge_dpkg

replace example.com/purge_perl => ./cmd/purge_perl

replace example.com/purge_sh => ./cmd/purge_sh

replace example.com/purge_userland => ./cmd/purge_userland
