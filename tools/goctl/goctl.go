package main

import (
	"github.com/mikokutou1/go-zero-m/core/load"
	"github.com/mikokutou1/go-zero-m/core/logx"
	"github.com/mikokutou1/go-zero-m/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
