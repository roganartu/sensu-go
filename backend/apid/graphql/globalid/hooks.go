package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Hooks
//

var hookName = "hooks"

// HookTranslator global ID resource
var HookTranslator = commonTranslator{
	name:       hookName,
	encodeFunc: standardEncoder(hookName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.HookConfig)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(HookTranslator) }
