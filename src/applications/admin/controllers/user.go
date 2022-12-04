package controllers

import (
    "github.com/caijw-go/library/param"
    "github.com/gin-gonic/gin"
    "net/http"
    "service/applications/admin/middlewares"
    "service/applications/admin/services/user"
    "service/library/e"
)

type UserController struct {
}

func (t *UserController) Init(g *gin.RouterGroup) {
    g.POST("/login", t.login)
    g.GET("/info", middlewares.CheckAuth(), t.info)
    g.GET("/logout", middlewares.CheckAuth(), t.logout)
}

type loginParameter struct {
    Username string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required"`
}

func (t *UserController) login(c *gin.Context) {
    var parameter loginParameter
    if param.Validate(c, &parameter) != nil {
        c.JSON(http.StatusOK, e.ValidateError())
        return
    }
    token, err := user.Login(parameter.Username, parameter.Password)
    if err != nil {
        c.JSON(http.StatusOK, err.Json())
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code": 0,
        "data": gin.H{
            "username": parameter.Username,
            "token":    token,
        },
    })
    return
}
func (t *UserController) info(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "code": 0,
        "data": gin.H{
            "username": user.GetUser(c),
        },
    })
    return
}

func (t *UserController) logout(c *gin.Context) {
    if user.Logout(c) {
        c.Status(http.StatusNoContent)
        return
    }
    c.JSON(http.StatusOK, e.New(e.UserLogoutError).Json())
    return
}
