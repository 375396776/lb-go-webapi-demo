package application

import (
	"invoice/application/initial"
)

func Run() {
	r := initial.Router()
	//TODO 这里可以创建其他的服务
	r.Run(":8089")
}
