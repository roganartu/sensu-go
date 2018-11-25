package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Users
//

var userName = "users"

// UserTranslator global ID resource
var UserTranslator = commonTranslator{
	name:       userName,
	encodeFunc: standardEncoder(userName, "Username"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.User)
		return ok
	},
}

// Register user encoder/decoder
func init() { registerTranslator(UserTranslator) }
