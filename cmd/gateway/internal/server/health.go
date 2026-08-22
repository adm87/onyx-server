package server

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
)

type SvcHealthStatus map[string]HealthStatus

func healthzHandler(clients *svcClients, log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		healthStatus := SvcHealthStatus{}
		allHealthy := true

		for svcName, client := range *clients {
			if client == nil || client.Conn() == nil {
				log.Warn("gRPC client is nil or connection is nil", zap.String("service", svcName))
				healthStatus[svcName] = HealthStatusUnhealthy
				allHealthy = false
				continue
			}

			healthy, err := client.IsHealthy(r.Context())
			if err != nil || !healthy {
				if err != nil {
					log.Error("Error checking health of service", zap.String("service", svcName), zap.Error(err))
				}
				healthStatus[svcName] = HealthStatusUnhealthy
				allHealthy = false
			} else {
				healthStatus[svcName] = HealthStatusHealthy
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if allHealthy {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		if err := json.NewEncoder(w).Encode(healthStatus); err != nil {
			log.Error("Failed to encode health status", zap.Error(err))
		}
	}
}
