package system

import (
	"app-instance/internal/model"
	"app-instance/internal/pkg/render"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	ID       int    `form:"id"`
	Username string `form:"username" binding:"required"`
}

func UserAdd(c *gin.Context) {
	var userForm UserForm
	if err := c.ShouldBind(&userForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	userCreateOrUpdate(c, userForm)
}

func userCreateOrUpdate(c *gin.Context, userForm UserForm) {
	var (
		checkUsername *model.User
		exists        bool
		err           error
	)

	checkUsername = &model.User{
		ID:       userForm.ID,
		Username: userForm.Username,
	}
	exists, err = checkUsername.UserCheckExists()
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	if exists {
		render.RepeatError(c, "username have exists")
		return
	}

	u := &model.User{
		ID:       userForm.ID,
		Username: userForm.Username,
	}
	if err := u.CreateOrUpdate(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}
