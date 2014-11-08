package omxplayer

import (
	"fmt"
)

// dbusCall calls a D-Bus method that has no return value.
func dbusCall(path string) error {
	return fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetBool calls a D-Bus method that will return a boolean value.
func dbusGetBool(path string) (bool, error) {
	return false, fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetFloat64 calls a D-Bus method that will return an int64 value.
func dbusGetFloat64(path string) (float64, error) {
	return 0, fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetInt64 calls a D-Bus method that will return an int64 value.
func dbusGetInt64(path string) (int64, error) {
	return 0, fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetString calls a D-Bus method that will return a string value.
func dbusGetString(path string) (string, error) {
	return "", fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetStringArray calls a D-Bus method that will return a string array.
func dbusGetStringArray(path string) ([]string, error) {
	return nil, fmt.Errorf("omxplayer: %s not implemented yet", path)
}
