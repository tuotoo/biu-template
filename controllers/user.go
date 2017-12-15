package controllers

import (
	"github.com/tuotoo/biu"
	"github.com/tuotoo/biu-template/models"
)

// UserResource is the REST layer to the User domain
type UserResource struct{}

// WebService creates a new service that can handle REST requests for User resources.
func (ctl UserResource) WebService(ws biu.WS) {
	ws.Route(ws.GET("/").To(biu.Handle(ctl.findAllUsers)).
		Doc("get all Users").
		Writes([]models.User{}).
		Returns(200, "OK", []models.User{}), nil)

	ws.Route(ws.GET("/{user-id}").To(biu.Handle(ctl.findUser)).
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").
			DataType("integer").DefaultValue("1")).
		Writes(models.User{}).
		Returns(200, "OK", models.User{}).
		Returns(404, "Not Found", nil), nil)

	ws.Route(ws.PUT("/{user-id}").To(biu.Handle(ctl.updateUser)).
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Reads(models.User{}), nil)

	ws.Route(ws.PUT("/").To(biu.Handle(ctl.createUser)).
		Doc("create a user").
		Reads(models.User{}), nil)

	ws.Route(ws.DELETE("/{user-id}").To(biu.Handle(ctl.removeUser)).
		Doc("delete a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")), nil)
}

func (ctl UserResource) findAllUsers(ctx biu.Ctx) {
	biu.Debug("UserResource.findAllUsers", biu.Log())
	ctx.ResponseJSON(models.User{})
}

func (ctl UserResource) findUser(ctx biu.Ctx) {
	id := ctx.PathParameter("user-id")
	biu.Debug("UserResource.findUser", biu.Log().Str("id", id))
}

func (ctl *UserResource) updateUser(ctx biu.Ctx) {
	usr := new(models.User)
	err := ctx.ReadEntity(&usr)
	biu.Debug("UserResource.updateUser", biu.Log().Interface("user", usr).Err(err))
}

func (ctl *UserResource) createUser(ctx biu.Ctx) {
	usr := models.User{ID: ctx.PathParameter("user-id")}
	err := ctx.ReadEntity(&usr)
	biu.Debug("UserResource.createUser", biu.Log().Interface("user", usr).Err(err))
}

func (ctl *UserResource) removeUser(ctx biu.Ctx) {
	id := ctx.PathParameter("user-id")
	biu.Debug("UserResource.removeUser", biu.Log().Str("id", id))
}
