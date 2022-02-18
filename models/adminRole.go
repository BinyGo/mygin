/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-23 18:16:20
 * @LastEditTime: 2022-01-23 18:32:53
 */
package models

type AdminRole struct {
	Id      int64 `json:"id" form:"id"`
	AdminId int64 `json:"admin_id" form:"admin_id"`
	RoleId  int64 `json:"role_id" form:"role_id"`
}

type Role struct {
	Id    int64  `json:"id" form:"id"`
	Group int    `json:"group" form:"group"`
	Title string `json:"title" form:"title"`
	Auth  string `json:"auth" form:"auth"`
}

func (p *AdminRole) TableName() string {
	return "biny_admin_role"
}

func (p *Role) TableName() string {
	return "biny_role"
}

func GetAdminRoles(id int64) []*Role {
	var roles []*Role
	db.Raw("SELECT r.id,r.group,r.title,r.auth FROM biny_admin_role ar LEFT JOIN `biny_role` r on ar.`role_id` = r.id where ar.admin_id = ? ", id).Scan(&roles)
	return roles

}
