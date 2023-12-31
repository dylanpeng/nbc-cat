// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import "fmt"

const TableNameDataAdminUser = "data_admin_user"

// DataAdminUser mapped from table <data_admin_user>
type DataAdminUser struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键id" json:"id"` // 主键id
	Name       string `gorm:"column:name;not null;comment:登录名" json:"name"`                   // 登录名
	Password   string `gorm:"column:password;not null;comment:密码" json:"password"`            // 密码
	Roles      Int64s `gorm:"column:roles;not null;comment:角色" json:"roles"`                  // 角色
	Status     int32  `gorm:"column:status;not null;comment:状态(1.正常 2.删除)" json:"status"`     // 状态(1.正常 2.删除)
	CreateTime int64  `gorm:"column:create_time;not null;comment:创建时间" json:"create_time"`    // 创建时间
	UpdateTime int64  `gorm:"column:update_time;not null;comment:更新时间" json:"update_time"`    // 更新时间
}

// TableName DataAdminUser's table name
func (*DataAdminUser) TableName() string {
	return TableNameDataAdminUser
}

func (e *DataAdminUser) PrimaryPairs() []interface{} {
	return []interface{}{"id", e.ID}
}

func (e *DataAdminUser) PrimarySeted() bool {
	return e.ID > 0
}

func (e *DataAdminUser) String() string {
	return fmt.Sprintf("%+v", *e)
}
