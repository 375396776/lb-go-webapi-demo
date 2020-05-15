package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Success bool
	Code    int
	Msg     string
	Data    interface{}
}

func response(success bool, code int, msg string, data interface{}, ctx *gin.Context) {
	r := Result{success, code, msg, data}
	ctx.JSON(http.StatusOK, r)
}

func successResponse(data interface{}, ctx *gin.Context) {
	response(true, 0, "请求成功", data, ctx)
}

func failResponse(code int, msg string, ctx *gin.Context) {
	response(false, code, msg, nil, ctx)
}

func SuccessResultWithEmptyData(ctx *gin.Context) {
	successResponse(nil, ctx)
}

func SuccessResult(data interface{}, ctx *gin.Context) {
	successResponse(data, ctx)
}

func FailResultWithDefaultMsg(code int, ctx *gin.Context) {
	failResponse(code, "请求失败", ctx)
}

func FailResult(code int, msg string, ctx *gin.Context) {
	failResponse(code, msg, ctx)
}
