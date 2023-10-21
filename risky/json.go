package risky

import (
	"encoding/json"
)

func JSON[T any](data []byte) *T {
	var parsed T
	_ = json.Unmarshal(data, &parsed)
	return &parsed
}
