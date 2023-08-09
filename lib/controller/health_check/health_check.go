package healthcheckdata

import (
	"encoding/json"
	"net/http"

	healthcheck "github.com/subhroacharjee/auth/lib/model/health_check"
)

type HealthCheckRepositoryOptions interface{}

type HealthCheckRepositoryImpl struct {
	HealthCheckRepositoryOptions
}

func NewHealthCheckController(options HealthCheckRepositoryOptions) *HealthCheckRepositoryImpl {
	return &HealthCheckRepositoryImpl{
		options,
	}
}

func (h HealthCheckRepositoryImpl) HeathCheckHandler(w http.ResponseWriter, r *http.Request) {
	hc := healthcheck.HealthCheck{
		Status: "working good!",
	}
	json.NewEncoder(w).Encode(hc)
}
