package user

import (
	"context"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserListResp)
	resp.Page = req.PageNum
	resp.PageSize = req.PageSize
	list, err, count := l.svcCtx.UserModel.UserList(l.ctx, req.PageNum, req.PageSize)
	resp.Total = int64(count)

	userList := make([]*types.UserMessage, 0)
	for _, user := range list {
		//获取角色权限
		authorityUsers, err := l.svcCtx.Authority.GetAuthorityUser(l.ctx, user.Id)
		if err != nil {
			return nil, err
		}
		var Authotities = []*types.Authority{}
		for _, authorityUser := range authorityUsers {
			Authotities = append(Authotities, &types.Authority{
				CreateAt:      authorityUser.CreatedAt.Time.Unix(),
				UpdateAt:      authorityUser.UpdatedAt.Time.Unix(),
				AuthorityId:   authorityUser.AuthorityId,
				AuthorityName: authorityUser.AuthorityName,
				ParentId:      authorityUser.ParentId,
			})

		}
		userList = append(userList, &types.UserMessage{
			Id:          user.Id,
			Username:    user.Username,
			Phone:       user.Phone,
			CreateTime:  user.CreateTime.Unix(),
			UpdateTime:  user.UpdateAt.Unix(),
			AuthorityId: user.AuthorityId.Int64,
			Uuid:        user.Uuid,
			Authorities: Authotities,
		})
	}
	resp.UserList = userList
	return resp, nil
}
