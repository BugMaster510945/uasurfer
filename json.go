package uasurfer

import (
	"encoding/json"
)

// MarshalJSON blabla
func (i DeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// MarshalJSON blabla
func (i BrowserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// MarshalJSON blabla
func (i OSName) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// MarshalJSON blabla
func (i Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}
