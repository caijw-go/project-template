package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "service/applications/admin/services/user"
    "service/library/business"
    "service/library/e"
)

func CheckAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := user.GetTokenFromHeader(c)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, err.Json())
            return
        }
        if username, err := user.ResolveUser(token); err != nil {
            c.AbortWithStatusJSON(http.StatusOK, e.New(e.UserNotLogin).Json())
            return
        } else {
            c.Set(business.ContextUserKey, username)
            c.Next()
        }
    }
}
