// Package animation provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Animation domain.
//
// Generated by the chromedp-gen command.
package animation

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	. "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/runtime"
	"github.com/mailru/easyjson"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

// EnableParams enables animation domain notifications.
type EnableParams struct{}

// Enable enables animation domain notifications.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Animation.enable.
func (p *EnableParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationEnable, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// DisableParams disables animation domain notifications.
type DisableParams struct{}

// Disable disables animation domain notifications.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Animation.disable.
func (p *DisableParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationDisable, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// GetPlaybackRateParams gets the playback rate of the document timeline.
type GetPlaybackRateParams struct{}

// GetPlaybackRate gets the playback rate of the document timeline.
func GetPlaybackRate() *GetPlaybackRateParams {
	return &GetPlaybackRateParams{}
}

// GetPlaybackRateReturns return values.
type GetPlaybackRateReturns struct {
	PlaybackRate float64 `json:"playbackRate,omitempty"` // Playback rate for animations on page.
}

// Do executes Animation.getPlaybackRate.
//
// returns:
//   playbackRate - Playback rate for animations on page.
func (p *GetPlaybackRateParams) Do(ctxt context.Context, h FrameHandler) (playbackRate float64, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationGetPlaybackRate, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return 0, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetPlaybackRateReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return 0, ErrInvalidResult
			}

			return r.PlaybackRate, nil

		case error:
			return 0, v
		}

	case <-ctxt.Done():
		return 0, ErrContextDone
	}

	return 0, ErrUnknownResult
}

// SetPlaybackRateParams sets the playback rate of the document timeline.
type SetPlaybackRateParams struct {
	PlaybackRate float64 `json:"playbackRate"` // Playback rate for animations on page
}

// SetPlaybackRate sets the playback rate of the document timeline.
//
// parameters:
//   playbackRate - Playback rate for animations on page
func SetPlaybackRate(playbackRate float64) *SetPlaybackRateParams {
	return &SetPlaybackRateParams{
		PlaybackRate: playbackRate,
	}
}

// Do executes Animation.setPlaybackRate.
func (p *SetPlaybackRateParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationSetPlaybackRate, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// GetCurrentTimeParams returns the current time of the an animation.
type GetCurrentTimeParams struct {
	ID string `json:"id"` // Id of animation.
}

// GetCurrentTime returns the current time of the an animation.
//
// parameters:
//   id - Id of animation.
func GetCurrentTime(id string) *GetCurrentTimeParams {
	return &GetCurrentTimeParams{
		ID: id,
	}
}

// GetCurrentTimeReturns return values.
type GetCurrentTimeReturns struct {
	CurrentTime float64 `json:"currentTime,omitempty"` // Current time of the page.
}

// Do executes Animation.getCurrentTime.
//
// returns:
//   currentTime - Current time of the page.
func (p *GetCurrentTimeParams) Do(ctxt context.Context, h FrameHandler) (currentTime float64, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return 0, err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationGetCurrentTime, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return 0, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetCurrentTimeReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return 0, ErrInvalidResult
			}

			return r.CurrentTime, nil

		case error:
			return 0, v
		}

	case <-ctxt.Done():
		return 0, ErrContextDone
	}

	return 0, ErrUnknownResult
}

// SetPausedParams sets the paused state of a set of animations.
type SetPausedParams struct {
	Animations []string `json:"animations"` // Animations to set the pause state of.
	Paused     bool     `json:"paused"`     // Paused state to set to.
}

// SetPaused sets the paused state of a set of animations.
//
// parameters:
//   animations - Animations to set the pause state of.
//   paused - Paused state to set to.
func SetPaused(animations []string, paused bool) *SetPausedParams {
	return &SetPausedParams{
		Animations: animations,
		Paused:     paused,
	}
}

// Do executes Animation.setPaused.
func (p *SetPausedParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationSetPaused, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// SetTimingParams sets the timing of an animation node.
type SetTimingParams struct {
	AnimationID string  `json:"animationId"` // Animation id.
	Duration    float64 `json:"duration"`    // Duration of the animation.
	Delay       float64 `json:"delay"`       // Delay of the animation.
}

// SetTiming sets the timing of an animation node.
//
// parameters:
//   animationId - Animation id.
//   duration - Duration of the animation.
//   delay - Delay of the animation.
func SetTiming(animationId string, duration float64, delay float64) *SetTimingParams {
	return &SetTimingParams{
		AnimationID: animationId,
		Duration:    duration,
		Delay:       delay,
	}
}

// Do executes Animation.setTiming.
func (p *SetTimingParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationSetTiming, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// SeekAnimationsParams seek a set of animations to a particular time within
// each animation.
type SeekAnimationsParams struct {
	Animations  []string `json:"animations"`  // List of animation ids to seek.
	CurrentTime float64  `json:"currentTime"` // Set the current time of each animation.
}

// SeekAnimations seek a set of animations to a particular time within each
// animation.
//
// parameters:
//   animations - List of animation ids to seek.
//   currentTime - Set the current time of each animation.
func SeekAnimations(animations []string, currentTime float64) *SeekAnimationsParams {
	return &SeekAnimationsParams{
		Animations:  animations,
		CurrentTime: currentTime,
	}
}

// Do executes Animation.seekAnimations.
func (p *SeekAnimationsParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationSeekAnimations, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// ReleaseAnimationsParams releases a set of animations to no longer be
// manipulated.
type ReleaseAnimationsParams struct {
	Animations []string `json:"animations"` // List of animation ids to seek.
}

// ReleaseAnimations releases a set of animations to no longer be
// manipulated.
//
// parameters:
//   animations - List of animation ids to seek.
func ReleaseAnimations(animations []string) *ReleaseAnimationsParams {
	return &ReleaseAnimationsParams{
		Animations: animations,
	}
}

// Do executes Animation.releaseAnimations.
func (p *ReleaseAnimationsParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationReleaseAnimations, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// ResolveAnimationParams gets the remote object of the Animation.
type ResolveAnimationParams struct {
	AnimationID string `json:"animationId"` // Animation id.
}

// ResolveAnimation gets the remote object of the Animation.
//
// parameters:
//   animationId - Animation id.
func ResolveAnimation(animationId string) *ResolveAnimationParams {
	return &ResolveAnimationParams{
		AnimationID: animationId,
	}
}

// ResolveAnimationReturns return values.
type ResolveAnimationReturns struct {
	RemoteObject *runtime.RemoteObject `json:"remoteObject,omitempty"` // Corresponding remote object.
}

// Do executes Animation.resolveAnimation.
//
// returns:
//   remoteObject - Corresponding remote object.
func (p *ResolveAnimationParams) Do(ctxt context.Context, h FrameHandler) (remoteObject *runtime.RemoteObject, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandAnimationResolveAnimation, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r ResolveAnimationReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, ErrInvalidResult
			}

			return r.RemoteObject, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, ErrContextDone
	}

	return nil, ErrUnknownResult
}
