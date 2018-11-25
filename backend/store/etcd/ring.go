package etcd

import (
	"github.com/sensu/sensu-go/backend/ring"
	types "github.com/sensu/sensu-go/api/core/v2"
)

// GetRing gets a named Ring.
func (s *Store) GetRing(path ...string) types.Ring {
	return ring.EtcdGetter{Client: s.client}.GetRing(path...)
}
