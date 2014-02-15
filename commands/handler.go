package commands

import (
	"github.com/coreos/go-etcd/etcd"
)

type handlerFunc func(*etcd.Client)
