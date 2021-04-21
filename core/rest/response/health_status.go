package response

import (
	"bytes"
	"encoding/json"
)

type HealthStatus int

const (
	Ok HealthStatus = iota
	Partial
	Down
)

func (s HealthStatus) String() string {
	return toString[s]
}

var toString = map[HealthStatus]string{
	Ok:      "OK",
	Partial: "PARTIAL",
	Down:    "DOWN",
}

var toID = map[string]HealthStatus{
	"OK":      Ok,
	"PARTIAL": Partial,
	"DOWN":    Down,
}

// MarshalJSON marshals the enum as a quoted json string
func (s HealthStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals a quoted json string to the enum value
func (s *HealthStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toID[j]
	return nil
}
