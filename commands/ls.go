package commands

import (
	"github.com/codegangsta/cli"
)

type getHostsFn func() []string
type outputFn func(string)

func LsCommand(getHosts getHostsFn, output outputFn) cli.Command {
	return cli.Command{
		Name:  "ls",
		Usage: "retrieve the hosts",
		Action: func(c *cli.Context) {
			handleLs(getHosts, output)
		},
	}
}

func handleLs(getHosts getHostsFn, output outputFn) {
	for _, host := range getHosts() {
		output(host)
	}
}
