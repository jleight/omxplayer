package omxplayer

import (
	"os"
)

const (
	prefixOmxDbusFiles = "/tmp/omxplayerdbus."
	suffixOmxDbusPid   = ".pid"
	ifaceMpris         = "org.mpris.MediaPlayer2"
)

var (
	user            string
	home            string
	fileOmxDbusPath string
	fileOmxDbusPid  string
)

func init() {
	SetUser(os.Getenv("USER"), os.Getenv("HOME"))
}

func SetUser(u, h string) {
	user = u
	home = h
	fileOmxDbusPath = prefixOmxDbusFiles + user
	fileOmxDbusPid = prefixOmxDbusFiles + user + suffixOmxDbusPid
}

// New returns a new Player instance that can be used to control an OMXPlayer
// instance that is playing the video located at the specified URL.
func New(url string) (player *Player, err error) {
	removeDbusFiles()
	return
}

// removeDbusFiles removes the files that OMXPlayer creates containing the D-Bus
// path and PID. This ensures that when the path and PID are read in, the new
// files are read instead of the old ones.
func removeDbusFiles() {
	removeFile(fileOmxDbusPath)
	removeFile(fileOmxDbusPid)
}

// getDbusPath reads the D-Bus path from the file OMXPlayer writes it's path to.
// If the file cannot be read, it returns an error, otherwise it returns the
// path as a string.
func getDbusPath() (string, error) {
	if err := waitForFile(fileOmxDbusPath); err != nil {
		return "", err
	}
	return readFile(fileOmxDbusPath)
}

// getDbusPath reads the D-Bus PID from the file OMXPlayer writes it's PID to.
// If the file cannot be read, it returns an error, otherwise it returns the
// PID as a string.
func getDbusPid() (string, error) {
	if err := waitForFile(fileOmxDbusPid); err != nil {
		return "", err
	}
	return readFile(fileOmxDbusPid)
}
