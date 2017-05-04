package tape

const (
	// SUCCESS show status 200 when success
	SUCCESS = 200

	// FAIL show status 404 when fail
	FAIL = 404
)

var (
	// Exce exce
	Exce = new(exception)
)

type (
	// exception throw exception
	exception struct{}

	// Map map
	Map map[string]interface{}

	// Hash map
	Hash struct {
		STATUS  int         `json:"status"`            // show sataus
		RESULT  interface{} `json:"result,omitempty"`  // show result when success
		MESSAGE interface{} `json:"message,omitempty"` // show message when fail
	}
)
