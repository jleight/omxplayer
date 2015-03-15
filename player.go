package omxplayer

import (
	log "github.com/Sirupsen/logrus"
	"github.com/guelfey/go.dbus"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const (
	ifaceProps     = "org.freedesktop.DBus.Properties"
	ifaceOmxRoot   = ifaceMpris
	ifaceOmxPlayer = ifaceOmxRoot + ".Player"

	cmdQuit                 = ifaceOmxRoot + ".Quit"
	propCanQuit             = ifaceProps + ".CanQuit"
	propFullscreen          = ifaceProps + ".Fullscreen"
	propCanSetFullscreen    = ifaceProps + ".CanSetFullscreen"
	propCanRaise            = ifaceProps + ".CanRaise"
	propHasTrackList        = ifaceProps + ".HasTrackList"
	propIdentity            = ifaceProps + ".Identity"
	propSupportedUriSchemes = ifaceProps + ".SupportedUriSchemes"
	propSupportedMimeTypes  = ifaceProps + ".SupportedMimeTypes"
	propCanGoNext           = ifaceProps + ".CanGoNext"
	propCanGoPrevious       = ifaceProps + ".CanGoPrevious"
	propCanSeek             = ifaceProps + ".CanSeek"
	propCanControl          = ifaceProps + ".CanControl"
	propCanPlay             = ifaceProps + ".CanPlay"
	propCanPause            = ifaceProps + ".CanPause"
	cmdNext                 = ifaceOmxPlayer + ".Next"
	cmdPrevious             = ifaceOmxPlayer + ".Previous"
	cmdPause                = ifaceOmxPlayer + ".Pause"
	cmdPlayPause            = ifaceOmxPlayer + ".PlayPause"
	cmdStop                 = ifaceOmxPlayer + ".Stop"
	cmdSeek                 = ifaceOmxPlayer + ".Seek"
	cmdSetPosition          = ifaceOmxPlayer + ".SetPosition"
	propPlaybackStatus      = ifaceProps + ".PlaybackStatus"
	cmdVolume               = ifaceProps + ".Volume"
	cmdMute                 = ifaceProps + ".Mute"
	cmdUnmute               = ifaceProps + ".Unmute"
	propPosition            = ifaceProps + ".Position"
	propAspect              = ifaceProps + ".Aspect"
	propVideoStreamCount    = ifaceProps + ".VideoStreamCount"
	propResWidth            = ifaceProps + ".ResWidth"
	propResHeight           = ifaceProps + ".ResHeight"
	propDuration            = ifaceProps + ".Duration"
	propMinimumRate         = ifaceProps + ".MinimumRate"
	propMaximumRate         = ifaceProps + ".MaximumRate"
	cmdListSubtitles        = ifaceOmxPlayer + ".ListSubtitles"
	cmdHideVideo            = ifaceOmxPlayer + ".HideVideo"
	cmdUnHideVideo          = ifaceOmxPlayer + ".UnHideVideo"
	cmdListAudio            = ifaceOmxPlayer + ".ListAudio"
	cmdListVideo            = ifaceOmxPlayer + ".ListVideo"
	cmdSelectSubtitle       = ifaceOmxPlayer + ".SelectSubtitle"
	cmdSelectAudio          = ifaceOmxPlayer + ".SelectAudio"
	cmdShowSubtitles        = ifaceOmxPlayer + ".ShowSubtitles"
	cmdHideSubtitles        = ifaceOmxPlayer + ".HideSubtitles"
	cmdAction               = ifaceOmxPlayer + ".Action"
)

type Player struct {
	command    *exec.Cmd
	connection *dbus.Conn
	bus        *dbus.Object
	ready      bool
}

// IsRunning checks to see if the OMXPlayer process is running. If it is, the
// function returns true, otherwise it returns false.
func (p *Player) IsRunning() bool {
	pid := p.command.Process.Pid
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	} else {
		err = process.Signal(syscall.Signal(0))
		return err != nil
	}
}

// IsReady checks to see if the Player instance is ready to accept D-Bus
// commands. If the player is ready and can accept commands, the function
// returns true, otherwise it returns false.
func (p *Player) IsReady() bool {
	if p.ready {
		return true
	}

	_, err := p.CanQuit()
	p.ready = err != nil
	return p.ready
}

// WaitForReady waits until the Player instance is ready to accept D-Bus
// commands and then returns.
func (p *Player) WaitForReady() {
	for !p.IsReady() {
		time.Sleep(50 * time.Millisecond)
	}
}

func (p *Player) Quit() error {
	return dbusCall(p.bus, cmdQuit)
}

func (p *Player) CanQuit() (bool, error) {
	return dbusGetBool(p.bus, propCanQuit)
}

func (p *Player) Fullscreen() (bool, error) {
	return dbusGetBool(p.bus, propFullscreen)
}

func (p *Player) CanSetFullscreen() (bool, error) {
	return dbusGetBool(p.bus, propCanSetFullscreen)
}

func (p *Player) CanRaise() (bool, error) {
	return dbusGetBool(p.bus, propCanRaise)
}

func (p *Player) HasTrackList() (bool, error) {
	return dbusGetBool(p.bus, propHasTrackList)
}

func (p *Player) Identity() (string, error) {
	return dbusGetString(p.bus, propIdentity)
}

func (p *Player) SupportedUriSchemes() ([]string, error) {
	return dbusGetStringArray(p.bus, propSupportedUriSchemes)
}

func (p *Player) SupportedMimeTypes() ([]string, error) {
	return dbusGetStringArray(p.bus, propSupportedMimeTypes)
}

