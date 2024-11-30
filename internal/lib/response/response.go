package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Body   string `json:"body,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(errorMessage string) Response {
	return Response{
		Status: StatusError,
		Error:  errorMessage,
	}
}
