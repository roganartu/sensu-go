package etcd

import (
	"github.com/coreos/etcd/clientv3"
	types "github.com/sensu/sensu-go/api/core/v2"
)

func namespaceExistsForResource(r types.MultitenantResource) clientv3.Cmp {
	key := getNamespacePath(r.GetNamespace())
	return clientv3.Compare(clientv3.Version(key), ">", 0)
}
