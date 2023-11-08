package user

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"zero-server/server/model/user_auth"
	"zero-server/server/model/user_model"
	"zero-server/server/utils"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	//首先验证验证码
	//code := req.Code
	// todo: 从redis中获取验证码
	//realCode, err := l.svcCtx.Rdb.Get(l.ctx, req.Phone).Result()
	//if err != nil {
	//	err = errors.New("验证码错误或不存在")
	//	return nil, err
	//}
	//if code != realCode {
	//	err = errors.New("验证码错误")
	//	return nil, err
	//}

	user := new(user_model.SysUser)
	user.Username = req.Username
	user.Password = utils.BcryptHash(req.Password)
	user.Phone = req.Phone
	user.Uuid = uuid.NewV4().String()
	user.Email = req.Email
	user.Enable = req.Enabled
	user.AuthorityId = req.AuthorityId
	user.AuthorityIds = req.AuthorityIds
	users, err := l.svcCtx.UserModel.Find(l.ctx, user)

	if len(users) > 0 {
		err = errors.New("用户名已存在,请直接登录")
		return nil, err
	}
	//插入用户表用户数据
	InsertResut, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	userId, err := InsertResut.LastInsertId()
	if err != nil {
		err = errors.New("获取插入Id 失败，")
		return nil, err
	}
	//角色表关联角色
	var user_auths = make([]*user_auth.SysUserAuthority, 0)
	for _, item := range req.AuthorityIds {
		sysUser := user_auth.SysUserAuthority{
			SysUserId:               userId,
			SysAuthorityAuthorityId: item,
		}
		user_auths = append(user_auths, &sysUser)
	}
	//调用插入接口
	err = l.svcCtx.AuthUser.Inserts(user_auths)
	if err != nil {
		return nil, err
	}

	resp = new(types.RegisterResp)
	resp.Message = "注册成功"

	return resp, nil
}
