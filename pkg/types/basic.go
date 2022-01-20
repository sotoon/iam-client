package types

type ResponseError struct {
	Error    string   `json:"message"`
	Invalids []string `json:"invalids,omitempty"`
}

type RequestExecutionError struct {
	Err        error
	StatusCode int
	Data       []byte
}

func (ree *RequestExecutionError) Error() string {
	return ree.Err.Error()
}
