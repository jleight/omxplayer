package omxplayer

import (
	"os"
	"os/exec"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	dbus "github.com/guelfey/go.dbus"
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
	propSupportedURISchemes = ifaceProps + ".SupportedUriSchemes"
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

// The Player struct provides access to all of omxplayer's D-Bus methods.
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
	}

	err = process.Signal(syscall.Signal(0))
	return err != nil
}

// IsReady checks to see if the Player instance is ready to accept D-Bus
// commands. If the player is ready and can accept commands, the function
// returns true, otherwise it returns false.
func (p *Player) IsReady() bool {
	if p.ready {
		return true
	}

	result, err := p.CanQuit()
	if err == nil && result {
		p.ready = true
	}
	return p.ready
}

// WaitForReady waits until the Player instance is ready to accept D-Bus
// commands and then returns.
func (p *Player) WaitForReady() {
	for !p.IsReady() {
		time.Sleep(50 * time.Millisecond)
	}
}

// Quit stops the currently playing video and terminates the omxplayer process.
// See https://github.com/popcornmix/omxplayer#quit for more details.
func (p *Player) Quit() error {
	return dbusCall(p.bus, cmdQuit)
}

// CanQuit returns true if the player can quit, false otherwise. See
// https://github.com/popcornmix/omxplayer#canquit for more details.
func (p *Player) CanQuit() (bool, error) {
	return dbusGetBool(p.bus, propCanQuit)
}

// Fullscreen returns true if the player is fullscreen, false otherwise. See
// https://github.com/popcornmix/omxplayer#fullscreen for more details.
func (p *Player) Fullscreen() (bool, error) {
	return dbusGetBool(p.bus, propFullscreen)
}

// CanSetFullscreen returns true if the player can be set to fullscreen, false
// otherwise. See https://github.com/popcornmix/omxplayer#cansetfullscreen for
// more details.
func (p *Player) CanSetFullscreen() (bool, error) {
	return dbusGetBool(p.bus, propCanSetFullscreen)
}

// CanRaise returns true if the player can be brought to the front, false
// otherwise. See https://github.com/popcornmix/omxplayer#canraise for more
// details.
func (p *Player) CanRaise() (bool, error) {
	return dbusGetBool(p.bus, propCanRaise)
}

// HasTrackList returns true if the player has a track list, false otherwise.
// See https://github.com/popcornmix/omxplayer#hastracklist for more details.
func (p *Player) HasTrackList() (bool, error) {
	return dbusGetBool(p.bus, propHasTrackList)
}

// Identity returns the name of the player instance. See
// https://github.com/popcornmix/omxplayer#identity for more details.
func (p *Player) Identity() (string, error) {
	return dbusGetString(p.bus, propIdentity)
}

// SupportedURISchemes returns a list of playable URI formats. See
// https://github.com/popcornmix/omxplayer#supportedurischemes for more details.
func (p *Player) SupportedURISchemes() ([]string, error) {
	return dbusGetStringArray(p.bus, propSupportedURISchemes)
}

// SupportedMimeTypes returns a list of supported MIME types. See
// https://github.com/popcornmix/omxplayer#supportedmimetypes for more details.
func (p *Player) SupportedMimeTypes() ([]string, error) {
	return dbusGetStringArray(p.bus, propSupportedMimeTypes)
}

// CanGoNext returns true if the player can skip to the next track, false
// otherwise. See https://github.com/popcornmix/omxplayer#cangonext for more
// details.
func (p *Player) CanGoNext() (bool, error) {
	return dbusGetBool(p.bus, propCanGoNext)
}

// CanGoPrevious returns true if the player can skip to previous track, false
// otherwise. See https://github.com/popcornmix/omxplayer#cangoprevious for more
// details.
func (p *Player) CanGoPrevious() (bool, error) {
	return dbusGetBool(p.bus, propCanGoPrevious)
}

// CanSeek returns true if the player can seek, false otherwise. See
// https://github.com/popcornmix/omxplayer#canseek for more details.
func (p *Player) CanSeek() (bool, error) {
	return dbusGetBool(p.bus, cmdSeek)
}

// CanControl returns true if the player can be controlled, false otherwise. See
// https://github.com/popcornmix/omxplayer#cancontrol for more details.
func (p *Player) CanControl() (bool, error) {
	return dbusGetBool(p.bus, propCanControl)
}

// CanPlay returns true if the player can play, false otherwise. See
// https://github.com/popcornmix/omxplayer#canplay for more details.
func (p *Player) CanPlay() (bool, error) {
	return dbusGetBool(p.bus, propCanPlay)
}

// CanPause returns true if the player can pause, false otherwise. See
// https://github.com/popcornmix/omxplayer#canpause for more details.
func (p *Player) CanPause() (bool, error) {
	return dbusGetBool(p.bus, propCanPause)
}

// Next tells the player to skip to the next chapter. See
// https://github.com/popcornmix/omxplayer#next for more details.
func (p *Player) Next() error {
	return dbusCall(p.bus, cmdNext)
}

// Previous tells the player to skip to the previous chapter. See
// https://github.com/popcornmix/omxplayer#previous for more details.
func (p *Player) Previous() error {
	return dbusCall(p.bus, cmdPrevious)
}

// Pause pauses the player if it is playing. Otherwise, it resumes playback. See
// https://github.com/popcornmix/omxplayer#pause for more details.
func (p *Player) Pause() error {
	return dbusCall(p.bus, cmdPause)
}

// PlayPause pauses the player if it is playing. Otherwise, it resumes playback.
// See https://github.com/popcornmix/omxplayer#playpause for more details.
func (p *Player) PlayPause() error {
	return dbusCall(p.bus, cmdPlayPause)
}

