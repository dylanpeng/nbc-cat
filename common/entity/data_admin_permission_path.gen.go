// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import "fmt"

const TableNameDataAdminPermissionPath = "data_admin_permission_path"

// DataAdminPermissionPath mapped from table <data_admin_permission_path>
type DataAdminPermissionPath struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键id" json:"id"` // 主键id
	PermSymbol string `gorm:"column:perm_symbol;not null;comment:权限标识" json:"perm_symbol"`    // 权限标识
	Path       string `gorm:"column:path;not null;comment:接口路径" json:"path"`                  // 接口路径
	CreateTime int64  `gorm:"column:create_time;not null;comment:创建时间戳" json:"create_time"`   // 创建时间戳
	UpdateTime int64  `gorm:"column:update_time;not null;comment:更新时间戳" json:"update_time"`   // 更新时间戳
}

// TableName DataAdminPermissionPath's table name
func (*DataAdminPermissionPath) TableName() string {
	return TableNameDataAdminPermissionPath
}

func (e *DataAdminPermissionPath) PrimaryPairs() []interface{} {
	return []interface{}{"id", e.ID}
}

func (e *DataAdminPermissionPath) PrimarySeted() bool {
	return e.ID > 0
}

func (e *DataAdminPermissionPath) String() string {
	return fmt.Sprintf("%+v", *e)
}
