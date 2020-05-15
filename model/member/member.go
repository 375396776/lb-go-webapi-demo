/*
 description:

 @author lib
 @since 2020/05/13
*/
package member

import "time"

// 定义 Member 结构体
type Member struct {
	Name       string    `form:"name" json:"name" binding:"required"`
	Age        int       `form:"age"  json:"age"  binding:"gt=10,lt=120"`
	Email      string    `form:"email" json:"email" binding:"email"`
	Address    string    `form:"address"`
	Birthday   string    `form:"birthday" json:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" json:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" json:"unixTime" time_format:"unix"`
	Mobile     string    `form:"mobile" json:"mobile" binding:"omitempty,phoneValid"`
	IdCard     string    `form:"idCard" json:"idCard" binding:"required,len=18"`
}
