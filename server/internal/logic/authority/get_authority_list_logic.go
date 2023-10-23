package authority

import (
	"context"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuthorityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorityListLogic {
	return &GetAuthorityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthorityListLogic) GetAuthorityList(req *types.GetAuthorityListReq) (resp *types.GetAuthorityListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetAuthorityListResp)
	//查询到所有用户
	users, err := l.svcCtx.Authority.GetAllAuthorities(l.ctx)
	if err != nil {
		return nil, err
	}
	//获取到parentid =0 的作为一级结构
	treeMap := make(map[int64][]*types.GetAuthorityList)
	for _, user := range users {
		treeMap[user.ParentId] = append(treeMap[user.ParentId], &types.GetAuthorityList{
			AuthorityId:   user.AuthorityId,
			AuthorityName: user.AuthorityName,
			ParentId:      user.ParentId,
		})
	}
	respLists := make([]*types.GetAuthorityList, 0)
	for _, treeMap := range treeMap[0] {
		respLists = append(respLists, &types.GetAuthorityList{

			AuthorityId:   treeMap.AuthorityId,
			AuthorityName: treeMap.AuthorityName,
			ParentId:      treeMap.ParentId,
		})
	}
	//传入一节结构和所有的用户循环放入
	for _, respList := range respLists {
		getUserTree(respList, treeMap)
	}
	resp.AuthorityList = respLists

	return
}

func getUserTree(parent *types.GetAuthorityList, treeMap map[int64][]*types.GetAuthorityList) {
	parent.Children = treeMap[parent.AuthorityId]
	for _, item := range parent.Children {
		getUserTree(item, treeMap)
	}

}
