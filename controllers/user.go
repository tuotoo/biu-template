package controllers

import (
	"github.com/emicklei/go-restful"
	"github.com/tuotoo/biu"
	"github.com/tuotoo/biu-template/models"
)

// UserResource is the REST layer to the User domain
type UserResource struct{}

// WebService creates a new service that can handle REST requests for User resources.
func (u UserResource) WebService(ws biu.WS) {
	ws.Route(ws.GET("/").To(u.findAllUsers).
		Doc("get all Users").
		Writes([]models.User{}).
		Returns(200, "OK", []models.User{}), nil)

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Writes(models.User{}).
		Returns(200, "OK", models.User{}).
		Returns(404, "Not Found", nil), nil)

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Reads(models.User{}), nil)

	ws.Route(ws.PUT("/").To(u.createUser).
		Doc("create a user").
		Reads(models.User{}), nil)

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		Doc("delete a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")), nil)
}

func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	biu.Debug("UserResource.findAllUsers", biu.Log())
	biu.ResponseJSON(response, models.User{})
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	biu.Debug("UserResource.findUser", biu.Log().Str("id", id))
}

func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(models.User)
	err := request.ReadEntity(&usr)
	biu.Debug("UserResource.updateUser", biu.Log().Interface("user", usr).Err(err))
}

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := models.User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	biu.Debug("UserResource.createUser", biu.Log().Interface("user", usr).Err(err))
}

func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	biu.Debug("UserResource.removeUser", biu.Log().Str("id", id))
}
