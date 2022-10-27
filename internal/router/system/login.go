package system

import (
	miniapp "app-instance/internal/mini-app"
	"app-instance/internal/pkg/gojwt"
	"app-instance/internal/pkg/render"

	"github.com/gin-gonic/gin"
)

type LoginBind struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginBind
	if err := c.ShouldBind(&form); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	// 验证用户名和密码
	user := &UserForm{}
	if form.UserName != "" {
		user.Username = form.UserName
	}

	if err := user.Detail(); err != nil {
		render.CustomerError(c, render.CODE_ERR_LOGIN_FAILED, "username or password incorrect")
		return
	}
	if form.PassWord != user.Password {
		render.CustomerError(c, render.CODE_ERR_LOGIN_FAILED, "password incorrect")
		return
	}

	userinfo := &gojwt.UserInfo{
		Id:   user.ID,
		Name: user.Username,
	}
	tokenString, _ := gojwt.GenerateToken(*userinfo)
	render.JSON(c, tokenString)
	miniapp.App.Logger.Infof("user: %s Login\n", user.Username)
}
