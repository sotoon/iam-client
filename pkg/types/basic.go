package types

type ResponseError struct {
	Error    string   `json:"message"`
	Invalids []string `json:"invalids,omitempty"`
}
