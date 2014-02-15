package hosts

import (
  "log"
  "github.com/coreos/go-etcd/etcd"
)

type etcdClient interface {
	Get(string, bool, bool) (*etcd.Response, error)
}

type Client struct {
  etcd etcdClient
}

func NewClient(c etcdClient) *Client {
  return &Client{
    etcd: c,
  }
}

func (c *Client) GetHosts() []string {
  response, err := c.etcd.Get("hosts", false, false)
  values := []string{}

  if(err != nil) {
    log.Fatal(err)
  }

  for _, n := range response.Node.Nodes {
    values = append(values, n.Key)
  }

	return values
}
