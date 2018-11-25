package testing

import types "github.com/sensu/sensu-go/api/core/v2"

func (c *MockClient) Health() (*types.HealthResponse, error) {
	args := c.Called()
	return args.Get(0).(*types.HealthResponse), args.Error(1)
}
