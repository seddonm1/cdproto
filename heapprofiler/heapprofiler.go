// Package heapprofiler provides the Chrome DevTools Protocol
// commands, types, and events for the HeapProfiler domain.
//
// Generated by the cdproto-gen command.
package heapprofiler

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/seddonm1/cdproto/cdp"
	"github.com/seddonm1/cdproto/runtime"
)

// AddInspectedHeapObjectParams enables console to refer to the node with
// given id via $x (see Command Line API for more details $x functions).
type AddInspectedHeapObjectParams struct {
	HeapObjectID HeapSnapshotObjectID `json:"heapObjectId"` // Heap snapshot object id to be accessible by means of $x command line API.
}

// AddInspectedHeapObject enables console to refer to the node with given id
// via $x (see Command Line API for more details $x functions).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-addInspectedHeapObject
//
// parameters:
//   heapObjectID - Heap snapshot object id to be accessible by means of $x command line API.
func AddInspectedHeapObject(heapObjectID HeapSnapshotObjectID) *AddInspectedHeapObjectParams {
	return &AddInspectedHeapObjectParams{
		HeapObjectID: heapObjectID,
	}
}

// Do executes HeapProfiler.addInspectedHeapObject against the provided context.
func (p *AddInspectedHeapObjectParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandAddInspectedHeapObject, p, nil)
}

// CollectGarbageParams [no description].
type CollectGarbageParams struct{}

// CollectGarbage [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-collectGarbage
func CollectGarbage() *CollectGarbageParams {
	return &CollectGarbageParams{}
}

// Do executes HeapProfiler.collectGarbage against the provided context.
func (p *CollectGarbageParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandCollectGarbage, nil, nil)
}

// DisableParams [no description].
type DisableParams struct{}

// Disable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes HeapProfiler.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams [no description].
type EnableParams struct{}

// Enable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes HeapProfiler.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// GetHeapObjectIDParams [no description].
type GetHeapObjectIDParams struct {
	ObjectID runtime.RemoteObjectID `json:"objectId"` // Identifier of the object to get heap object id for.
}

// GetHeapObjectID [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-getHeapObjectId
//
// parameters:
//   objectID - Identifier of the object to get heap object id for.
func GetHeapObjectID(objectID runtime.RemoteObjectID) *GetHeapObjectIDParams {
	return &GetHeapObjectIDParams{
		ObjectID: objectID,
	}
}

// GetHeapObjectIDReturns return values.
type GetHeapObjectIDReturns struct {
	HeapSnapshotObjectID HeapSnapshotObjectID `json:"heapSnapshotObjectId,omitempty"` // Id of the heap snapshot object corresponding to the passed remote object id.
}

// Do executes HeapProfiler.getHeapObjectId against the provided context.
//
// returns:
//   heapSnapshotObjectID - Id of the heap snapshot object corresponding to the passed remote object id.
func (p *GetHeapObjectIDParams) Do(ctx context.Context) (heapSnapshotObjectID HeapSnapshotObjectID, err error) {
	// execute
	var res GetHeapObjectIDReturns
	err = cdp.Execute(ctx, CommandGetHeapObjectID, p, &res)
	if err != nil {
		return "", err
	}

	return res.HeapSnapshotObjectID, nil
}

// GetObjectByHeapObjectIDParams [no description].
type GetObjectByHeapObjectIDParams struct {
	ObjectID    HeapSnapshotObjectID `json:"objectId"`
	ObjectGroup string               `json:"objectGroup,omitempty"` // Symbolic group name that can be used to release multiple objects.
}

// GetObjectByHeapObjectID [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-getObjectByHeapObjectId
//
// parameters:
//   objectID
func GetObjectByHeapObjectID(objectID HeapSnapshotObjectID) *GetObjectByHeapObjectIDParams {
	return &GetObjectByHeapObjectIDParams{
		ObjectID: objectID,
	}
}

// WithObjectGroup symbolic group name that can be used to release multiple
// objects.
func (p GetObjectByHeapObjectIDParams) WithObjectGroup(objectGroup string) *GetObjectByHeapObjectIDParams {
	p.ObjectGroup = objectGroup
	return &p
}

// GetObjectByHeapObjectIDReturns return values.
type GetObjectByHeapObjectIDReturns struct {
	Result *runtime.RemoteObject `json:"result,omitempty"` // Evaluation result.
}

