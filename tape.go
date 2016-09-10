package tape

// E exception
var (
	E = &exception{}
)

// H hash
type (
	H         map[string]interface{}
	exception struct{}
)
