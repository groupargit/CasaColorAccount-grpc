package handler

import (
	"fmt"
	"net/http"

	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/logic"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/svc"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Println("req=////==:", req)
		l := logic.NewCreateLoggerLogic(r.Context(), svcCtx)
		err := l.CreateLogger(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, "ok")
			httpx.Ok(w)
		}
	}
}
