package common

type ErrorJSON struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
