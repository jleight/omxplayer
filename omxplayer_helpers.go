package omxplayer

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
)

// removeFile removes the specified file. Errors are ignored.
func removeFile(path string) {
	log.WithFields(log.Fields{
		"path": path,
	}).Debug("omxplayer: removing file")
	os.Remove(path)
}

// waitForFile waits for the specified file to exist before returning. If the an
// error, other than the file not existing, occurs, the error is returned. If,
// after 100 attempts, the file does not exist, an error is returned.
func waitForFile(path string) error {
	log.WithFields(log.Fields{
		"file": path,
	}).Debug("omxplayer: waiting for file")
	for i := 0; i < 100; i++ {
		_, err := os.Stat(path)
		if err == nil || !os.IsNotExist(err) {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("omxplayer: file does not exist: %s", path)
}
