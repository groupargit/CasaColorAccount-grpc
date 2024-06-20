package svc

import (
	model "github.com/groupargit/casacoloraccount-grpc/model/mongo/nocache"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/config"
	//"k8s.io/client-go/kubernetes"
)

type ServiceContext struct {
	Config config.Config
	DB     model.AccountModel
}

func NewServiceContext(c config.Config, db model.AccountModel) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

