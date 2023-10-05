package types

type AdminUserAddReq struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Roles    []int64 `json:"roles"`
}

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRsp struct {
	AdminUser *AdminUser
	Token     string
}

type AdminUser struct {
	Name   string  `json:"name"`
	Roles  []int64 `json:"roles"`
	Status int     `json:"status"`
}