func (p *Player) CanGoNext() (bool, error) {
	return dbusGetBool(p.bus, propCanGoNext)
}

func (p *Player) CanGoPrevious() (bool, error) {
	return dbusGetBool(p.bus, propCanGoPrevious)
}

func (p *Player) CanSeek() (bool, error) {
	return dbusGetBool(p.bus, cmdSeek)
}

func (p *Player) CanControl() (bool, error) {
	return dbusGetBool(p.bus, propCanControl)
}

func (p *Player) CanPlay() (bool, error) {
	return dbusGetBool(p.bus, propCanPlay)
}

func (p *Player) CanPause() (bool, error) {
	return dbusGetBool(p.bus, propCanPause)
}

func (p *Player) Next() error {
	return dbusCall(p.bus, cmdNext)
}

func (p *Player) Previous() error {
	return dbusCall(p.bus, cmdPrevious)
}

func (p *Player) Pause() error {
	return dbusCall(p.bus, cmdPause)
}

func (p *Player) PlayPause() error {
	return dbusCall(p.bus, cmdPlayPause)
}

func (p *Player) Stop() error {
	return dbusCall(p.bus, cmdStop)
}

func (p *Player) Seek(amount int64) (int64, error) {
	log.WithFields(log.Fields{
		"path":        cmdSeek,
		"paramAmount": amount,
	}).Debug("omxplayer: dbus call")
	call := p.bus.Call(cmdSeek, 0, amount)
	if call.Err != nil {
		return 0, call.Err
	}
	return call.Body[0].(int64), nil
}

func (p *Player) SetPosition(path string, position int64) (int64, error) {
	log.WithFields(log.Fields{
		"path":          cmdSetPosition,
		"paramPath":     path,
		"paramPosition": position,
	}).Debug("omxplayer: dbus call")
	call := p.bus.Call(cmdSetPosition, 0, path, position)
	if call.Err != nil {
		return 0, call.Err
	}
	return call.Body[0].(int64), nil
}

func (p *Player) PlaybackStatus() (string, error) {
	return dbusGetString(p.bus, propPlaybackStatus)
}

func (p *Player) Volume(volume ...float64) (float64, error) {
	log.WithFields(log.Fields{
		"path":        cmdVolume,
		"paramVolume": volume,
	}).Debug("omxplayer: dbus call")
	if len(volume) == 0 {
		return dbusGetFloat64(p.bus, cmdVolume)
	}
	call := p.bus.Call(cmdVolume, 0, volume[0])
	if call.Err != nil {
		return 0, call.Err
	}
	return call.Body[0].(float64), nil
}

func (p *Player) Mute() error {
	return dbusCall(p.bus, cmdMute)
}

func (p *Player) Unmute() error {
	return dbusCall(p.bus, cmdUnmute)
}

func (p *Player) Position() (int64, error) {
	return dbusGetInt64(p.bus, propPosition)
}

func (p *Player) Aspect() (float64, error) {
	return dbusGetFloat64(p.bus, propAspect)
}

func (p *Player) VideoStreamCount() (int64, error) {
	return dbusGetInt64(p.bus, propVideoStreamCount)
}

func (p *Player) ResWidth() (int64, error) {
	return dbusGetInt64(p.bus, propResWidth)
}

func (p *Player) ResHeight() (int64, error) {
	return dbusGetInt64(p.bus, propResHeight)
}

func (p *Player) Duration() (int64, error) {
	return dbusGetInt64(p.bus, propDuration)
}

func (p *Player) MinimumRate() (float64, error) {
	return dbusGetFloat64(p.bus, propMinimumRate)
}

func (p *Player) MaximumRate() (float64, error) {
	return dbusGetFloat64(p.bus, propMaximumRate)
}

func (p *Player) ListSubtitles() ([]string, error) {
	return dbusGetStringArray(p.bus, cmdListSubtitles)
}

func (p *Player) HideVideo() error {
	return dbusCall(p.bus, cmdHideVideo)
}

func (p *Player) UnHideVideo() error {
	return dbusCall(p.bus, cmdUnHideVideo)
}

func (p *Player) ListAudio() ([]string, error) {
	return dbusGetStringArray(p.bus, cmdListAudio)
}

func (p *Player) ListVideo() ([]string, error) {
	return dbusGetStringArray(p.bus, cmdListVideo)
}

func (p *Player) SelectSubtitle(index int32) (bool, error) {
	log.WithFields(log.Fields{
		"path":       cmdSelectSubtitle,
		"paramIndex": index,
	}).Debug("omxplayer: dbus call")
	call := p.bus.Call(cmdSelectSubtitle, 0, index)
	if call.Err != nil {
		return false, call.Err
	}
	return call.Body[0].(bool), nil
}

func (p *Player) SelectAudio(index int32) (bool, error) {
	log.WithFields(log.Fields{
		"path":       cmdSelectAudio,
		"paramIndex": index,
	}).Debug("omxplayer: dbus call")
	call := p.bus.Call(cmdSelectAudio, 0, index)
	if call.Err != nil {
		return false, call.Err
	}
	return call.Body[0].(bool), nil
}

func (p *Player) ShowSubtitles() error {
	return dbusCall(p.bus, cmdShowSubtitles)
}

func (p *Player) HideSubtitles() error {
	return dbusCall(p.bus, cmdHideSubtitles)
}

func (p *Player) Action(action int32) error {
	log.WithFields(log.Fields{
		"path":        cmdAction,
		"paramAction": action,
	}).Debug("omxplayer: dbus call")
	return p.bus.Call(cmdAction, 0, action).Err
}
