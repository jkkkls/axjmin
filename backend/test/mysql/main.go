package main

import (
	"log"
	"time"

	"github.com/gookit/goutil/dump"
)

func initUser() {
	now := time.Now()
	user := &SysUser{
		Username:    "guangbo",
		Name:        "guangbo",
		Password:    "123456",
		Type:        1,
		Status:      1,
		CreateTs:    &now,
		LastLoginTs: &now,
	}
	err := adminDB.Save(user)
	if err != nil {
		log.Println(err)
		return
	}

	err = adminDB.Save(&SysRoleUser{
		SysUserID: user.Id,
		SysRoleID: 1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	err = adminDB.Save(&SysRoleUser{
		SysUserID: user.Id,
		SysRoleID: 2,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	err := InitAdminDB("root:ghb.com123@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		return
	}

	//
	time.Sleep(time.Second)
	// initUser()
	user := &SysUser{}

	err = adminDB.DB.Where("username = ?", "guangbo").Preload("SysRoleUsers.SysRole").First(user).Error
	// err = adminDB.Query(user, "username = ?", "guangbo")
	log.Println(err)
	// log.Println(user.Id, user.Username, user.Name, user.Status, user.LastLoginTs, user.Type, user.CreateTs, len(user.SysRoleUsers))
	dump.Println(user)

	// user := &SysUser{Id: 1}
	// err = adminDB.DB.Unscoped().Select(clause.Associations).Delete(user).Error
	// // err = adminDB.Delete(&SysUser{}, "id = ?", 1)
	// // Select("Account").
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// err = adminDB.Save(&SysRoleUser{
	// 	SysUserID: 1,
	// 	SysRoleID: 1,
	// })
	// err = adminDB.Save(&SysRoleUser{
	// 	SysUserID: 1,
	// 	SysRoleID: 2,
	// })

	// err = adminDB.Save(&SysRole{
	// 	Name: "test",
	// 	Desc: "测试",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// now := time.Now()
	// err = adminDB.Save(&SysUser{
	// 	Username:    "guangbo",
	// 	Name:        "guangbo",
	// 	Password:    "123456",
	// 	Type:        1,
	// 	Status:      1,
	// 	CreateTs:    &now,
	// 	LastLoginTs: &now,
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = adminDB.Save(&SysUser{
	// 	Username: "admin",
	// 	Name:     "admin",
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	log.Println("----")
}
