package route

import (
	miniapp "app-instance/internal/mini-app"
	reqApi "app-instance/internal/router/route/api"
	"app-instance/internal/router/system"
)

func RegisterRoute() {
	api := miniapp.App.Gin.Group("/api")
	{
		api.POST(reqApi.USER_ADD, system.UserAdd)
	}
}
