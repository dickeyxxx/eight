package eight

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

func (c *Client) GetDir(key string) []string {
  response, err := c.etcd.Get(key, false, false)
  values := []string{}

  if(err != nil) {
    log.Fatal(err)
  }

  for _, n := range response.Node.Nodes {
    values = append(values, n.Key)
  }

	return values
}
