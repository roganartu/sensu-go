package globalid

import (
	"testing"

	types "github.com/sensu/sensu-go/api/core/v2"
	"github.com/stretchr/testify/assert"
)

func TestStandardDecoder(t *testing.T) {
	assert := assert.New(t)

	handler := types.FixtureHandler("myHandler")
	encoderFn := standardEncoder("handlers", "Name")
	components := encoderFn(handler)

	assert.Equal("handlers", components.Resource())
	assert.Equal("default", components.Namespace())
	assert.Equal("myHandler", components.UniqueComponent())
}
