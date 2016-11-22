package tape

import (
	"log"
	"strings"

	"github.com/kataras/iris"
)

// OK throw exception if !ok
func (e *exception) OK(ok bool, message interface{}) {
	if !ok {
		panic(message)
	}
}

// Err throw exception if err!=nil
func (e *exception) Err(err error, message interface{}) {
	if err != nil {
		if message == "" {
			switch {
			case strings.HasPrefix(err.Error(), "sql:"):
				message = "数据查询失败"
			case strings.HasPrefix(err.Error(), "Error 1"):
				message = "数据操作异常"
			default:
				message = err.Error()
			}
		}
		log.Println(err.Error(), message)
		panic(message)
	}
}

// Catch throw exception if recover err
func (e *exception) Catch(ctx *iris.Context) func() {
	return func() {
		if r := recover(); r != nil {
			ctx.JSON(200, Hash{
				STATUS:  FAIL,
				MESSAGE: r,
			})
		}
	}
}
