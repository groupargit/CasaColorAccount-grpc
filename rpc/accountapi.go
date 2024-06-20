package main

import (
	"flag"
	"os"

	model "github.com/groupargit/casacoloraccount-grpc/model/mongo/nocache"
	"github.com/groupargit/casacoloraccount-grpc/pkg/secrets"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/config"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/constants"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/handler"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/svc"
	"github.com/groupargit/casacoloraccount-grpc/rpc/pkg/bylogger"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/accountapi.yaml", "the config file")

func main() {

	secrets.LoadSecrets()
	flag.Parse()
	var (
		mongo_uri  = os.Getenv("MONGODB_URI")
		db_name    = os.Getenv("MONGODB_DB_NAME")
		collection = constants.COLLECTION
	)

	bylogger.LogInfo(constants.MSG_CONNECT_DB, db_name)

	db := model.NewAccountModel(mongo_uri, db_name, collection)
	bylogger.LogInfo(constants.MSG_START)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c, db)

	// s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
	// 	v1alpha1.RegisterAccountAPIServiceServer(grpcServer, server.NewAccountAPIServiceServer(ctx))

	// 	reflection.Register(grpcServer)

	// })
	// defer s.Stop()

	// bylogger.LogInfo(constants.START_SERVER, c.ListenOn)
	// s.Start()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	bylogger.LogInfo(constants.START_SERVER, c.Host, c.Port)
	server.Start()
}
