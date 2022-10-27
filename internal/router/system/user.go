package system

import (
	miniapp "app-instance/internal/mini-app"
	"app-instance/internal/model"
	"app-instance/internal/pkg/render"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	ID       int    `form:"id"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
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
		Password: userForm.Password,
	}
	if err := u.CreateOrUpdate(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
	miniapp.App.Logger.Infof("user %s add success\n", userForm.Username)
}

func (user *UserForm) Detail() error {
	var where []model.WhereParam
	u := model.User{}

	if user.ID != 0 {
		where = append(where, model.WhereParam{
			Field:   "id",
			Prepare: user.ID,
		})
	}

	if user.Username != "" {
		where = append(where, model.WhereParam{
			Field:   "username",
			Prepare: user.Username,
		})
	}

	if ok := u.GetOne(model.QueryParam{
		Where: where,
	}); !ok {
		return errors.New("get user detail failed")
	}

	if u.ID == 0 {
		return errors.New("user not exists")
	}

	user.ID = u.ID
	user.Password = u.Password
	user.Username = u.Username

	return nil
}
