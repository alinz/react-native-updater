package data

import (
	"encoding/json"
	"fmt"
)

//Platform defines type of platform
type Platform int

const (
	//IOS represents apple ios devices
	IOS Platform = iota
	//ANDROID represents google android devices
	ANDROID
)

var (
	platformNameToValue = map[string]Platform{
		"IOS":     IOS,
		"ANDROID": ANDROID,
	}

	platformValueToName = map[Platform]string{
		IOS:     "IOS",
		ANDROID: "ANDROID",
	}
)

//MarshalJSON for type Platform
func (p Platform) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := platformValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid Platform: %d", p)
	}
	return json.Marshal(s)
}

//UnmarshalJSON for type Platform
func (p *Platform) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Platform should be a string, got %s", data)
	}
	v, ok := platformNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Platform %q", s)
	}
	*p = v
	return nil
}
