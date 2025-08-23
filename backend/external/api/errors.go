package api

import "strings"

type apiError struct {
	messages []string `json:"messages"`
}

func (e apiError) Error() string {
	return "Error: " + strings.Join(e.messages, ", ")
}
