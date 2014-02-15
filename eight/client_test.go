package eight_test

import (
	. "github.com/dickeyxxx/eight/eight"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/coreos/go-etcd/etcd"
)

type fakeEtcdClient struct{}

func (c *fakeEtcdClient) Get(string, bool, bool) (*etcd.Response, error) {
  node := &etcd.Response{
    Node: &etcd.Node{
      Nodes: []etcd.Node{
        etcd.Node{
          Key: "bar",
        },
        etcd.Node{
          Key: "bar",
        },
        etcd.Node{
          Key: "bam",
        },
      },
    },
  }
  return node, nil
}

var _ = Describe("Eight Client", func() {
	var client *Client

	BeforeEach(func() {
		client = NewClient(&fakeEtcdClient{})
	})

	It("returns a whole directory", func() {
		values := client.GetDir("foo")
		Ω(values).Should(Equal([]string{"bar", "bar", "bam"}))
	})
})
