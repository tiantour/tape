# tape
a tape of echo,iris with go

### how to use

Exce

    tape.Exce.OK    // if !ok throw exce

    tape.Exce.Err   // if err!=nil throw exce

    tape.Exce.Catch // catch exce

Value

    tape.SUCCESS    // const "200"

    tape.FAIL       // const "404"

    tape.Map        // map[string]interface{}

    tape.Hash       // struct


### demo

```
func demo(ctx echo.Context) error {

	defer tape.Exce.Catch(ctx)()

	data := map[string]string{
		"go":   "golang",
		"java": "javalang",
		"rust": "rustlang",
	}

	value, ok := data["python"]
	tape.Exce.OK(ok, "python is not exist")

	_, err := strconv.Atoi(value)
	tape.Exce.Err(err, "")

	result := tape.Map{
		"data": data,
	}

	return ctx.JSON(200, tape.Hash{
		STATUS: tape.SUCCESS,
		RESULT: result,
	})
}

```
