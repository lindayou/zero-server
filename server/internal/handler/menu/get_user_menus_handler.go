package menu

import (
	"net/http"
	"zero-server/server/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/menu"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func GetUserMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserMenusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetUserMenusLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMenus(&req)
		response.Response(w, resp, err)
	}
}
