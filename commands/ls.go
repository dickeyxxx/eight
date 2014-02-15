package commands

import (
	"github.com/codegangsta/cli"
)

type clientInterface interface {
	GetHosts() []string
}
type outputFn func(...interface{}) (int, error)

func LsCommand(client clientInterface, output outputFn) cli.Command {
	return cli.Command{
		Name:  "ls",
		Usage: "retrieve the hosts",
		Action: func(c *cli.Context) {
			for _, host := range client.GetHosts() {
				output(host)
			}
		},
	}
}
