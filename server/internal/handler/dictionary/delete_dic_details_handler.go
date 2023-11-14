package dictionary

import (
	"net/http"
	"zero-server/server/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/dictionary"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func DeleteDicDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteDicDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewDeleteDicDetailsLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDicDetails(&req)
		response.Response(w, resp, err)
	}
}
