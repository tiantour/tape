# tape
a tape of echo,iris,gin with go

### how to use

**tape**

    tape.SUCCESS    // const 200

    tape.FAIL       // const 404

    tape.Map        // map[string]interface{}

    tape.Hash      // struct

**Exce**

    tape.Exce.Catch // catch exce

    tape.Exce.OK    // if !ok throw exce

    tape.Exce.Err   // if err!=nil throw exce


### demo

**Echo**

```
func demo(ctx echo.Context) error {
	// Catch exce
	defer tape.Exce.Catch(ctx)()
	data := map[string]string{
		"go":   "golang",
		"java": "javalang",
		"rust": "rustlang",
	}
	value, ok := data["python"]
	// check ok
	tape.Exce.OK(ok, "python is not exist")
	_, err := strconv.Atoi(value)
	// check err
	tape.Exce.Err(err, "")
	result := tape.Map{
		"data": data,
	}
	// return data
	return ctx.JSON(200, tape.Hash{
		STATUS: tape.SUCCESS,
		RESULT: result,
	})
}
```

**Iris**

```
func demo(ctx *iris.Context) {
	// Catch exce
	defer tape.Exce.Catch(ctx)()
	data := map[string]string{
		"go":   "golang",
		"java": "javalang",
		"rust": "rustlang",
	}
	value, ok := data["python"]
	// check ok
	tape.Exce.OK(ok, "python is not exist")
	_, err := strconv.Atoi(value)
	// check err
	tape.Exce.Err(err, "")
	result := tape.Map{
		"data": data,
	}
	// return data
	ctx.JSON(200, tape.Hash{
		STATUS: tape.SUCCESS,
		RESULT: result,
	})
}
```
