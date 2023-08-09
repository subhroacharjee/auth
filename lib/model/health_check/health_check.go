package healthcheck

import (
	"net/http"
)

type HealthCheck struct {
	Status string `json:"status"`
}

type Repository interface {
	HeathCheckHandler(w http.ResponseWriter, r *http.Request)
}
