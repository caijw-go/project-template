package admin

import (
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/server"
    "github.com/gin-gonic/gin"
    "service/applications/admin/controllers"
)

func Init() {
    server.Create(server.Config{
        AppName:        "admin",
        Address:        base.Config().GetString("server.address"),
        FuncController: routes,
    })
}

func routes(engine *gin.Engine) {
    new(controllers.UserController).Init(engine.Group("/user"))
}
