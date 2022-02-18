/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-23 18:29:01
 * @LastEditTime: 2022-01-23 18:52:19
 */
package models

import "fmt"

type AdminMenu struct {
	Id         int64        `json:"id" form:"id"`
	ParentId   int64        `json:"parent_id" form:"parent_id"`
	Type       int          `json:"type" form:"type"`
	Status     int          `json:"status" form:"status"`
	Sort       int64        `json:"sort" form:"sort"`
	Controller string       `json:"controller" form:"controller"`
	Action     string       `json:"action" form:"action"`
	Param      string       `json:"param" form:"param"`
	Path       string       `json:"path" form:"path"`
	Title      string       `json:"title" form:"title"`
	Icon       string       `json:"icon" form:"icon"`
	IsMenu     int          `json:"is_menu" form:"is_menu"`
	Level      int          `json:"level" form:"level"`
	Children   []*AdminMenu `json:"children" form:"children" orm:"-" gorm:"-"`
}

//1. Id 字段不暴露给用户，则使用 `json:"-"` 修饰。
//2. Inputs、Outputs 在某些情况下不返回字段数据。(1)、使用 `json:"omitempty"`（当字段为空时忽略此字段） 修饰字段；(2)、当不需要该字段返回时，让其赋值为空即可。

func (p *AdminMenu) TableName() string {
	return "biny_admin_menu"
}

func GetAllMenu() []*AdminMenu {
	var menus []*AdminMenu
	db.Where("is_menu = ?", 1).Find(&menus)
	menus = makeTree(menus, 0, 0)
	return menus
}

// * // 递归实现无限分类
func makeTree(menus []*AdminMenu, pid int64, level int) []*AdminMenu {
	fmt.Println(&menus)
	var tree []*AdminMenu
	for i := 0; i < len(menus); {
		row := menus[i]
		if row.ParentId == pid {
			row.Level = level
			menus = append(menus[:i], menus[i+1:]...)
			fmt.Println(menus)
			children := makeTree(menus, row.Id, level+1)
			if children != nil {
				row.Children = children
			}
			tree = append(tree, row)
		} else {
			i++
		}
	}
	return tree
}
