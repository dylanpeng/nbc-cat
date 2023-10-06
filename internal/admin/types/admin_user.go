package types

import "fmt"

type AdminUserAddReq struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Roles    []int64 `json:"roles"`
}

func (e *AdminUserAddReq) String() string {
	return fmt.Sprintf("%+v", *e)
}

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (e *LoginReq) String() string {
	return fmt.Sprintf("%+v", *e)
}

type LoginRsp struct {
	AdminUser *AdminUser `json:"admin_user"`
	Token     string     `json:"token"`
}

func (e *LoginRsp) String() string {
	return fmt.Sprintf("%+v", *e)
}

type AdminUser struct {
	ID                   int64                  `json:"id"`
	Name                 string                 `json:"name"`
	Roles                []int64                `json:"roles"`
	Status               int                    `json:"status"`
	AdminRoles           []*AdminRole           `json:"admin_roles"`
	AdminPermissions     []*AdminPermission     `json:"admin_permissions"`
	AdminPermissionPaths []*AdminPermissionPath `json:"admin_permission_paths"`
}

func (e *AdminUser) String() string {
	return fmt.Sprintf("%+v", *e)
}

type AdminRole struct {
	ID               int64              `json:"id"`
	Name             string             `json:"name"`
	Perms            []int64            `json:"perms"`
	AdminPermissions []*AdminPermission `json:"admin_permissions"`
}

func (e *AdminRole) String() string {
	return fmt.Sprintf("%+v", *e)
}

type AdminPermission struct {
	ID                   int64                  `json:"id"`
	Name                 string                 `json:"name"`
	PermSymbol           string                 `json:"perm_symbol"`
	Sort                 int                    `json:"sort"`
	ParentId             int64                  `json:"parent_id"`
	AdminPermissionPaths []*AdminPermissionPath `json:"admin_permission_paths"`
}

func (e *AdminPermission) String() string {
	return fmt.Sprintf("%+v", *e)
}

type AdminPermissionPath struct {
	ID         int64  `json:"id"`
	PermSymbol string `json:"perm_symbol"`
	Path       string `json:"path"`
}

func (e *AdminPermissionPath) String() string {
	return fmt.Sprintf("%+v", *e)
}
