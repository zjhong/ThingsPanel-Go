// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysVersion = "sys_version"

// SysVersion mapped from table <sys_version>
type SysVersion struct {
	VersionNumber int32  `gorm:"column:version_number;primaryKey" json:"version_number"`
	Version       string `gorm:"column:version;not null" json:"version"`
}

// TableName SysVersion's table name
func (*SysVersion) TableName() string {
	return TableNameSysVersion
}