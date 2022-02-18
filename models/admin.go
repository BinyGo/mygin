/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 21:38:34
 * @LastEditTime: 2022-01-23 23:06:41
 */
package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

//var db *gorm.DB

type Admin struct {
	Id       int64  `json:"id" form:"id" binding:"required,gt=0,lt=999999999" zh:"id" `        //大于0小于999999999
	Username string `json:"username" form:"username" binding:"required,min=6,max=20" zh:"用户名"` //大于6位,小于20位 gt=6,lt=20
	Password string `json:"password" form:"password" binding:"required,min=6,max=20" zh:"密码"`  //gorm:"index:,sort:desc,collate:utf8,type:btree,length:10,where:name3 != 'asd'"
	//RePassword string `form:"repassword" binding:"required,gt=5,lt=20" zh:"确认密码"`
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile" zh:"手机号码"`
	Email  string `json:"email" form:"email" binding:"required,email" zh:"邮箱地址"`
	//Group      int    `form:"group"`
	CreateTime MyTime `json:"create_time" form:"create_time,omitempty" gorm:"autoCreateTime;type:timestamp"` //gorm:"autoCreateTime" default:null
	//CreateTime int64 `json:"create_time" form:"create_time,omitempty"`  //default:null
	UpdateTime MyTime `json:"update_time" form:"update_time,omitempty" gorm:"default:null;type:timestamp"` //validate:"omitempty"
}

//如果没指定表名,GORM 使用结构体名的蛇形复数作为表名。例如：结构体名为 DockerInstance ，则表名为 dockerInstances
func (p *Admin) TableName() string {
	return "biny_admin2"
}

type AdminCount struct {
	Id    int64 `json:"id"`
	Count int64 `json:"count"`
}

func (p *AdminCount) TableName() string {
	return "biny_admin_count"
}

func GetAdminById(id int64) *Admin {
	admin := &Admin{}
	db.First(admin, "id = ?", id)
	return admin
}

func GetAdminByUsername(username string) *Admin {
	admin := &Admin{}
	db.First(admin, "username = ?", username)
	return admin
}

func GetAdminList(limit int, offset int) []*Admin {
	var admins []*Admin
	db.Limit(limit).Offset(offset).Order("id desc").Find(&admins)
	//fmt.Println(admins)
	return admins
}

func AdminCreate(admin *Admin) error {
	adminCount := &AdminCount{}
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		result := tx.Create(admin)
		row := result.RowsAffected
		fmt.Println(result.RowsAffected)
		if row == 0 {
			return errors.New("添加失败")
		}

		if err := tx.Model(adminCount).Where("id = ?", 1).Update("count", gorm.Expr("count + ? ", 1)).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	return err
}

func AdminUpdate(admin *Admin) bool {
	result := db.Model(admin).Updates(admin)
	return result.RowsAffected > 0
}

func AdminDelete(id int64) error {
	admin := &Admin{}
	adminCount := &AdminCount{}
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）  //RowsAffected
		result := tx.Model(admin).Where("id = ?", id).Delete(admin)
		row := result.RowsAffected //fmt.Println(result.RowsAffected)
		if row == 0 {
			return errors.New("删除失败")
		}

		if err := tx.Model(adminCount).Where("id = ?", 1).Update("count", gorm.Expr("count - ? ", 1)).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	return err
}

func GetAdminCount() int64 {
	adminCount := &AdminCount{}
	db.First(adminCount, "id = ?", 1)
	return adminCount.Count
}

// func (admin *Admin) AfterFind(db *gorm.DB) (err error) {

// 	fmt.Println("-----AfterFind-------")

// 	for _, field := range db.Statement.Schema.Fields {
// 		fmt.Println(field)
// 		return nil
// 		// err := field.Set(db.Statement.ReflectValue, 11)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 	}

// 	return nil
// }

// AfterCreate run after create database record. create后的触手怪
// func (admin *Admin) AfterCreate(tx *gorm.DB) error {
// 	admin.Username = "test--" + admin.Username //新建数据后,用户名前面加"test--"
// 	return tx.Save(admin).Error
// }

// func AdminDelete(id int64) bool {
// 	admin := &Admin{}
// 	result := db.Where("id = ?", id).Delete(admin)
// 	return result.RowsAffected > 0
// }

// func (u *dbs) Get(ctx context.Context, username string) (*Admin, error) {
// 	admin := &Admin{}
// 	err := u.db.Where("name = ? ", username).First(&admin).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.WithCode(500, err.Error())
// 		}

// 		return nil, errors.WithCode(500, err.Error())
// 	}

// 	return admin, nil
// }
