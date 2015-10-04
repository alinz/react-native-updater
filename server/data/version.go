package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//Version holds version in uint64
type Version uint64

//VersionEncode convert Version to string represantation
func VersionEncode(major uint64, minor uint64, patch uint64) string {
	return strconv.FormatUint(major, 10) + "." +
		strconv.FormatUint(minor, 10) + "." +
		strconv.FormatUint(patch, 10)
}

//VersionDecode gets an major.minor.pathc and returns array of parsed version
func VersionDecode(value string) ([]uint64, error) {
	versionSegments := strings.Split(value, ".")

	if len(versionSegments) != 3 {
		return nil, fmt.Errorf("Version should have 3 parts, got %d", len(versionSegments))
	}

	major, err := strconv.ParseUint(versionSegments[0], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Major part of Version is not parrsable")
	}

	minor, err := strconv.ParseUint(versionSegments[1], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("Minor part of Version is not parrsable")
	}

	patch, err := strconv.ParseUint(versionSegments[2], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Patch part of Version is not parrsable")
	}

	return []uint64{major, minor, patch}, nil
}

//MarshalJSON for type Platform
func (v Version) MarshalJSON() ([]byte, error) {
	value := uint64(v)

	major := (value & 0xffff000000000000) >> 48
	minor := (value & 0x0000ffff00000000) >> 32
	patch := (value & 0x00000000ffffffff)

	version := VersionEncode(major, minor, patch)

	return json.Marshal(version)
}

//UnmarshalJSON for type Platform
func (v *Version) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Version should be a string, got %s", data)
	}

	version, err := VersionDecode(s)

	if err != nil {
		return err
	}

	var value uint64

	version[0] = (version[0] << 48) & 0xffff000000000000
	version[1] = (version[1] << 32) & 0x0000ffff00000000
	version[2] = version[2] & 0x00000000ffffffff

	value = version[0] | version[1] | version[2]

	*v = Version(value)
	return nil
}
