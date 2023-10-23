package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-server/server/internal/logic/user"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
)

func SendIdentityMesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendIdentityMesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendIdentityMesLogic(r.Context(), svcCtx)
		resp, err := l.SendIdentityMes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
