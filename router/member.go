/*
 description:

 @author lib
 @since 2020/05/13
*/
package router

import (
	"github.com/gin-gonic/gin"
	v1 "invoice/api/v1"
)

func InitMemberRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/members")
	routerGroup.POST("/", v1.AddMember)
}
