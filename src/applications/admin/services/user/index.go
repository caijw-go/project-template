package user

import (
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/log"
    "github.com/gin-gonic/gin"
    "service/library/business"
    "service/library/e"
    "service/library/redis_key"
)

func Login(username, password string) (string, *e.Error) {
    return "xxx", nil
}

func ResolveUser(token string) (string, *e.Error) {
    if len(token) != 32 {
        log.Error("admin resolveUser token error", token)
        return "", e.New(e.UserTokenError)
    }
    bytes, err := base.Redis("admin").Get(redis_key.Get(redis_key.AdminLoggedUser, token)).Bytes()
    if err != nil {
        log.Error("admin resolveUser redis get error", err.Error())
        return "", e.New(e.UserTokenResolveError, err.Error())
    }
    return string(bytes), nil
}

func GetUser(c *gin.Context) string {
    username, _ := c.Get(business.ContextUserKey)
    return username.(string)
}

func Logout(c *gin.Context) bool {
    token, err := GetTokenFromHeader(c)
    if err != nil {
        log.Error("admin Logout get cookie error", err.Error())
        return false
    }
    base.Redis("admin").Del(redis_key.Get(redis_key.AdminLoggedUser, token))
    return true
}
