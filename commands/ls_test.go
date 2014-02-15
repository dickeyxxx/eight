package commands_test
import (
  "github.com/codegangsta/cli"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "github.com/dickeyxxx/eight/commands"
)

var hosts = []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}

var lines []string

func output(args ...interface{}) (int, error) {
  for _, s := range args {
    lines = append(lines, s.(string))
  }
  return 0, nil
}

type fakeClient struct{}

func (fakeClient) GetDir(string) []string {
  return hosts
}

var _ = Describe("ls", func() {
  var command cli.Command

  BeforeEach(func() {
    client := fakeClient{}
    command = LsCommand(client, output)
  })

  It("shows a list of hosts", func() {
    c := cli.NewContext(nil, nil, nil)
    command.Action(c)
    Ω(lines).Should(Equal(hosts))
  })
})
