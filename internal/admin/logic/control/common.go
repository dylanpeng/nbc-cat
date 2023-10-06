package control

import (
	"cat/common/consts"
	"cat/common/exception"
	"cat/internal/admin/types"
	"github.com/gin-gonic/gin"
)

func GetCurrUser(ctx *gin.Context) (user *types.AdminClaims, err *exception.Exception) {
	if v, ok := ctx.Get(consts.CtxValueAdminAuth); ok && v != nil {
		user = v.(*types.AdminClaims)
	}

	if user == nil || user.UserId <= 0 {
		err = exception.New(exception.CodeUnauthorized)
		return
	}

	return
}
