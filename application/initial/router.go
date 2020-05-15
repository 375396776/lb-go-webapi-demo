package initial

import (
	"github.com/gin-gonic/gin"
	"invoice/router"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/")
	router.InitInvoiceRouter(api)
	router.InitMemberRouter(api)
	return r
}
