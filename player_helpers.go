package omxplayer

import (
	dbus "github.com/guelfey/go.dbus"
	log "github.com/sirupsen/logrus"
)

// dbusCall calls a D-Bus method that has no return value.
func dbusCall(bus *dbus.Object, path string) error {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	return bus.Call(path, 0).Err
}

// dbusGetBool calls a D-Bus method that will return a boolean value.
func dbusGetBool(bus *dbus.Object, path string) (bool, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	call := bus.Call(path, 0)
	if call.Err != nil {
		return false, call.Err
	}
	return call.Body[0].(bool), nil
}

// dbusGetFloat64 calls a D-Bus method that will return an int64 value.
func dbusGetFloat64(bus *dbus.Object, path string) (float64, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	call := bus.Call(path, 0)
	if call.Err != nil {
		return 0, call.Err
	}
	return call.Body[0].(float64), nil
}

// dbusGetInt64 calls a D-Bus method that will return an int64 value.
func dbusGetInt64(bus *dbus.Object, path string) (int64, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	call := bus.Call(path, 0)
	if call.Err != nil {
		return 0, call.Err
	}
	return call.Body[0].(int64), nil
}

// dbusGetString calls a D-Bus method that will return a string value.
func dbusGetString(bus *dbus.Object, path string) (string, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	call := bus.Call(path, 0)
	if call.Err != nil {
		return "", call.Err
	}
	return call.Body[0].(string), nil
}

// dbusGetStringArray calls a D-Bus method that will return a string array.
func dbusGetStringArray(bus *dbus.Object, path string) ([]string, error) {
	log.WithFields(log.Fields{"path": path}).Debug("omxplayer: dbus call")
	call := bus.Call(path, 0)
	if call.Err != nil {
		return nil, call.Err
	}
	return call.Body[0].([]string), nil
}
