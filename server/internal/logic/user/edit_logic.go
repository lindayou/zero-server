package user

import (
	"context"
	"zero-server/server/model/user_model"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.EditUserRequest) (resp *types.EditUserResponse, err error) {
	// todo: add your logic here and delete this line
	user := new(user_model.SysUser)
	user.Username = req.Username
	user.Email = req.Email
	user.Phone = req.Phone
	user.Id = req.Id
	err = l.svcCtx.UserModel.Update(l.ctx, user)

	return
}
