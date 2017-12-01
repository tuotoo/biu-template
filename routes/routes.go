package routes

import (
	"github.com/tuotoo/biu"
	"github.com/tuotoo/biu-template/controllers"
)

func init() {
	biu.AddServices("/v1", nil,
		biu.NS{
			NameSpace:  "user",
			Controller: controllers.UserResource{},
			Desc:       "用户管理",
		},
	)
}
