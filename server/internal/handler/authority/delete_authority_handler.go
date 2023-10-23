package authority

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/authority"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func DeleteAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authority.NewDeleteAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAuthority(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
