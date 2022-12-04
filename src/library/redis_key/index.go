package redis_key

import "fmt"

const (
    AdminLoggedUser = "admin:user:login:%s" //登陆用户的session key
)

func Get(key string, param ...interface{}) string {
    return fmt.Sprintf(key, param...)
}
