package e

import (
    "errors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

const (
    ParameterError        = 4000
    DBRecordNotFound      = 4040
    DBRecordAlreadyExists = 4100
    SystemError           = 5000
    DBError               = 5010
    HttpError             = 5100
    UserLoginError        = 9000
    UserTokenError        = 9010
    UserTokenResolveError = 9011
    UserLogoutError       = 9020
    UserNotLogin          = 9999 //用户未登录
)

var statusText = map[int]string{
    ParameterError:        "参数校验失败，请检查参数",
    DBRecordNotFound:      "数据库记录不存在",
    DBRecordAlreadyExists: "数据库记录已经存在",
    SystemError:           "系统错误",
    DBError:               "数据库操作错误",
    HttpError:             "http请求失败",
    UserLoginError:        "用户登陆错误",
    UserTokenError:        "用户token错误",
    UserTokenResolveError: "用户token解析失败",
    UserLogoutError:       "用户退出登录失败",
}

//Error 自定义Error
type Error struct {
    error
    code int
    data gin.H
}

func (t *Error) Error() string {
    if t.error == nil {
        return ""
    }
    return t.error.Error()
}

func (t *Error) Code() int {
    return t.code
}

func (t *Error) SetMsg(msg string) *Error {
    t.error = errors.New(msg)
    return t
}

func (t *Error) SetData(data gin.H) *Error {
    t.data = data
    return t
}

func (t *Error) Json() gin.H {
    res := gin.H{"code": t.code}
    if msg := t.Error(); msg != "" {
        res["msg"] = msg
    }
    if t.data != nil {
        res["data"] = t.data
    }
    return res
}

func New(code int, param ...string) *Error {
    msg := statusText[code]

    if len(param) > 0 {
        if len(msg) > 0 {
            msg += ":"
        }
        msg += param[0]
    }

    e := new(Error)
    e.code = code
    return e.SetMsg(msg)
}

// FormatDBError 如果数据库报错，无论是数据库错误还是未找到都要报错，则使用这个
func FormatDBError(err error, param ...string) *Error {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return New(DBRecordNotFound, param...)
    }
    return New(DBError, param...)
}

func ValidateError() gin.H {
    return New(ParameterError).Json()
}
