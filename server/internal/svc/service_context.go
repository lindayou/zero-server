package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"zero-server/server/internal/config"
	"zero-server/server/internal/middleware"
	"zero-server/server/model/admin_dic"
	"zero-server/server/model/admin_dic_detail"
	"zero-server/server/model/admin_operation"
	"zero-server/server/model/anthority_model"
	"zero-server/server/model/menu"
	"zero-server/server/model/user_auth"
	"zero-server/server/model/user_model"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       user_model.SysUserModel
	MenuModel       menu.SysBaseMenusModel
	Authority       anthority_model.SysAuthoritiesModel
	AuthUser        user_auth.SysUserAuthorityModel
	DicModel        admin_dic.SysDictionariesModel
	DicDetail       admin_dic_detail.SysDictionaryDetailsModel
	Operation       admin_operation.SysOperationRecordsModel
	Rdb             *redis.Client
	OperationRecord rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: user_model.NewSysUserModel(conn),
		MenuModel: menu.NewSysBaseMenusModel(conn),
		Authority: anthority_model.NewSysAuthoritiesModel(conn),
		AuthUser:  user_auth.NewSysUserAuthorityModel(conn),
		DicModel:  admin_dic.NewSysDictionariesModel(conn),
		DicDetail: admin_dic_detail.NewSysDictionaryDetailsModel(conn),
		Operation: admin_operation.NewSysOperationRecordsModel(conn),
		Rdb: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host + ":" + c.Redis.Port,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
		OperationRecord: middleware.NewOperationRecordMiddleware().Handle,
	}
}
