package user

import (
	"context"
	"errors"
	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
	"zero-server/server/model/user_model"
	"zero-server/server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequset) (resp *types.LoginResponse, err error) {
	//根据用户名查找密码并比对
	findUser := new(user_model.SysUser)
	findUser.Username = req.Username
	user, err := l.svcCtx.UserModel.FindByUsername(l.ctx, findUser)
	if user == nil {
		err = errors.New("用户名不存在，请先注册")
		return nil, err
	}
	if !utils.BcryptCheck(req.Password, user.Password) {
		err = errors.New("密码错误")
		return nil, err
	}
	token, err := utils.GenerateToken(int(user.Id), user.Uuid, user.Username, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Id = int(user.Id)
	resp.Token = token
	return resp, nil
}
