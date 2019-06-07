// Package browser provides the Chrome DevTools Protocol
// commands, types, and events for the Browser domain.
//
// The Browser domain defines methods and events for browser managing.
//
// Generated by the cdproto-gen command.
package browser

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/seddonm1/cdproto/cdp"
	"github.com/seddonm1/cdproto/target"
)

// GrantPermissionsParams grant specific permissions to the given origin and
// reject all others.
type GrantPermissionsParams struct {
	Origin           string                  `json:"origin"`
	Permissions      []PermissionType        `json:"permissions"`
	BrowserContextID target.BrowserContextID `json:"browserContextId,omitempty"` // BrowserContext to override permissions. When omitted, default browser context is used.
}

// GrantPermissions grant specific permissions to the given origin and reject
// all others.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-grantPermissions
//
// parameters:
//   origin
//   permissions
func GrantPermissions(origin string, permissions []PermissionType) *GrantPermissionsParams {
	return &GrantPermissionsParams{
		Origin:      origin,
		Permissions: permissions,
	}
}

// WithBrowserContextID browserContext to override permissions. When omitted,
// default browser context is used.
func (p GrantPermissionsParams) WithBrowserContextID(browserContextID target.BrowserContextID) *GrantPermissionsParams {
	p.BrowserContextID = browserContextID
	return &p
}

// Do executes Browser.grantPermissions against the provided context.
func (p *GrantPermissionsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandGrantPermissions, p, nil)
}

// ResetPermissionsParams reset all permission management for all origins.
type ResetPermissionsParams struct {
	BrowserContextID target.BrowserContextID `json:"browserContextId,omitempty"` // BrowserContext to reset permissions. When omitted, default browser context is used.
}

// ResetPermissions reset all permission management for all origins.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-resetPermissions
//
// parameters:
func ResetPermissions() *ResetPermissionsParams {
	return &ResetPermissionsParams{}
}

// WithBrowserContextID browserContext to reset permissions. When omitted,
// default browser context is used.
func (p ResetPermissionsParams) WithBrowserContextID(browserContextID target.BrowserContextID) *ResetPermissionsParams {
	p.BrowserContextID = browserContextID
	return &p
}

// Do executes Browser.resetPermissions against the provided context.
func (p *ResetPermissionsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandResetPermissions, p, nil)
}

// CloseParams close browser gracefully.
type CloseParams struct{}

// Close close browser gracefully.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-close
func Close() *CloseParams {
	return &CloseParams{}
}

// Do executes Browser.close against the provided context.
func (p *CloseParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClose, nil, nil)
}

// CrashParams crashes browser on the main thread.
type CrashParams struct{}

// Crash crashes browser on the main thread.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-crash
func Crash() *CrashParams {
	return &CrashParams{}
}

// Do executes Browser.crash against the provided context.
func (p *CrashParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandCrash, nil, nil)
}

// CrashGpuProcessParams crashes GPU process.
type CrashGpuProcessParams struct{}

// CrashGpuProcess crashes GPU process.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-crashGpuProcess
func CrashGpuProcess() *CrashGpuProcessParams {
	return &CrashGpuProcessParams{}
}

// Do executes Browser.crashGpuProcess against the provided context.
func (p *CrashGpuProcessParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandCrashGpuProcess, nil, nil)
}

// GetVersionParams returns version information.
type GetVersionParams struct{}

// GetVersion returns version information.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-getVersion
func GetVersion() *GetVersionParams {
	return &GetVersionParams{}
}

// GetVersionReturns return values.
type GetVersionReturns struct {
	ProtocolVersion string `json:"protocolVersion,omitempty"` // Protocol version.
	Product         string `json:"product,omitempty"`         // Product name.
	Revision        string `json:"revision,omitempty"`        // Product revision.
	UserAgent       string `json:"userAgent,omitempty"`       // User-Agent.
	JsVersion       string `json:"jsVersion,omitempty"`       // V8 version.
}

// Do executes Browser.getVersion against the provided context.
//
// returns:
//   protocolVersion - Protocol version.
//   product - Product name.
//   revision - Product revision.
//   userAgent - User-Agent.
//   jsVersion - V8 version.
func (p *GetVersionParams) Do(ctx context.Context) (protocolVersion string, product string, revision string, userAgent string, jsVersion string, err error) {
	// execute
	var res GetVersionReturns
	err = cdp.Execute(ctx, CommandGetVersion, nil, &res)
	if err != nil {
		return "", "", "", "", "", err
	}

	return res.ProtocolVersion, res.Product, res.Revision, res.UserAgent, res.JsVersion, nil
}

