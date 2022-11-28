package functions

import (
	"fmt"
	"net/http"

	"github.com/gravitl/netmaker/models"
)

// GetNodeMetrics - fetch a single node's metrics
func GetNodeMetrics(networkName, nodeID string) *models.Metrics {
	return request[models.Metrics](http.MethodGet, fmt.Sprintf("/api/metrics/%s/%s", networkName, nodeID), nil)
}

func GetNetworkNodeMetrics(networkName string) *models.NetworkMetrics {
	return request[models.NetworkMetrics](http.MethodGet, "/api/metrics/"+networkName, nil)
}

func GetAllMetrics() *models.NetworkMetrics {
	return request[models.NetworkMetrics](http.MethodGet, "/api/metrics", nil)
}

func GetNetworkExtMetrics(networkName string) *map[string]models.Metric {
	return request[map[string]models.Metric](http.MethodGet, "/api/metrics-ext/"+networkName, nil)
}
