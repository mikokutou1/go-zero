package main

import (
	"flag"

	"github.com/mikokutou1/go-zero-m/core/conf"
	"github.com/mikokutou1/go-zero-m/gateway"
)

var configFile = flag.String("f", "etc/gateway.yaml", "config file")

func main() {
	flag.Parse()

	var c gateway.GatewayConf
	conf.MustLoad(*configFile, &c)
	gw := gateway.MustNewServer(c)
	defer gw.Stop()
	gw.Start()
}
