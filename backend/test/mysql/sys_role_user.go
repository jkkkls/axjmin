package main

import (
	"time"
)

type SysUser struct {
	Id          int64 `gorm:"primarykey"`
	Username    string
	Password    string
	Name        string
	Type        int
	CreateTs    *time.Time
	LastLoginTs *time.Time
	Status      int

	SysRoleUsers []*SysRoleUser
}

type SysRole struct {
	Id       int64 `gorm:"primarykey"`
	Name     string
	Desc     string
	CreateTs *time.Time
	UpdateTs *time.Time
	CreateBy string
}

type SysRoleUser struct {
	Id        int64 `gorm:"primarykey"`
	SysUserID int64
	SysRoleID int64
	SysRole   *SysRole
}

func init() {
	tables = append(tables, &SysRole{}, &SysUser{}, &SysRoleUser{})
}
