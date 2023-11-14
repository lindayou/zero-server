package dictionary

import (
	"net/http"
	"zero-server/server/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/dictionary"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func EditDicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditDicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewEditDicLogic(r.Context(), svcCtx)
		resp, err := l.EditDic(&req)
		response.Response(w, resp, err)
	}
}
