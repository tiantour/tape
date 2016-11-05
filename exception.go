package tape

import (
	"log"
	"strings"

	"github.com/labstack/echo"
)

// OK
func (e *exception) OK(ok bool, message interface{}) {
	if !ok {
		panic(message)
	}
}

// Err
func (e *exception) Err(err error, message interface{}) {
	if err != nil {
		if message == "" {
			if strings.HasPrefix(err.Error(), "sql:") {
				message = "数据查询失败"
			} else if strings.HasPrefix(err.Error(), "Error 1") {
				message = "数据查询异常"
			} else {
				message = err.Error()
			}
		}
		panic(message)
	}
}

// Catch
func (e *exception) Catch(ctx echo.Context) func() {
	return func() {
		if r := recover(); r != nil {
			log.Println(r)
			ctx.(echo.Context).JSON(200, H{
				"status":  "404",
				"message": r,
			})

		}
	}
}