// Stop tells the player to stop playing the video. See
// https://github.com/popcornmix/omxplayer#stop for more details.
func (p *Player) Stop() error {
	return dbusCall(p.bus, cmdStop)
}

// Seek performs a relative seek from the current video position. See
// https://github.com/popcornmix/omxplayer#seek for more details.
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

// SetPosition performs an absolute seek to the specified video position. See
// https://github.com/popcornmix/omxplayer#setposition for more details.
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

// PlaybackStatus returns the current state of the player. See
// https://github.com/popcornmix/omxplayer#playbackstatus for more details.
func (p *Player) PlaybackStatus() (string, error) {
	return dbusGetString(p.bus, propPlaybackStatus)
}

// Volume returns the current volume. Sets a new volume when an argument is
// specified. See https://github.com/popcornmix/omxplayer#volume for more
// details.
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

// Mute mutes the video's audio stream. See
// https://github.com/popcornmix/omxplayer#mute for more details.
func (p *Player) Mute() error {
	return dbusCall(p.bus, cmdMute)
}

// Unmute unmutes the video's audio stream. See
// https://github.com/popcornmix/omxplayer#unmute for more details.
func (p *Player) Unmute() error {
	return dbusCall(p.bus, cmdUnmute)
}

// Position returns the current position in the video in milliseconds. See
// https://github.com/popcornmix/omxplayer#position for more details.
func (p *Player) Position() (int64, error) {
	return dbusGetInt64(p.bus, propPosition)
}

// Aspect returns the aspect ratio. See
// https://github.com/popcornmix/omxplayer/blob/master/OMXControl.cpp#L362.
func (p *Player) Aspect() (float64, error) {
	return dbusGetFloat64(p.bus, propAspect)
}

// VideoStreamCount returns the number of available video streams. See
// https://github.com/popcornmix/omxplayer/blob/master/OMXControl.cpp#L369.
func (p *Player) VideoStreamCount() (int64, error) {
	return dbusGetInt64(p.bus, propVideoStreamCount)
}

// ResWidth returns the width of the video. See
// https://github.com/popcornmix/omxplayer/blob/master/OMXControl.cpp#L376.
func (p *Player) ResWidth() (int64, error) {
	return dbusGetInt64(p.bus, propResWidth)
}

// ResHeight returns the height of the video. See
// https://github.com/popcornmix/omxplayer/blob/master/OMXControl.cpp#L383.
func (p *Player) ResHeight() (int64, error) {
	return dbusGetInt64(p.bus, propResHeight)
}

// Duration returns the total length of the video in milliseconds. See
// https://github.com/popcornmix/omxplayer#duration for more details.
func (p *Player) Duration() (int64, error) {
	return dbusGetInt64(p.bus, propDuration)
}

// MinimumRate returns the minimum playback rate. See
// https://github.com/popcornmix/omxplayer#minimumrate for more details.
func (p *Player) MinimumRate() (float64, error) {
	return dbusGetFloat64(p.bus, propMinimumRate)
}

// MaximumRate returns the maximum playback rate. See
// https://github.com/popcornmix/omxplayer#maximumrate for more details.
func (p *Player) MaximumRate() (float64, error) {
	return dbusGetFloat64(p.bus, propMaximumRate)
}

// ListSubtitles returns a list of the subtitles available in the video file.
// See https://github.com/popcornmix/omxplayer#listsubtitles for more details.
func (p *Player) ListSubtitles() ([]string, error) {
	return dbusGetStringArray(p.bus, cmdListSubtitles)
}

// HideVideo is an undocumented D-Bus method. See
// https://github.com/popcornmix/omxplayer/blob/master/OMXControl.cpp#L457.
func (p *Player) HideVideo() error {
	return dbusCall(p.bus, cmdHideVideo)
}

// UnHideVideo is an undocumented D-Bus method. See
// https://github.com/popcornmix/omxplayer/blob/master/OMXControl.cpp#L462.
func (p *Player) UnHideVideo() error {
	return dbusCall(p.bus, cmdUnHideVideo)
}

// ListAudio returns a list of the audio tracks available in the video file. See
// https://github.com/popcornmix/omxplayer#listaudio for more details.
func (p *Player) ListAudio() ([]string, error) {
	return dbusGetStringArray(p.bus, cmdListAudio)
}

// ListVideo returns a list of the video tracks available in the video file. See
// https://github.com/popcornmix/omxplayer#listvideo for more details.
func (p *Player) ListVideo() ([]string, error) {
	return dbusGetStringArray(p.bus, cmdListVideo)
}

// SelectSubtitle specifies which subtitle track should be used. See
// https://github.com/popcornmix/omxplayer#selectsubtitle for more details.
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

// SelectAudio specifies which audio track should be used. See
// https://github.com/popcornmix/omxplayer#selectaudio for more details.
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

// ShowSubtitles starts displaying subtitles. See
// https://github.com/popcornmix/omxplayer#showsubtitles for more details.
func (p *Player) ShowSubtitles() error {
	return dbusCall(p.bus, cmdShowSubtitles)
}

// HideSubtitles stops displaying subtitles. See
// https://github.com/popcornmix/omxplayer#hidesubtitles for more details.
func (p *Player) HideSubtitles() error {
	return dbusCall(p.bus, cmdHideSubtitles)
}

// Action allows for executing keyboard commands. See
// https://github.com/popcornmix/omxplayer#action for more details.
func (p *Player) Action(action int32) error {
	log.WithFields(log.Fields{
		"path":        cmdAction,
		"paramAction": action,
	}).Debug("omxplayer: dbus call")
	return p.bus.Call(cmdAction, 0, action).Err
}
