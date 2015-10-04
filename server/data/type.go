package data

import (
	"encoding/json"
	"fmt"
)

//Type represents type of bundle store
type Type int

const (
	//JAVASCRIPT represents source code in JS
	JAVASCRIPT Type = iota
	//IMAGE represents picture and image types
	IMAGE
)

var (
	typeNameToValue = map[string]Type{
		"JAVASCRIPT": JAVASCRIPT,
		"IMAGE":      IMAGE,
	}

	typeValueToName = map[Type]string{
		JAVASCRIPT: "JAVASCRIPT",
		IMAGE:      "IMAGE",
	}
)

//MarshalJSON for type Type
func (a Type) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(a).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := typeValueToName[a]
	if !ok {
		return nil, fmt.Errorf("invalid Type: %d", a)
	}
	return json.Marshal(s)
}

//UnmarshalJSON for type Type
func (a *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Type should be a string, got %s", data)
	}
	v, ok := typeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Type %q", s)
	}
	*a = v
	return nil
}
