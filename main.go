package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	time.Local = time.FixedZone("CST", 3600*8)
}

func main() {
	app := cli.NewApp()
	app.Name = "Nexus docker image clean CLI"
	app.Usage = "Manage Docker Private Registry on Nexus"
	app.Version = "1.0.0-beta"
	app.Authors = []*cli.Author{
		{
			Name:  "XMapst",
			Email: "xmapst@gmail.com",
		},
	}
	app.Commands = Command()
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
