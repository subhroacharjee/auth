package controller

import (
	healthcheckdata "github.com/subhroacharjee/auth/lib/controller/health_check"
	userdata "github.com/subhroacharjee/auth/lib/controller/user"
	healthcheck "github.com/subhroacharjee/auth/lib/model/health_check"
	"github.com/subhroacharjee/auth/lib/model/user"
	"gorm.io/gorm"
)

type ResolverOptions struct {
	*gorm.DB
}

type ControllerResolver struct {
	HealthCheck healthcheck.Repository
	User        user.Repository
}

func NewController(options ResolverOptions) *ControllerResolver {
	return &ControllerResolver{
		HealthCheck: healthcheckdata.NewHealthCheckController(options),
		User: userdata.NewUserRepository(userdata.RepositoryOptions{
			DB: options.DB,
		}),
	}
}
