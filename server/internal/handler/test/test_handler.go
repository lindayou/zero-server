package test

import (
	"net/http"
	"zero-server/server/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/test"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := test.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		response.Response(w, resp, err)
	}
}
