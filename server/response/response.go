package response

import (
	"context"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
)

func Response(w http.ResponseWriter, resp interface{}, err error) {
	ctx := context.Background()
	if err != nil {
		xhttp.JsonBaseResponseCtx(ctx, w, err)
	} else {
		xhttp.JsonBaseResponseCtx(ctx, w, resp)
	}

}
