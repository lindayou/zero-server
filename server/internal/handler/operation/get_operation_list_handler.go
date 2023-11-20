package operation

import (
	"net/http"
	"zero-server/server/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/operation"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func GetOperationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOperationListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operation.NewGetOperationListLogic(r.Context(), svcCtx)
		resp, err := l.GetOperationList(&req)
		response.Response(w, resp, err)
	}
}
