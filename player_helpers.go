package omxplayer

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/guelfey/go.dbus"
)

// dbusCall calls a D-Bus method that has no return value.
func dbusCall(bus *dbus.Object, path string) error {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetBool calls a D-Bus method that will return a boolean value.
func dbusGetBool(bus *dbus.Object, path string) (bool, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return false, fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetFloat64 calls a D-Bus method that will return an int64 value.
func dbusGetFloat64(bus *dbus.Object, path string) (float64, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return 0, fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetInt64 calls a D-Bus method that will return an int64 value.
func dbusGetInt64(bus *dbus.Object, path string) (int64, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return 0, fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetString calls a D-Bus method that will return a string value.
func dbusGetString(bus *dbus.Object, path string) (string, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return "", fmt.Errorf("omxplayer: %s not implemented yet", path)
}

// dbusGetStringArray calls a D-Bus method that will return a string array.
func dbusGetStringArray(bus *dbus.Object, path string) ([]string, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return nil, fmt.Errorf("omxplayer: %s not implemented yet", path)
}
