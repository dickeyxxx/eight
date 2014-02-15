package commands_test

import (
	"github.com/codegangsta/cli"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dickeyxxx/eight/commands"
)

var lines []string

func output(s string) {
	lines = append(lines, s)
}

func getHosts() []string {
	return []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}
}

var _ = Describe("ls", func() {
	var command cli.Command

	BeforeEach(func() {
		command = LsCommand(getHosts, output)
	})

	It("shows a list of hosts", func() {
		c := cli.NewContext(nil, nil, nil)
		command.Action(c)
		Ω(lines).Should(Equal([]string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}))
	})
})
