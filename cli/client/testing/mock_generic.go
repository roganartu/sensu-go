package testing

import types "github.com/sensu/sensu-go/api/core/v2"

// PutGeneric ...
func (c *MockClient) PutResource(r types.Resource) error {
	args := c.Called(r)
	return args.Error(0)
}
