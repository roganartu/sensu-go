package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Mutators
//

var mutatorName = "mutators"

// MutatorTranslator global ID resource
var MutatorTranslator = commonTranslator{
	name:       mutatorName,
	encodeFunc: standardEncoder(mutatorName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.Mutator)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(MutatorTranslator) }
