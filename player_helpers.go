package omxplayer

import (
	"fmt"
)

// dbusGetBool calls a D-Bus method that will return a boolean value.
func dbusGetBool(path string) (bool, error) {
	return false, fmt.Errorf("omxplayer: %s not implemented yet", path)
}
