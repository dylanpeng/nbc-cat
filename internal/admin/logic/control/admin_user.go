package control

import (
	"cat/common/consts"
	ctrl "cat/common/control"
	"cat/internal/admin/logic/service"
	"cat/internal/admin/types"
	"github.com/gin-gonic/gin"
)

var AdminUser = &adminUserCtrl{}

type adminUserCtrl struct{}

func (c *adminUserCtrl) Login(ctx *gin.Context) {
	req := &types.LoginReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, len(req.Name) == 0 || len(req.Password) == 0) {
		return
	}

	rsp, err := service.AdminUser.Login(req)
	if err != nil {
		ctrl.Exception(ctx, err)
		return
	}

	resp := &ctrl.Response{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data:    rsp,
	}

	ctrl.SendRsp(ctx, resp)
}

func (c *adminUserCtrl) Add(ctx *gin.Context) {
	req := &types.AdminUserAddReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, len(req.Name) == 0 || len(req.Password) == 0) {
		return
	}

	err := service.AdminUser.AddAdminUser(req)
	if err != nil {
		ctrl.Exception(ctx, err)
		return
	}

	resp := &ctrl.Response{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
	}

	ctrl.SendRsp(ctx, resp)
}
