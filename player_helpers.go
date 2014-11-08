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
