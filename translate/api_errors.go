package deepl

import (
	"fmt"
)

type APIError struct {
	StatusCode int
	Body       string
}

func (e APIError) Error() string {
	return fmt.Sprintf("API Error: status code %d, body %s", e.StatusCode, e.Body)
}
