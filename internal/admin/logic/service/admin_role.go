package service

import (
	"cat/common"
	"cat/common/exception"
	"cat/common/model"
	"cat/internal/admin/types"
)

var AdminRole = &adminRoleSrv{}

type adminRoleSrv struct{}

func (s *adminRoleSrv) GetAdminRolesByIds(ids []int64) (result []*types.AdminRole, err *exception.Exception) {
	roles, e := model.AdminRole.GetByIds(ids)
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		common.Logger.Infof("adminRoleSrv GetAdminRolesByIds GetByIds fail. | err: %s | id: %+v", e, ids)
		return
	}

	result = make([]*types.AdminRole, len(roles))
	if len(roles) > 0 {
		e = common.ConvertStruct(roles, &result)
		if e != nil {
			err = exception.New(exception.CodeConvertFailed, e)
			common.Logger.Infof("adminRoleSrv GetAdminRolesByIds convert fail. | err: %s | id: %+v", e, ids)
			return
		}
	}

	return
}

func (s *adminRoleSrv) GetAdminPermissionByIds(ids []int64) (result []*types.AdminPermission, err *exception.Exception) {
	permissions, e := model.AdminPermission.GetByIds(ids)
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		common.Logger.Infof("adminRoleSrv GetAdminPermissionByIds GetByIds fail. | err: %s | id: %+v", e, ids)
		return
	}

	result = make([]*types.AdminPermission, len(permissions))
	if len(permissions) > 0 {
		e = common.ConvertStruct(permissions, &result)
		if e != nil {
			err = exception.New(exception.CodeConvertFailed, e)
			common.Logger.Infof("adminRoleSrv GetAdminPermissionByIds convert fail. | err: %s | id: %+v", e, ids)
			return
		}
	}

	return
}

func (s *adminRoleSrv) GetAdminPermissionPathBySymbols(symbols []string) (result []*types.AdminPermissionPath, err *exception.Exception) {
	paths, e := model.AdminPermissionPath.GetPathBySymbol(symbols)
	if e != nil {
		err = exception.New(exception.CodeInternalError, e)
		common.Logger.Infof("adminRoleSrv GetAdminPermissionPathBySymbols GetByIds fail. | err: %s | symbols: %+v", e, symbols)
		return
	}

	result = make([]*types.AdminPermissionPath, len(paths))
	if len(paths) > 0 {
		e = common.ConvertStruct(paths, &result)
		if e != nil {
			err = exception.New(exception.CodeConvertFailed, e)
			common.Logger.Infof("adminRoleSrv GetAdminPermissionPathBySymbols convert fail. | err: %s | symbols: %+v", e, symbols)
			return
		}
	}

	return
}
