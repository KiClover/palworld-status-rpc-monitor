package svc

import (
	"github.com/KiClover/palworld-status-rpc-monitor/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
