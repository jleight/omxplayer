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
