package middleware

import (
	"cat/common"
	"cat/common/consts"
	ctrl "cat/common/control"
	"cat/common/exception"
	"cat/internal/admin/config"
	"cat/internal/admin/logic/service"
	"cat/internal/admin/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func AdminAuth(ctx *gin.Context) {
	var accessToken string
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	getToken := ctx.DefaultQuery("token", "")

	if authorizationHeader == "" && getToken == "" {
		common.Logger.Infof("AdminAuth token is empty.")
		ctrl.Error(ctx, exception.CodeUnauthorized)
		ctx.Abort()
		return
	}

	header := strings.Split(authorizationHeader, " ")
	if len(header) == 2 && header[0] == "Bearer" {
		accessToken = header[1]
	} else {
		accessToken = getToken
	}

	//解析token
	token, err := jwt.ParseWithClaims(accessToken, &types.AdminClaims{}, GetJwtKey)
	if err != nil {
		common.Logger.Infof("AdminAuth token parse failed. | err: %s", err)
		ctrl.Error(ctx, exception.CodeUnauthorized, err)
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(*types.AdminClaims)
	if !ok || !token.Valid {
		common.Logger.Infof("AdminAuth token invalid. | err: %s", err)
		ctrl.Error(ctx, exception.CodeUnauthorized, err)
		ctx.Abort()
		return

	}

	ctx.Set(consts.CtxValueAdminAuth, claims)
	ctx.Next()
}

func GetJwtKey(token *jwt.Token) (i interface{}, err error) {
	return []byte(config.GetConfig().App.Secret), nil
}

func PermissionCheck(ctx *gin.Context) {
	v, ok := ctx.Get(consts.CtxValueAdminAuth)

	if !ok || v == nil {
		ctrl.Error(ctx, exception.CodeNoPermission)
		ctx.Abort()
		return
	}

	claim, ok := v.(*types.AdminClaims)

	if !ok || claim == nil {
		ctrl.Error(ctx, exception.CodeNoPermission)
		ctx.Abort()
		return
	}

	p := ctx.Request.URL.Path
	uri := ctx.Request.URL.RequestURI()
	fmt.Printf("path: %s\n uri:%s\n", p, uri)

	adminUser, err := service.AdminUser.GetAdminUserById(claim.UserId)
	if err != nil || len(adminUser.AdminPermissionPaths) == 0 {
		ctrl.Error(ctx, exception.CodeNoPermission)
		ctx.Abort()
		return
	}

	var hasPermission bool
	for _, path := range adminUser.AdminPermissionPaths {
		if path.Path == ctx.Request.URL.Path {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		ctrl.Error(ctx, exception.CodeNoPermission)
		ctx.Abort()
		return
	}

	ctx.Next()
	return
}
