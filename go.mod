module github.com/sonnes/photos

go 1.22

toolchain go1.22.2

replace (
	github.com/xdb-dev/xdb => ../xdb
	github.com/xdb-dev/xdb/stores/sqlite => ../xdb/stores/sqlite
)

require (
	github.com/karrick/godirwalk v1.17.0
	github.com/lmittmann/tint v1.0.4
	github.com/sourcegraph/conc v0.3.0
	github.com/urfave/cli/v2 v2.27.1
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
)
