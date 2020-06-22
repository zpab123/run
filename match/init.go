package main

import (
	"github.com/zpab123/sco"
)

//
func init() {
	app := sco.GetApp()
	app.Options.Name = "match_1v1"
	app.Options.Mid = 202
	app.Options.Cluster = true
	app.Options.RpcServer.Laddr = "192.168.1.88:4036"

	app.Options.Discovery.Endpoints = []string{
		"http://192.168.1.69:2379",
		"http://192.168.1.69:2479",
		"http://192.168.1.69:2579",
	}

	app.SetHandler(&Hander{})
}
