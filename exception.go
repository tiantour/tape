package tape

import (
	"log"

	"github.com/labstack/echo"
	"github.com/tiantour/conf"
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
		if conf.Options.Server.Debug == true {
			message = err.Error()
		}
		panic(message)
	}
}

// Catch
func (e *exception) Catch(ctx echo.Context) func() {
	return func() {
		if r := recover(); r != nil {
			if conf.Options.Server.Debug == true {
				log.Println(r)
			}
			ctx.(echo.Context).JSON(200, H{
				"status":  "404",
				"message": r,
			})

		}
	}
}
