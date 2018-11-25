package client

import (
	"encoding/json"
	"fmt"

	types "github.com/sensu/sensu-go/api/core/v2"
)

const healthPath = "/health"

func (c *RestClient) Health() (*types.HealthResponse, error) {
	res, err := c.R().Get(healthPath)
	if err != nil {
		return nil, fmt.Errorf("GET %q: %s", healthPath, err)
	}
	var healthResponse *types.HealthResponse
	return healthResponse, json.Unmarshal(res.Body(), &healthResponse)
}
