package route

import (
	miniapp "app-instance/internal/mini-app"
	"app-instance/internal/router/middleware"
	reqApi "app-instance/internal/router/route/api"
	"app-instance/internal/router/system"
)

func RegisterRoute() {
	api := miniapp.App.Gin.Group("/api", middleware.ApiPriv())
	{
		api.POST(reqApi.USER_ADD, system.UserAdd)
		api.POST(reqApi.LOGIN, system.Login)
	}
}
