package main

import (
	"flag"
	"os"

	model "github.com/groupargit/casacoloraccount-grpc/model/mongo/nocache"
	"github.com/groupargit/casacoloraccount-grpc/pkg/secrets"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/config"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/constants"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/handler"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/server"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/svc"
	"github.com/groupargit/casacoloraccount-grpc/rpc/pkg/bylogger"
	"github.com/groupargit/casacoloraccount-grpc/rpc/types/account/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/accountapi.yaml", "the config file")

func main() {
	// Load secrets and parse flags
	secrets.LoadSecrets()
	flag.Parse()

	// Retrieve MongoDB connection information from environment variables
	mongo_uri := os.Getenv("MONGODB_URI")
	db_name := os.Getenv("MONGODB_DB_NAME")
	collection := constants.COLLECTION

	// Initialize MongoDB connection and context
	db := model.NewAccountModel(mongo_uri, db_name, collection)
	cfg := getServiceConfig()
	ctx := svc.NewServiceContext(*cfg, db)

	// Start gRPC server in a goroutine
	go func() {
		startGRPCServer(ctx)
	}()

	// Start REST server
	startRESTServer(ctx)
}

// Function to start gRPC server
func startGRPCServer(ctx *svc.ServiceContext) {
	var g config.ConfigGrpc
	g.Name = "account.rpc"
	g.ListenOn = ":50050"
	conf.MustLoad(*configFile, &g)

	s := zrpc.MustNewServer(g.RpcServerConf, func(grpcServer *grpc.Server) {
		v1alpha1.RegisterAccountAPIServiceServer(grpcServer, server.NewAccountAPIServiceServer(ctx))

		if ctx.Config.Mode == service.DevMode || ctx.Config.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	bylogger.LogInfo(constants.START_SERVER, g.ListenOn)
	s.Start()
}

// Function to start REST server
func startRESTServer(ctx *svc.ServiceContext) {
	c := getServiceConfig()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	bylogger.LogInfo(constants.START_SERVER, c.Host, c.Port)
	server.Start()
}

// Function to load service configuration
func getServiceConfig() *config.Config {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	return &c
}

// func main() {

// 	secrets.LoadSecrets()
// 	flag.Parse()
// 	var (
// 		mongo_uri  = os.Getenv("MONGODB_URI")
// 		db_name    = os.Getenv("MONGODB_DB_NAME")
// 		collection = constants.COLLECTION
// 	)

// 	bylogger.LogInfo(constants.MSG_CONNECT_DB, db_name)

// 	db := model.NewAccountModel(mongo_uri, db_name, collection)
// 	bylogger.LogInfo(constants.MSG_START)

// 	var c config.Config
// 	conf.MustLoad(*configFile, &c)
// 	ctx := svc.NewServiceContext(c, db)

// 	var g config.ConfigGrpc
// 	g.Name = "account.rpc"
// 	g.ListenOn = ":50050"
// 	conf.MustLoad(*configFile, &c)

// 	s := zrpc.MustNewServer(g.RpcServerConf, func(grpcServer *grpc.Server) {
// 		v1alpha1.RegisterAccountAPIServiceServer(grpcServer, server.NewAccountAPIServiceServer(ctx))

// 		if c.Mode == service.DevMode || c.Mode == service.TestMode {
// 			reflection.Register(grpcServer)
// 		}
// 	})
// 	defer s.Stop()

// 	bylogger.LogInfo(constants.START_SERVER, g.ListenOn)
// 	s.Start()

// 	server := rest.MustNewServer(c.RestConf)
// 	defer server.Stop()

// 	handler.RegisterHandlers(server, ctx)

// 	bylogger.LogInfo(constants.START_SERVER, c.Host, c.Port)
// 	server.Start()
// }
