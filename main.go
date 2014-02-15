package main

import (
	"fmt"
	"os"

	"github.com/dickeyxxx/eight/commands"
	"github.com/dickeyxxx/eight/eight"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

func main() {
  etcdClient := etcd.NewClient([]string{"http://127.0.0.1:4001"})
  client := eight.NewClient(etcdClient)
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		commands.LsCommand(client, fmt.Println),
	}
	app.Run(os.Args)
}
