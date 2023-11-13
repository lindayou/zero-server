package dictionary

import (
	"context"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDicListLogic {
	return &GetDicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDicListLogic) GetDicList(req *types.GetDicListReq) (resp *types.GetDicListResp, err error) {
	resp = new(types.GetDicListResp)
	DicList := make([]*types.Dictionary, 0)
	list, err := l.svcCtx.DicModel.GetDicList(l.ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		DicList = append(DicList, &types.Dictionary{
			Id:        int(item.Id),
			CreatedAt: int(item.CreatedAt.Unix()),
			UpdatedAt: int(item.UpdatedAt.Unix()),
			//DeletedAt: item.,
			Name:   item.Name,
			Type:   item.Type,
			Status: item.Status,
			Desc:   item.Desc,
		})

	}
	resp.DicList = DicList
	return resp, nil
}
