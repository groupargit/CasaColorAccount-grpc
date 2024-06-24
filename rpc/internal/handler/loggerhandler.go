package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/logic"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/svc"
	"github.com/groupargit/casacoloraccount-grpc/rpc/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type User struct {
}

func LoggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request

		// Leer el cuerpo de la solicitud una vez
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			httpx.Error(w, err)
			return
		}

		// Mostrar el cuerpo de la solicitud
		fmt.Println("request Body:", string(body))

		// Parsear el cuerpo de la solicitud
		err = json.Unmarshal(body, &req)
		if err != nil {
			fmt.Printf("Error parsing body: %v", err)
			httpx.Error(w, err)
			return
		}

		// Crear y llamar a la lógica del logger
		l := logic.NewCreateLoggerLogic(r.Context(), svcCtx)
		err = l.CreateLogger(body)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, json.RawMessage(`{"msg":"ok"}`))
		}
	}
}
