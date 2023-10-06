package service

import (
	"cat/common"
	"cat/common/consts"
	"cat/common/entity"
	"cat/common/exception"
	"cat/common/model"
	"cat/internal/admin/config"
	"cat/internal/admin/types"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var AdminUser = &adminUserSrv{}

type adminUserSrv struct{}

func (s *adminUserSrv) Login(req *types.LoginReq) (result *types.LoginRsp, err *exception.Exception) {
	adminUser, e := model.AdminUser.GetByName(req.Name)
	if e != nil {
		err = exception.New(exception.CodeQueryFailed, e, req)
		return
	}

	if adminUser == nil {
		err = exception.New(exception.CodeDataNotExist)
		return
	}

	e = bcrypt.CompareHashAndPassword([]byte(adminUser.Password), []byte(req.Password))
	if e != nil {
		err = exception.New(exception.CodeUserPasswordErr)
		return
	}

	result = &types.LoginRsp{
		AdminUser: &types.AdminUser{},
	}

	admin, err := s.GetAdminUserById(adminUser.ID)
	if err != nil {
		return
	}

	result.AdminUser = admin

	claims := &types.AdminClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "nbc-cat-admin-token",
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 1, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
		UserId: adminUser.ID,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result.Token, e = jwtToken.SignedString([]byte(config.GetConfig().App.Secret))
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		return
	}

	return
}

func (s *adminUserSrv) AddAdminUser(req *types.AdminUserAddReq) (err *exception.Exception) {
	existUser, e := model.AdminUser.GetByName(req.Name)
	if e != nil {
		err = exception.New(exception.CodeQueryFailed, e, req)
		return
	}

	if existUser != nil {
		err = exception.New(exception.CodeDataAlreadyExist)
		return
	}

	adminUser := &entity.DataAdminUser{
		Name:   req.Name,
		Roles:  req.Roles,
		Status: consts.AdminStatusNormal,
	}

	//生成hash
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	//转换成字符串
	adminUser.Password = string(hashBytes)

	e = model.AdminUser.Add(adminUser)
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		common.Logger.Infof("adminUserSrv AddAdminUser add fail. | err: %s | req: %s", e, req)
		return
	}

	return
}

func (s *adminUserSrv) GetAdminUserById(id int64) (result *types.AdminUser, err *exception.Exception) {
	result = &types.AdminUser{}

	adminUser, err := s.GetAdminUserEntityById(id)
	if err != nil {
		return
	}

	e := common.ConvertStruct(adminUser, result)
	if e != nil {
		err = exception.New(exception.CodeConvertFailed, e)
		common.Logger.Infof("adminUserSrv GetAdminUserById convert fail. | err: %s | id: %d", e, id)
		return
	}

	roles, err := AdminRole.GetAdminRolesByIds(adminUser.Roles)
	if err != nil {
		return
	}

	if len(roles) == 0 {
		return
	}

	result.AdminRoles = roles

	permissionIds := make([]int64, 0)
	for _, role := range roles {
		if len(role.Perms) > 0 {
			permissionIds = append(permissionIds, role.Perms...)
		}
	}

	permissions, err := AdminRole.GetAdminPermissionByIds(permissionIds)
	if err != nil {
		return
	}

	result.AdminPermissions = permissions

	if len(permissions) == 0 {
		return
	}

	symbols := make([]string, 0)
	for _, permission := range permissions {
		symbols = append(symbols, permission.PermSymbol)
	}

	paths, err := AdminRole.GetAdminPermissionPathBySymbols(symbols)
	if err != nil {
		return
	}

	result.AdminPermissionPaths = paths

	return
}

func (s *adminUserSrv) GetAdminUserEntityById(id int64) (result *entity.DataAdminUser, err *exception.Exception) {
	result = &entity.DataAdminUser{ID: id}
	exist, e := model.AdminUser.Get(result)
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		common.Logger.Infof("adminUserSrv GetAdminUserById fail. | err: %s | id: %d", e, id)
		return
	}

	if !exist {
		err = exception.New(exception.CodeDataNotExist)
		common.Logger.Infof("adminUserSrv GetAdminUserById data not exists. | id: %d", e, id)
		return
	}

	return
}
