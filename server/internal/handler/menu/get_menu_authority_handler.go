package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/menu"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func GetMenuAuthorityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMenuAuthorityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetMenuAuthorityLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuAuthority(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
