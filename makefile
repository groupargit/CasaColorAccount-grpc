BUILDPATH=$(CURDIR)
API_NAME=accounts
# external services required to run the service

ctl:
	@cd ./rpc && goctl rpc protoc  account_api.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

go-ctl:
	@cd ./rpc && go run accountapi.go -f etc/accountapi.yaml

pg:
	@rancher kubectl -n ccp port-forward svc/postgresql-payment 6432:5432
mongo:
	@rancher kubectl -n ccp port-forward svc/mongodb-applications 27018:27017

coverage:
	@echo "Coverfile... required 60% coverage"
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -func coverfile_out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > coverage.txt

test:
	@echo "Running tests..."
	@go test -v --coverprofile=coverage.out ./... ./...

.PHONY: test