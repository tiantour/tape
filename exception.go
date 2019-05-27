package tape

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

// Exce exce
var Exce = new(Exception)

// Exception throw exception
type Exception struct{}

// OK throw exception if !ok
func (e *Exception) OK(ok bool, message interface{}) {
	if !ok {
		panic(message)
	}
}

// Err throw exception if err!=nil
func (e *Exception) Err(err error, message interface{}) {
	if err != nil {
		if message == "" {
			switch {
			case strings.HasPrefix(err.Error(), "Error 1046"):
				message = "查询表名错误"
			case strings.HasPrefix(err.Error(), "Error 1049"):
				message = "查询库名错误"
			case strings.HasPrefix(err.Error(), "Error 1054"):
				message = "查询列名错误"
			case strings.HasPrefix(err.Error(), "Error 1062"):
				message = "写入数据重复"
			case strings.HasPrefix(err.Error(), "Error 1064"):
				message = "查询语句错误"
			case strings.HasPrefix(err.Error(), "sql:"):
				message = "查询数据失败"
			case strings.HasPrefix(err.Error(), "code=400,message=Unmarshal type error"):
				message = "数据类型错误"
			case strings.HasPrefix(err.Error(), "rpc error: code = Unavailable desc"):
				message = "连接服务错误"
			case strings.HasPrefix(err.Error(), "rpc error: code = Unknown desc"):
				message = "调用服务错误"
			case strings.HasPrefix(err.Error(), "rpc error: code = Unimplemented desc"):
				message = "解析服务错误"
			default:
				message = err.Error()
			}
		}
		fmt.Println(reflect.TypeOf(err), err.Error(), message)
		panic(message)
	}
}

// Catch throw exception if recover err
func (e *Exception) Catch(ctx echo.Context) func() {
	return func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusOK, Hash{
				STATUS:  FAIL,
				MESSAGE: r,
			})
		}
	}
}
