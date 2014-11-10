package omxplayer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/guelfey/go.dbus"
	"os"
	"os/exec"
	"strings"
)

const (
	envDisplay         = "DISPLAY"
	envDbusAddress     = "DBUS_SESSION_BUS_ADDRESS"
	envDbusPid         = "DBUS_SESSION_BUS_PID"
	prefixOmxDbusFiles = "/tmp/omxplayerdbus."
	suffixOmxDbusPid   = ".pid"
	ifaceMpris         = "org.mpris.MediaPlayer2"
	exeOxmPlayer       = "omxplayer"
	keyPause           = "p"
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

// getDbusConnection establishes and returns a D-Bus connection. The connection
// is made to the D-Bus service that has been set via the two `DBUS_*`
// environment variables. Since the connection's `Auth` method attempts to use
// Go's `os/user` package to get the current user's name and home directory, and
// `os/user` is not implemented for Linux-ARM, the `authMethods` parameter is
// specified explicitly rather than passing `nil`.
func getDbusConnection() (conn *dbus.Conn, err error) {
	authMethods := []dbus.Auth{
		dbus.AuthExternal(user),
		dbus.AuthCookieSha1(user, home),
	}

	log.Debug("omxplayer: opening dbus session")
	if conn, err = dbus.SessionBusPrivate(); err != nil {
		return
	}

	log.Debug("omxplayer: authenticating dbus session")
	if err = conn.Auth(authMethods); err != nil {
		return
	}

	log.Debug("omxplayer: initializing dbus session")
	err = conn.Hello()
	return
}

// setupDbusEnvironment sets the environment variables that are necessary to
// establish a D-Bus connection. If the connection's path or PID cannot be read,
// the associated error is returned.
func setupDbusEnvironment() (err error) {
	log.Debug("omxplayer: setting up dbus environment")

	path, err := getDbusPath()
	if err != nil {
		return
	}

	pid, err := getDbusPid()
	if err != nil {
		return
	}

	setEnv(envDisplay, ":0")
	setEnv(envDbusAddress, path)
	setEnv(envDbusPid, pid)
	return
}

// execOmxplayer starts a new OMXPlayer process and tells it to pause the video
// by passing a "p" on standard input.
func execOmxplayer(url string) (cmd *exec.Cmd, err error) {
	log.Debug("omxplayer: starting omxplayer process")

	cmd = exec.Command(exeOxmPlayer, "--no-osd", url)
	cmd.Stdin = strings.NewReader(keyPause)
	err = cmd.Start()
	return
}
