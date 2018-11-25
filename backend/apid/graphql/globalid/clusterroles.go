package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Cluster Roles
//
var clusterRoleName = "clusterroles"

// ClusterRoleTranslator global ID resource
var ClusterRoleTranslator = commonTranslator{
	name:       clusterRoleName,
	encodeFunc: standardEncoder(clusterRoleName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.ClusterRole)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(ClusterRoleTranslator) }
