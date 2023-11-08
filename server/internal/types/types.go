// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequset struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message     string `json:"message"`
	Token       string `json:"token"`
	Id          int    `json:"id"`
	AuthorityId int    `json:"authorityId"`
}

type EditUserRequest struct {
	Id           int64   `json:"id"`
	Username     string  `json:"username"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email"`
	Enable       int64   `json:"enable"`
	AuthorityIds []int64 `json:"authorityIds"`
}

type EditUserResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type RegisterReq struct {
	Username     string  `json:"username"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email,optional"`
	Password     string  `json:"password"`
	Enabled      int64   `json:"enable"`
	AuthorityIds []int64 `json:"authorityIds"`
}

type RegisterResp struct {
	Message string `json:"message"`
}

type SendIdentityMesReq struct {
	Phone string `json:"phone"`
}

type SendIdentityMesResp struct {
	Message string `json:"message"`
}

type ChangePasswordReq struct {
	UserId         int64  `json:"userId"`
	PriPassword    string `json:"priPassword"`
	ChangePassword string `json:"changePassword"`
}

type ChangePasswordResp struct {
	Message string `json:"message"`
}

type UserListReq struct {
	PageNum  int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type UserListResp struct {
	Page     int64          `json:"page"`
	PageSize int64          `json:"pageSize"`
	UserList []*UserMessage `json:"userList"`
	Total    int64          `json:"total"`
}

type UserMessage struct {
	Id          int64        `json:"id"`
	Username    string       `json:"username"`
	Phone       string       `json:"phone"`
	Email       string       `json:"email"`
	CreateTime  int64        `json:"createTime"`
	UpdateTime  int64        `json:"updateTime"`
	AuthorityId int64        `json:"authorityId"`
	Enable      int          `json:"enable"`
	Uuid        string       `json:"uuid"`
	Authorities []*Authority `json:"authorities"`
}

type Authority struct {
	CreateAt      int64  `json:"createAt"`
	UpdateAt      int64  `json:"updateAt"`
	AuthorityId   int64  `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
	ParentId      int64  `json:"parentId"`
}

type SetUserAuthorityReq struct {
	AuthorityId []int64 `json:"authorityId"`
	UserId      int64   `json:"userId"`
}

type SetUserAuthorityResp struct {
	Message string `json:"message"`
}

type DeleteUserReq struct {
	Id int64 `json:"id"`
}

type DeleteUserResp struct {
	Message string `json:"message"`
}

type MenuRequest struct {
	Id int `json:"id"`
}

type MenuResponse struct {
	Resp string `json:"resp"`
}

type Menu struct {
	Id        int     `json:"id"`
	CreatedAt int64   `json:"createdAt,optional"`
	UpdatedAt int64   `json:"updatedAt,optional"`
	ParentId  string  `json:"parentId,optional"`
	Path      string  `json:"path,optional"`
	Name      string  `json:"name,optional"`
	Hidden    bool    `json:"hidden,optional"`
	Sort      int     `json:"sort,optional"`
	KeepAlive bool    `json:"keepAlive,optional"`
	Title     string  `json:"title,optional"`
	Icon      string  `json:"icon,optional"`
	CloseTab  int     `json:"closeTab,optional"`
	Children  []*Menu `json:"children,optional"`
}

type GetMenuListReq struct {
}

type GetMenuListResp struct {
	Total    int     `json:"total"`
	MenuList []*Menu `json:"menuList"`
}

type AddMenuReq struct {
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Sort      int    `json:"sort"`
	KeepAlive bool   `json:"keepAlive"`
	Title     string `json:"title"`
	Icon      string `json:"icon ,optional"`
	CloseTab  int    `json:"closeTab, optional"`
}

type AddMenuResp struct {
	Message string `json:"message"`
}

type DeleteMenuReq struct {
	Id int64 `json:"id"`
}

type DeleteMenuResp struct {
	Message string `json:"message"`
}

type EditMenuReq struct {
	Id        int64  `json:"id"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Sort      int    `json:"sort"`
	KeepAlive bool   `json:"keepAlive"`
	Title     string `json:"title"`
	Icon      string `json:"icon, optional"`
	CloseTab  int    `json:"closeTab ,optional"`
}

type EditMenuResp struct {
	Message string `json:"message"`
}

type GetMenuAuthorityReq struct {
	AuthorityId int64 `json:"authorityId"`
}

type GetMenuAuthorityResp struct {
	MenuList []*Menu `json:"menuList"`
}

type AddMenuAuthorityReq struct {
	MenuList    []*Menu `json:"menuList"`
	AuthorityId int64   `json:"authorityId"`
}

type AddMenuAuthorityResp struct {
	Message string `json:"message"`
}

type CreateAuthorityReq struct {
	AuthorityId   int64  `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
	ParentId      int64  `json:"parentId"`
}

type CreateAuthorityResp struct {
	Message string `json:"message"`
}

type GetAuthorityListReq struct {
	PageNum  int64 `json:"pageNum,optional"`
	PageSize int64 `json:"pageSize,optional"`
}

type GetAuthorityList struct {
	CreateAt      int64               `json:"createAt"`
	UpdateAt      int64               `json:"updateAt"`
	AuthorityId   int64               `json:"authorityId"`
	AuthorityName string              `json:"authorityName"`
	ParentId      int64               `json:"parentId"`
	Children      []*GetAuthorityList `json:"children"`
}

type GetAuthorityListResp struct {
	AuthorityList []*GetAuthorityList `json:"authorityList"`
}

type DeleteAuthorityReq struct {
	AuthorityId int64 `json:"authorityId "`
}

type DeleteAuthorityResp struct {
	Message string `json:"message"`
}

type UpdateAuthorityReq struct {
	AuthorityId   int64  `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
	ParentId      int64  `json:"parentId"`
}

type UpdateAuthorityResp struct {
	Message string `json:"message"`
}

type SetDataAuthorityReq struct {
}

type SetDataAuthorityResp struct {
}

type TestReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TestResp struct {
	Message string `json:"message"`
}
