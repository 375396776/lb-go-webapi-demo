package router

import (
	"github.com/gin-gonic/gin"
	v1 "invoice/api/v1"
)

func InitInvoiceRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/invoices")
	routerGroup.POST("/register", v1.RegisterEnterprise)
	routerGroup.POST("/register2", v1.Register)
	routerGroup.GET("/test", v1.Test)
	routerGroup.GET("/ent/:id", v1.GetEnterpriseById)
	routerGroup.PATCH("/ent/:id", v1.UpdateEnterpriseById)
	routerGroup.DELETE("/ent/:id", v1.DelById)
}
