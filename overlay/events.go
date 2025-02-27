package overlay

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/seddonm1/cdproto/cdp"
	"github.com/seddonm1/cdproto/page"
)

// EventInspectNodeRequested fired when the node should be inspected. This
// happens after call to setInspectMode or when user manually inspects an
// element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Overlay#event-inspectNodeRequested
type EventInspectNodeRequested struct {
	BackendNodeID cdp.BackendNodeID `json:"backendNodeId"` // Id of the node to inspect.
}

// EventNodeHighlightRequested fired when the node should be highlighted.
// This happens after call to setInspectMode.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Overlay#event-nodeHighlightRequested
type EventNodeHighlightRequested struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// EventScreenshotRequested fired when user asks to capture screenshot of
// some area on the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Overlay#event-screenshotRequested
type EventScreenshotRequested struct {
	Viewport *page.Viewport `json:"viewport"` // Viewport to capture, in device independent pixels (dip).
}

// EventInspectModeCanceled fired when user cancels the inspect mode.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Overlay#event-inspectModeCanceled
type EventInspectModeCanceled struct{}
