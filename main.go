package main

import (
	"os"
	"fmt"

	"github.com/dickeyxxx/eight/commands"
	"github.com/dickeyxxx/eight/hosts"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		commands.LsCommand(hosts.List, fmt.Println),
	}
	app.Run(os.Args)
}
