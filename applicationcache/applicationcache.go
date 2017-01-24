// Package applicationcache provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome ApplicationCache domain.
//
// Generated by the chromedp-gen command.
package applicationcache

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	. "github.com/knq/chromedp/cdp"
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

// GetFramesWithManifestsParams returns array of frame identifiers with
// manifest urls for each frame containing a document associated with some
// application cache.
type GetFramesWithManifestsParams struct{}

// GetFramesWithManifests returns array of frame identifiers with manifest
// urls for each frame containing a document associated with some application
// cache.
func GetFramesWithManifests() *GetFramesWithManifestsParams {
	return &GetFramesWithManifestsParams{}
}

// GetFramesWithManifestsReturns return values.
type GetFramesWithManifestsReturns struct {
	FrameIds []*FrameWithManifest `json:"frameIds,omitempty"` // Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
}

// Do executes ApplicationCache.getFramesWithManifests.
//
// returns:
//   frameIds - Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (p *GetFramesWithManifestsParams) Do(ctxt context.Context, h FrameHandler) (frameIds []*FrameWithManifest, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandApplicationCacheGetFramesWithManifests, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetFramesWithManifestsReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, ErrInvalidResult
			}

			return r.FrameIds, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, ErrContextDone
	}

	return nil, ErrUnknownResult
}

// EnableParams enables application cache domain notifications.
type EnableParams struct{}

// Enable enables application cache domain notifications.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes ApplicationCache.enable.
func (p *EnableParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandApplicationCacheEnable, Empty)

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

// GetManifestForFrameParams returns manifest URL for document in the given
// frame.
type GetManifestForFrameParams struct {
	FrameID FrameID `json:"frameId"` // Identifier of the frame containing document whose manifest is retrieved.
}

// GetManifestForFrame returns manifest URL for document in the given frame.
//
// parameters:
//   frameId - Identifier of the frame containing document whose manifest is retrieved.
func GetManifestForFrame(frameId FrameID) *GetManifestForFrameParams {
	return &GetManifestForFrameParams{
		FrameID: frameId,
	}
}

// GetManifestForFrameReturns return values.
type GetManifestForFrameReturns struct {
	ManifestURL string `json:"manifestURL,omitempty"` // Manifest URL for document in the given frame.
}

// Do executes ApplicationCache.getManifestForFrame.
//
// returns:
//   manifestURL - Manifest URL for document in the given frame.
func (p *GetManifestForFrameParams) Do(ctxt context.Context, h FrameHandler) (manifestURL string, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return "", err
	}

	// execute
	ch := h.Execute(ctxt, CommandApplicationCacheGetManifestForFrame, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return "", ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetManifestForFrameReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return "", ErrInvalidResult
			}

			return r.ManifestURL, nil

		case error:
			return "", v
		}

	case <-ctxt.Done():
		return "", ErrContextDone
	}

	return "", ErrUnknownResult
}

// GetApplicationCacheForFrameParams returns relevant application cache data
// for the document in given frame.
type GetApplicationCacheForFrameParams struct {
	FrameID FrameID `json:"frameId"` // Identifier of the frame containing document whose application cache is retrieved.
}

// GetApplicationCacheForFrame returns relevant application cache data for
// the document in given frame.
//
// parameters:
//   frameId - Identifier of the frame containing document whose application cache is retrieved.
func GetApplicationCacheForFrame(frameId FrameID) *GetApplicationCacheForFrameParams {
	return &GetApplicationCacheForFrameParams{
		FrameID: frameId,
	}
}

// GetApplicationCacheForFrameReturns return values.
type GetApplicationCacheForFrameReturns struct {
	ApplicationCache *ApplicationCache `json:"applicationCache,omitempty"` // Relevant application cache data for the document in given frame.
}

// Do executes ApplicationCache.getApplicationCacheForFrame.
//
// returns:
//   applicationCache - Relevant application cache data for the document in given frame.
func (p *GetApplicationCacheForFrameParams) Do(ctxt context.Context, h FrameHandler) (applicationCache *ApplicationCache, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandApplicationCacheGetApplicationCacheForFrame, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetApplicationCacheForFrameReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, ErrInvalidResult
			}

			return r.ApplicationCache, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, ErrContextDone
	}

	return nil, ErrUnknownResult
}
