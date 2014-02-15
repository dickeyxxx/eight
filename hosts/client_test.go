package hosts_test

import (
	. "github.com/dickeyxxx/eight/hosts"
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
          Key: "10.0.0.1",
        },
        etcd.Node{
          Key: "10.0.0.2",
        },
        etcd.Node{
          Key: "10.0.0.3",
        },
      },
    },
  }
  return node, nil
}

var _ = Describe("Hosts client", func() {
	var client *Client

	BeforeEach(func() {
		client = NewClient(&fakeEtcdClient{})
	})

	It("returns all the hosts", func() {
		values := client.GetHosts()
		Ω(values).Should(Equal([]string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}))
	})
})
