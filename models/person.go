/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 21:04:04
 * @LastEditTime: 2022-01-22 21:32:50
 */
package models

type Person struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
	// Address string `form:"address"`
	// Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	// CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	// UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}
