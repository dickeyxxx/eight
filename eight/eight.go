package eight

import (
  "github.com/coreos/go-etcd/etcd"
)

type EtcdClient interface{
  Get(string, bool, bool) (*etcd.Response, error)
}
