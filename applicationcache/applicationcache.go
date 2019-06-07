// Package applicationcache provides the Chrome DevTools Protocol
// commands, types, and events for the ApplicationCache domain.
//
// Generated by the cdproto-gen command.
package applicationcache

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/seddonm1/cdproto/cdp"
)

// EnableParams enables application cache domain notifications.
type EnableParams struct{}

// Enable enables application cache domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes ApplicationCache.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// GetApplicationCacheForFrameParams returns relevant application cache data
// for the document in given frame.
type GetApplicationCacheForFrameParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Identifier of the frame containing document whose application cache is retrieved.
}

// GetApplicationCacheForFrame returns relevant application cache data for
// the document in given frame.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache#method-getApplicationCacheForFrame
//
// parameters:
//   frameID - Identifier of the frame containing document whose application cache is retrieved.
func GetApplicationCacheForFrame(frameID cdp.FrameID) *GetApplicationCacheForFrameParams {
	return &GetApplicationCacheForFrameParams{
		FrameID: frameID,
	}
}

// GetApplicationCacheForFrameReturns return values.
type GetApplicationCacheForFrameReturns struct {
	ApplicationCache *ApplicationCache `json:"applicationCache,omitempty"` // Relevant application cache data for the document in given frame.
}

// Do executes ApplicationCache.getApplicationCacheForFrame against the provided context.
//
// returns:
//   applicationCache - Relevant application cache data for the document in given frame.
func (p *GetApplicationCacheForFrameParams) Do(ctx context.Context) (applicationCache *ApplicationCache, err error) {
	// execute
	var res GetApplicationCacheForFrameReturns
	err = cdp.Execute(ctx, CommandGetApplicationCacheForFrame, p, &res)
	if err != nil {
		return nil, err
	}

	return res.ApplicationCache, nil
}

// GetFramesWithManifestsParams returns array of frame identifiers with
// manifest urls for each frame containing a document associated with some
// application cache.
type GetFramesWithManifestsParams struct{}

// GetFramesWithManifests returns array of frame identifiers with manifest
// urls for each frame containing a document associated with some application
// cache.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache#method-getFramesWithManifests
func GetFramesWithManifests() *GetFramesWithManifestsParams {
	return &GetFramesWithManifestsParams{}
}

// GetFramesWithManifestsReturns return values.
type GetFramesWithManifestsReturns struct {
	FrameIds []*FrameWithManifest `json:"frameIds,omitempty"` // Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
}

// Do executes ApplicationCache.getFramesWithManifests against the provided context.
//
// returns:
//   frameIds - Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (p *GetFramesWithManifestsParams) Do(ctx context.Context) (frameIds []*FrameWithManifest, err error) {
	// execute
	var res GetFramesWithManifestsReturns
	err = cdp.Execute(ctx, CommandGetFramesWithManifests, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.FrameIds, nil
}

// GetManifestForFrameParams returns manifest URL for document in the given
// frame.
type GetManifestForFrameParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Identifier of the frame containing document whose manifest is retrieved.
}

// GetManifestForFrame returns manifest URL for document in the given frame.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache#method-getManifestForFrame
//
// parameters:
//   frameID - Identifier of the frame containing document whose manifest is retrieved.
func GetManifestForFrame(frameID cdp.FrameID) *GetManifestForFrameParams {
	return &GetManifestForFrameParams{
		FrameID: frameID,
	}
}

// GetManifestForFrameReturns return values.
type GetManifestForFrameReturns struct {
	ManifestURL string `json:"manifestURL,omitempty"` // Manifest URL for document in the given frame.
}

// Do executes ApplicationCache.getManifestForFrame against the provided context.
//
// returns:
//   manifestURL - Manifest URL for document in the given frame.
func (p *GetManifestForFrameParams) Do(ctx context.Context) (manifestURL string, err error) {
	// execute
	var res GetManifestForFrameReturns
	err = cdp.Execute(ctx, CommandGetManifestForFrame, p, &res)
	if err != nil {
		return "", err
	}

	return res.ManifestURL, nil
}

// Command names.
const (
	CommandEnable                      = "ApplicationCache.enable"
	CommandGetApplicationCacheForFrame = "ApplicationCache.getApplicationCacheForFrame"
	CommandGetFramesWithManifests      = "ApplicationCache.getFramesWithManifests"
	CommandGetManifestForFrame         = "ApplicationCache.getManifestForFrame"
)
