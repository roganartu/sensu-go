package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Asset
//

var assetName = "assets"

// AssetTranslator global ID resource
var AssetTranslator = commonTranslator{
	name:       assetName,
	encodeFunc: standardEncoder(assetName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.Asset)
		return ok
	},
}

// Register asset encoder/decoder
func init() { registerTranslator(AssetTranslator) }
