package service

import (
	"cat/common"
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

	e = common.ConvertStruct(adminUser, result.AdminUser)
	if e != nil {
		err = exception.New(exception.CodeConvertFailed)
		return
	}

	claims := &types.AdminClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "nbc-cat-admin-token",
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 7)),
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
		Name:  req.Name,
		Roles: req.Roles,
	}

	//生成hash
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	//转换成字符串
	adminUser.Password = string(hashBytes)

	e = model.AdminUser.Add(adminUser)
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		common.Logger.Infof("adminUserSrv AddAdminUser add fail. | err: %s | req: %+v", e, *req)
		return
	}

	return
}
