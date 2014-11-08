package omxplayer

import (
	"fmt"
)

type Player struct {
}

func (p *Player) Quit() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanQuit() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Fullscreen() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanSetFullscreen() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanRaise() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) HasTrackList() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Identity() (string, error) {
	return "", fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) SupportedUriSchemes() ([]string, error) {
	return nil, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) SupportedMimeTypes() ([]string, error) {
	return nil, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanGoNext() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanGoPrevious() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanSeek() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanControl() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanPlay() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) CanPause() (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Next() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Previous() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Pause() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) PlayPause() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Stop() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Seek(amount int64) (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) SetPosition(path string, position int64) (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) PlaybackStatus() (string, error) {
	return "", fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Volume(volume ...float64) (float64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Mute() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Unmute() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Position() (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Aspect() (float64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) VideoStreamCount() (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) ResWidth() (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) ResHeight() (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Duration() (int64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) MinimumRate() (float64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) MaximumRate() (float64, error) {
	return 0, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) ListSubtitles() ([]string, error) {
	return nil, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) HideVideo() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) UnHideVideo() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) ListAudio() ([]string, error) {
	return nil, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) ListVideo() ([]string, error) {
	return nil, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) SelectSubtitle(index int32) (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) SelectAudio(index int32) (bool, error) {
	return false, fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) ShowSubtitles() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) HideSubtitles() error {
	return fmt.Errorf("omxplayer: not implemented yet")
}

func (p *Player) Action(action int32) error {
	return fmt.Errorf("omxplayer: not implemented yet")
}
