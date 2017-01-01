package tape

import (
	"github.com/kataras/iris"
	"github.com/labstack/echo"
)

// exception exception
type exception struct{}

// Exce new exce
var Exce = new(exception)

// Catch catch exce
// date 2017-01-01
// author andy.jiang
func (e exception) Catch(ctx interface{}) func() {
	return func() {
		if r := recover(); r != nil {
			switch ctx.(type) {
			case *iris.Context:
				ctx.(*iris.Context).JSON(200, Hash{
					STATUS:  FAIL,
					MESSAGE: r,
				})
			case echo.Context:
				ctx.(echo.Context).JSON(200, Hash{
					STATUS:  FAIL,
					MESSAGE: r,
				})
			}
		}
	}
}

// OK check ok
// date 2017-01-01
// author andy.jiang
func (e exception) OK(ok bool, message interface{}) {
	if !ok {
		panic(message)
	}
}

// Err check err
// date 2017-01-01
// author andy.jiang
func (e exception) Err(err error, message interface{}) {
	if err != nil {
		if message == "" {
			message = err.Error()
		}
		panic(message)
	}
}
