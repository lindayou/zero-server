package user

import (
	"context"
	"fmt"
	"time"

	"zero-server/server/internal/svc"
	"zero-server/server/internal/types"
	"zero-server/server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendIdentityMesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendIdentityMesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendIdentityMesLogic {
	return &SendIdentityMesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendIdentityMesLogic) SendIdentityMes(req *types.SendIdentityMesReq) (resp *types.SendIdentityMesResp, err error) {
	// todo: add your logic here and delete this line
	phone := req.Phone
	fmt.Println("get phone is  ", phone)
	code := utils.GenerateSmsCode(4)
	fmt.Println("this is code", code)
	// todo:存储到redis 并设置过期时间  手机号：验证码的形式

	err = l.svcCtx.Rdb.Set(l.ctx, phone, code, 60*time.Second).Err()
	if err != nil {
		return nil, err
	}
	//调用阿里云服务去发送验证码
	//没有错误则返回成功
	resp = new(types.SendIdentityMesResp)
	resp.Message = "发送成功"
	return resp, nil
}