// Do executes HeapProfiler.getObjectByHeapObjectId against the provided context.
//
// returns:
//   result - Evaluation result.
func (p *GetObjectByHeapObjectIDParams) Do(ctx context.Context) (result *runtime.RemoteObject, err error) {
	// execute
	var res GetObjectByHeapObjectIDReturns
	err = cdp.Execute(ctx, CommandGetObjectByHeapObjectID, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

// GetSamplingProfileParams [no description].
type GetSamplingProfileParams struct{}

// GetSamplingProfile [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-getSamplingProfile
func GetSamplingProfile() *GetSamplingProfileParams {
	return &GetSamplingProfileParams{}
}

// GetSamplingProfileReturns return values.
type GetSamplingProfileReturns struct {
	Profile *SamplingHeapProfile `json:"profile,omitempty"` // Return the sampling profile being collected.
}

// Do executes HeapProfiler.getSamplingProfile against the provided context.
//
// returns:
//   profile - Return the sampling profile being collected.
func (p *GetSamplingProfileParams) Do(ctx context.Context) (profile *SamplingHeapProfile, err error) {
	// execute
	var res GetSamplingProfileReturns
	err = cdp.Execute(ctx, CommandGetSamplingProfile, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Profile, nil
}

// StartSamplingParams [no description].
type StartSamplingParams struct {
	SamplingInterval float64 `json:"samplingInterval,omitempty"` // Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
}

// StartSampling [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-startSampling
//
// parameters:
func StartSampling() *StartSamplingParams {
	return &StartSamplingParams{}
}

// WithSamplingInterval average sample interval in bytes. Poisson
// distribution is used for the intervals. The default value is 32768 bytes.
func (p StartSamplingParams) WithSamplingInterval(samplingInterval float64) *StartSamplingParams {
	p.SamplingInterval = samplingInterval
	return &p
}

// Do executes HeapProfiler.startSampling against the provided context.
func (p *StartSamplingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartSampling, p, nil)
}

// StartTrackingHeapObjectsParams [no description].
type StartTrackingHeapObjectsParams struct {
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

// StartTrackingHeapObjects [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-startTrackingHeapObjects
//
// parameters:
func StartTrackingHeapObjects() *StartTrackingHeapObjectsParams {
	return &StartTrackingHeapObjectsParams{}
}

// WithTrackAllocations [no description].
func (p StartTrackingHeapObjectsParams) WithTrackAllocations(trackAllocations bool) *StartTrackingHeapObjectsParams {
	p.TrackAllocations = trackAllocations
	return &p
}

// Do executes HeapProfiler.startTrackingHeapObjects against the provided context.
func (p *StartTrackingHeapObjectsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStartTrackingHeapObjects, p, nil)
}

// StopSamplingParams [no description].
type StopSamplingParams struct{}

// StopSampling [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-stopSampling
func StopSampling() *StopSamplingParams {
	return &StopSamplingParams{}
}

// StopSamplingReturns return values.
type StopSamplingReturns struct {
	Profile *SamplingHeapProfile `json:"profile,omitempty"` // Recorded sampling heap profile.
}

// Do executes HeapProfiler.stopSampling against the provided context.
//
// returns:
//   profile - Recorded sampling heap profile.
func (p *StopSamplingParams) Do(ctx context.Context) (profile *SamplingHeapProfile, err error) {
	// execute
	var res StopSamplingReturns
	err = cdp.Execute(ctx, CommandStopSampling, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Profile, nil
}

// StopTrackingHeapObjectsParams [no description].
type StopTrackingHeapObjectsParams struct {
	ReportProgress bool `json:"reportProgress,omitempty"` // If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
}

// StopTrackingHeapObjects [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-stopTrackingHeapObjects
//
// parameters:
func StopTrackingHeapObjects() *StopTrackingHeapObjectsParams {
	return &StopTrackingHeapObjectsParams{}
}

// WithReportProgress if true 'reportHeapSnapshotProgress' events will be
// generated while snapshot is being taken when the tracking is stopped.
func (p StopTrackingHeapObjectsParams) WithReportProgress(reportProgress bool) *StopTrackingHeapObjectsParams {
	p.ReportProgress = reportProgress
	return &p
}

// Do executes HeapProfiler.stopTrackingHeapObjects against the provided context.
func (p *StopTrackingHeapObjectsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandStopTrackingHeapObjects, p, nil)
}

// TakeHeapSnapshotParams [no description].
type TakeHeapSnapshotParams struct {
	ReportProgress bool `json:"reportProgress,omitempty"` // If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
}

// TakeHeapSnapshot [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#method-takeHeapSnapshot
//
// parameters:
func TakeHeapSnapshot() *TakeHeapSnapshotParams {
	return &TakeHeapSnapshotParams{}
}

// WithReportProgress if true 'reportHeapSnapshotProgress' events will be
// generated while snapshot is being taken.
func (p TakeHeapSnapshotParams) WithReportProgress(reportProgress bool) *TakeHeapSnapshotParams {
	p.ReportProgress = reportProgress
	return &p
}

// Do executes HeapProfiler.takeHeapSnapshot against the provided context.
func (p *TakeHeapSnapshotParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTakeHeapSnapshot, p, nil)
}

// Command names.
const (
	CommandAddInspectedHeapObject   = "HeapProfiler.addInspectedHeapObject"
	CommandCollectGarbage           = "HeapProfiler.collectGarbage"
	CommandDisable                  = "HeapProfiler.disable"
	CommandEnable                   = "HeapProfiler.enable"
	CommandGetHeapObjectID          = "HeapProfiler.getHeapObjectId"
	CommandGetObjectByHeapObjectID  = "HeapProfiler.getObjectByHeapObjectId"
	CommandGetSamplingProfile       = "HeapProfiler.getSamplingProfile"
	CommandStartSampling            = "HeapProfiler.startSampling"
	CommandStartTrackingHeapObjects = "HeapProfiler.startTrackingHeapObjects"
	CommandStopSampling             = "HeapProfiler.stopSampling"
	CommandStopTrackingHeapObjects  = "HeapProfiler.stopTrackingHeapObjects"
	CommandTakeHeapSnapshot         = "HeapProfiler.takeHeapSnapshot"
)