// GetBrowserCommandLineParams returns the command line switches for the
// browser process if, and only if --enable-automation is on the commandline.
type GetBrowserCommandLineParams struct{}

// GetBrowserCommandLine returns the command line switches for the browser
// process if, and only if --enable-automation is on the commandline.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-getBrowserCommandLine
func GetBrowserCommandLine() *GetBrowserCommandLineParams {
	return &GetBrowserCommandLineParams{}
}

// GetBrowserCommandLineReturns return values.
type GetBrowserCommandLineReturns struct {
	Arguments []string `json:"arguments,omitempty"` // Commandline parameters
}

// Do executes Browser.getBrowserCommandLine against the provided context.
//
// returns:
//   arguments - Commandline parameters
func (p *GetBrowserCommandLineParams) Do(ctx context.Context) (arguments []string, err error) {
	// execute
	var res GetBrowserCommandLineReturns
	err = cdp.Execute(ctx, CommandGetBrowserCommandLine, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Arguments, nil
}

// GetHistogramsParams get Chrome histograms.
type GetHistogramsParams struct {
	Query string `json:"query,omitempty"` // Requested substring in name. Only histograms which have query as a substring in their name are extracted. An empty or absent query returns all histograms.
	Delta bool   `json:"delta,omitempty"` // If true, retrieve delta since last call.
}

// GetHistograms get Chrome histograms.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-getHistograms
//
// parameters:
func GetHistograms() *GetHistogramsParams {
	return &GetHistogramsParams{}
}

// WithQuery requested substring in name. Only histograms which have query as
// a substring in their name are extracted. An empty or absent query returns all
// histograms.
func (p GetHistogramsParams) WithQuery(query string) *GetHistogramsParams {
	p.Query = query
	return &p
}

// WithDelta if true, retrieve delta since last call.
func (p GetHistogramsParams) WithDelta(delta bool) *GetHistogramsParams {
	p.Delta = delta
	return &p
}

// GetHistogramsReturns return values.
type GetHistogramsReturns struct {
	Histograms []*Histogram `json:"histograms,omitempty"` // Histograms.
}

// Do executes Browser.getHistograms against the provided context.
//
// returns:
//   histograms - Histograms.
func (p *GetHistogramsParams) Do(ctx context.Context) (histograms []*Histogram, err error) {
	// execute
	var res GetHistogramsReturns
	err = cdp.Execute(ctx, CommandGetHistograms, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Histograms, nil
}

// GetHistogramParams get a Chrome histogram by name.
type GetHistogramParams struct {
	Name  string `json:"name"`            // Requested histogram name.
	Delta bool   `json:"delta,omitempty"` // If true, retrieve delta since last call.
}

// GetHistogram get a Chrome histogram by name.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-getHistogram
//
// parameters:
//   name - Requested histogram name.
func GetHistogram(name string) *GetHistogramParams {
	return &GetHistogramParams{
		Name: name,
	}
}

// WithDelta if true, retrieve delta since last call.
func (p GetHistogramParams) WithDelta(delta bool) *GetHistogramParams {
	p.Delta = delta
	return &p
}

// GetHistogramReturns return values.
type GetHistogramReturns struct {
	Histogram *Histogram `json:"histogram,omitempty"` // Histogram.
}

// Do executes Browser.getHistogram against the provided context.
//
// returns:
//   histogram - Histogram.
func (p *GetHistogramParams) Do(ctx context.Context) (histogram *Histogram, err error) {
	// execute
	var res GetHistogramReturns
	err = cdp.Execute(ctx, CommandGetHistogram, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Histogram, nil
}

// GetWindowBoundsParams get position and size of the browser window.
type GetWindowBoundsParams struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
}

// GetWindowBounds get position and size of the browser window.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-getWindowBounds
//
// parameters:
//   windowID - Browser window id.
func GetWindowBounds(windowID WindowID) *GetWindowBoundsParams {
	return &GetWindowBoundsParams{
		WindowID: windowID,
	}
}

// GetWindowBoundsReturns return values.
type GetWindowBoundsReturns struct {
	Bounds *Bounds `json:"bounds,omitempty"` // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// Do executes Browser.getWindowBounds against the provided context.
//
// returns:
//   bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (p *GetWindowBoundsParams) Do(ctx context.Context) (bounds *Bounds, err error) {
	// execute
	var res GetWindowBoundsReturns
	err = cdp.Execute(ctx, CommandGetWindowBounds, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Bounds, nil
}

// GetWindowForTargetParams get the browser window that contains the devtools
// target.
type GetWindowForTargetParams struct {
	TargetID target.ID `json:"targetId,omitempty"` // Devtools agent host id. If called as a part of the session, associated targetId is used.
}

// GetWindowForTarget get the browser window that contains the devtools
// target.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-getWindowForTarget
//
// parameters:
func GetWindowForTarget() *GetWindowForTargetParams {
	return &GetWindowForTargetParams{}
}

// WithTargetID devtools agent host id. If called as a part of the session,
// associated targetId is used.
func (p GetWindowForTargetParams) WithTargetID(targetID target.ID) *GetWindowForTargetParams {
	p.TargetID = targetID
	return &p
}

// GetWindowForTargetReturns return values.
type GetWindowForTargetReturns struct {
	WindowID WindowID `json:"windowId,omitempty"` // Browser window id.
	Bounds   *Bounds  `json:"bounds,omitempty"`   // Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}

// Do executes Browser.getWindowForTarget against the provided context.
//
// returns:
//   windowID - Browser window id.
//   bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (p *GetWindowForTargetParams) Do(ctx context.Context) (windowID WindowID, bounds *Bounds, err error) {
	// execute
	var res GetWindowForTargetReturns
	err = cdp.Execute(ctx, CommandGetWindowForTarget, p, &res)
	if err != nil {
		return 0, nil, err
	}

	return res.WindowID, res.Bounds, nil
}

// SetWindowBoundsParams set position and/or size of the browser window.
type SetWindowBoundsParams struct {
	WindowID WindowID `json:"windowId"` // Browser window id.
	Bounds   *Bounds  `json:"bounds"`   // New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
}

// SetWindowBounds set position and/or size of the browser window.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-setWindowBounds
//
// parameters:
//   windowID - Browser window id.
//   bounds - New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
func SetWindowBounds(windowID WindowID, bounds *Bounds) *SetWindowBoundsParams {
	return &SetWindowBoundsParams{
		WindowID: windowID,
		Bounds:   bounds,
	}
}

// Do executes Browser.setWindowBounds against the provided context.
func (p *SetWindowBoundsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetWindowBounds, p, nil)
}

// SetDockTileParams set dock tile details, platform-specific.
type SetDockTileParams struct {
	BadgeLabel string `json:"badgeLabel,omitempty"`
	Image      string `json:"image,omitempty"` // Png encoded image.
}

// SetDockTile set dock tile details, platform-specific.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Browser#method-setDockTile
//
// parameters:
func SetDockTile() *SetDockTileParams {
	return &SetDockTileParams{}
}

// WithBadgeLabel [no description].
func (p SetDockTileParams) WithBadgeLabel(badgeLabel string) *SetDockTileParams {
	p.BadgeLabel = badgeLabel
	return &p
}

// WithImage png encoded image.
func (p SetDockTileParams) WithImage(image string) *SetDockTileParams {
	p.Image = image
	return &p
}

// Do executes Browser.setDockTile against the provided context.
func (p *SetDockTileParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetDockTile, p, nil)
}

// Command names.
const (
	CommandGrantPermissions      = "Browser.grantPermissions"
	CommandResetPermissions      = "Browser.resetPermissions"
	CommandClose                 = "Browser.close"
	CommandCrash                 = "Browser.crash"
	CommandCrashGpuProcess       = "Browser.crashGpuProcess"
	CommandGetVersion            = "Browser.getVersion"
	CommandGetBrowserCommandLine = "Browser.getBrowserCommandLine"
	CommandGetHistograms         = "Browser.getHistograms"
	CommandGetHistogram          = "Browser.getHistogram"
	CommandGetWindowBounds       = "Browser.getWindowBounds"
	CommandGetWindowForTarget    = "Browser.getWindowForTarget"
	CommandSetWindowBounds       = "Browser.setWindowBounds"
	CommandSetDockTile           = "Browser.setDockTile"
)
