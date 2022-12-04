package user

import (
    "github.com/gin-gonic/gin"
    "service/library/e"
    "strings"
)

func GetTokenFromHeader(c *gin.Context) (string, *e.Error) {
    header := c.GetHeader("Authorization")
    tmp := strings.Split(header, " ")
    if len(tmp) != 2 || tmp[0] != "Basic" || len(tmp[1]) != 32 {
        return "", e.New(e.UserNotLogin)
    }
    return tmp[1], nil
}
