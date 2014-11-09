package omxplayer

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

// removeFile removes the specified file. Errors are ignored.
func removeFile(path string) {
	log.WithFields(log.Fields{
		"path": path,
	}).Debug("omxplayer: removing file")
	os.Remove(path)
}
