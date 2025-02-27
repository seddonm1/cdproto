// Package security provides the Chrome DevTools Protocol
// commands, types, and events for the Security domain.
//
// Security.
//
// Generated by the cdproto-gen command.
package security

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/seddonm1/cdproto/cdp"
)

// DisableParams disables tracking security state changes.
type DisableParams struct{}

// Disable disables tracking security state changes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Security.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enables tracking security state changes.
type EnableParams struct{}

// Enable enables tracking security state changes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Security.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// SetIgnoreCertificateErrorsParams enable/disable whether all certificate
// errors should be ignored.
type SetIgnoreCertificateErrorsParams struct {
	Ignore bool `json:"ignore"` // If true, all certificate errors will be ignored.
}

// SetIgnoreCertificateErrors enable/disable whether all certificate errors
// should be ignored.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Security#method-setIgnoreCertificateErrors
//
// parameters:
//   ignore - If true, all certificate errors will be ignored.
func SetIgnoreCertificateErrors(ignore bool) *SetIgnoreCertificateErrorsParams {
	return &SetIgnoreCertificateErrorsParams{
		Ignore: ignore,
	}
}

// Do executes Security.setIgnoreCertificateErrors against the provided context.
func (p *SetIgnoreCertificateErrorsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetIgnoreCertificateErrors, p, nil)
}

// Command names.
const (
	CommandDisable                    = "Security.disable"
	CommandEnable                     = "Security.enable"
	CommandSetIgnoreCertificateErrors = "Security.setIgnoreCertificateErrors"
)
