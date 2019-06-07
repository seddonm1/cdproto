package cdp

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/knq/sysutil"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Executor is the common interface for executing a command.
type Executor interface {
	// Execute executes the command.
	Execute(context.Context, string, easyjson.Marshaler, easyjson.Unmarshaler) error
}

// contextKey is the context key type.
type contextKey int

// context keys.
const (
	executorKey contextKey = iota
)

// WithExecutor sets the message executor for the context.
func WithExecutor(parent context.Context, executor Executor) context.Context {
	return context.WithValue(parent, executorKey, executor)
}

// ExecutorFromContext returns the message executor for the context.
func ExecutorFromContext(ctx context.Context) Executor {
	return ctx.Value(executorKey).(Executor)
}

// Execute uses the context's message executor to send a command or event
// method marshaling the provided parameters, and unmarshaling to res.
func Execute(ctx context.Context, method string, params easyjson.Marshaler, res easyjson.Unmarshaler) error {
	if executor := ctx.Value(executorKey); executor != nil {
		return executor.(Executor).Execute(ctx, method, params, res)
	}
	return ErrInvalidContext
}

// Error is a error.
type Error string

// Error values.
const (
	// ErrInvalidContext is the invalid context error.
	ErrInvalidContext Error = "invalid context"

	// ErrMsgMissingParamsOrResult is the msg missing params or result error.
	ErrMsgMissingParamsOrResult Error = "msg missing params or result"
)

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// ErrUnknownCommandOrEvent is an unknown command or event error.
type ErrUnknownCommandOrEvent string

// Error satisfies the error interface.
func (err ErrUnknownCommandOrEvent) Error() string {
	return fmt.Sprintf("unknown command or event %q", string(err))
}

// NodeID unique DOM node identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-NodeId
type NodeID int64

// Int64 returns the NodeID as int64 value.
func (t NodeID) Int64() int64 {
	return int64(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *NodeID) UnmarshalEasyJSON(in *jlexer.Lexer) {
	buf := in.Raw()
	if l := len(buf); l > 2 && buf[0] == '"' && buf[l-1] == '"' {
		buf = buf[1 : l-1]
	}

	v, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		in.AddError(err)
	}

	*t = NodeID(v)
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *NodeID) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// BackendNodeID unique DOM node identifier used to reference a node that may
// not have been pushed to the front-end.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-BackendNodeId
type BackendNodeID int64

// Int64 returns the BackendNodeID as int64 value.
func (t BackendNodeID) Int64() int64 {
	return int64(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *BackendNodeID) UnmarshalEasyJSON(in *jlexer.Lexer) {
	buf := in.Raw()
	if l := len(buf); l > 2 && buf[0] == '"' && buf[l-1] == '"' {
		buf = buf[1 : l-1]
	}

	v, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		in.AddError(err)
	}

	*t = BackendNodeID(v)
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *BackendNodeID) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// BackendNode backend node with a friendly name.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-BackendNode
type BackendNode struct {
	NodeType      NodeType      `json:"nodeType"` // Node's nodeType.
	NodeName      string        `json:"nodeName"` // Node's nodeName.
	BackendNodeID BackendNodeID `json:"backendNodeId"`
}

// PseudoType pseudo element type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-PseudoType
type PseudoType string

// String returns the PseudoType as string value.
func (t PseudoType) String() string {
	return string(t)
}

// PseudoType values.
const (
	PseudoTypeFirstLine           PseudoType = "first-line"
	PseudoTypeFirstLetter         PseudoType = "first-letter"
	PseudoTypeBefore              PseudoType = "before"
	PseudoTypeAfter               PseudoType = "after"
	PseudoTypeBackdrop            PseudoType = "backdrop"
	PseudoTypeSelection           PseudoType = "selection"
	PseudoTypeFirstLineInherited  PseudoType = "first-line-inherited"
	PseudoTypeScrollbar           PseudoType = "scrollbar"
	PseudoTypeScrollbarThumb      PseudoType = "scrollbar-thumb"
	PseudoTypeScrollbarButton     PseudoType = "scrollbar-button"
	PseudoTypeScrollbarTrack      PseudoType = "scrollbar-track"
	PseudoTypeScrollbarTrackPiece PseudoType = "scrollbar-track-piece"
	PseudoTypeScrollbarCorner     PseudoType = "scrollbar-corner"
	PseudoTypeResizer             PseudoType = "resizer"
	PseudoTypeInputListButton     PseudoType = "input-list-button"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PseudoType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PseudoType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PseudoType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PseudoType(in.String()) {
	case PseudoTypeFirstLine:
		*t = PseudoTypeFirstLine
	case PseudoTypeFirstLetter:
		*t = PseudoTypeFirstLetter
	case PseudoTypeBefore:
		*t = PseudoTypeBefore
	case PseudoTypeAfter:
		*t = PseudoTypeAfter
	case PseudoTypeBackdrop:
		*t = PseudoTypeBackdrop
	case PseudoTypeSelection:
		*t = PseudoTypeSelection
	case PseudoTypeFirstLineInherited:
		*t = PseudoTypeFirstLineInherited
	case PseudoTypeScrollbar:
		*t = PseudoTypeScrollbar
	case PseudoTypeScrollbarThumb:
		*t = PseudoTypeScrollbarThumb
	case PseudoTypeScrollbarButton:
		*t = PseudoTypeScrollbarButton
	case PseudoTypeScrollbarTrack:
		*t = PseudoTypeScrollbarTrack
	case PseudoTypeScrollbarTrackPiece:
		*t = PseudoTypeScrollbarTrackPiece
	case PseudoTypeScrollbarCorner:
		*t = PseudoTypeScrollbarCorner
	case PseudoTypeResizer:
		*t = PseudoTypeResizer
	case PseudoTypeInputListButton:
		*t = PseudoTypeInputListButton

	default:
		in.AddError(errors.New("unknown PseudoType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PseudoType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ShadowRootType shadow root type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-ShadowRootType
type ShadowRootType string

// String returns the ShadowRootType as string value.
func (t ShadowRootType) String() string {
	return string(t)
}

// ShadowRootType values.
const (
	ShadowRootTypeUserAgent ShadowRootType = "user-agent"
	ShadowRootTypeOpen      ShadowRootType = "open"
	ShadowRootTypeClosed    ShadowRootType = "closed"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ShadowRootType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ShadowRootType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ShadowRootType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ShadowRootType(in.String()) {
	case ShadowRootTypeUserAgent:
		*t = ShadowRootTypeUserAgent
	case ShadowRootTypeOpen:
		*t = ShadowRootTypeOpen
	case ShadowRootTypeClosed:
		*t = ShadowRootTypeClosed

	default:
		in.AddError(errors.New("unknown ShadowRootType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ShadowRootType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Node DOM interaction is implemented in terms of mirror objects that
// represent the actual DOM nodes. DOMNode is a base node mirror type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-Node
type Node struct {
	NodeID           NodeID         `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the nodeId. Backend will only push node with given id once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	ParentID         NodeID         `json:"parentId,omitempty"`         // The id of the parent node if any.
	BackendNodeID    BackendNodeID  `json:"backendNodeId"`              // The BackendNodeId for this node.
	NodeType         NodeType       `json:"nodeType"`                   // Node's nodeType.
	NodeName         string         `json:"nodeName"`                   // Node's nodeName.
	LocalName        string         `json:"localName"`                  // Node's localName.
	NodeValue        string         `json:"nodeValue"`                  // Node's nodeValue.
	ChildNodeCount   int64          `json:"childNodeCount,omitempty"`   // Child count for Container nodes.
	Children         []*Node        `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string       `json:"attributes,omitempty"`       // Attributes of the Element node in the form of flat array [name1, value1, name2, value2].
	DocumentURL      string         `json:"documentURL,omitempty"`      // Document URL that Document or FrameOwner node points to.
	BaseURL          string         `json:"baseURL,omitempty"`          // Base URL that Document or FrameOwner node uses for URL completion.
	PublicID         string         `json:"publicId,omitempty"`         // DocumentType's publicId.
	SystemID         string         `json:"systemId,omitempty"`         // DocumentType's systemId.
	InternalSubset   string         `json:"internalSubset,omitempty"`   // DocumentType's internalSubset.
	XMLVersion       string         `json:"xmlVersion,omitempty"`       // Document's XML version in case of XML documents.
	Name             string         `json:"name,omitempty"`             // Attr's name.
	Value            string         `json:"value,omitempty"`            // Attr's value.
	PseudoType       PseudoType     `json:"pseudoType,omitempty"`       // Pseudo element type for this node.
	ShadowRootType   ShadowRootType `json:"shadowRootType,omitempty"`   // Shadow root type.
	FrameID          FrameID        `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *Node          `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []*Node        `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *Node          `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []*Node        `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	ImportedDocument *Node          `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
	DistributedNodes []*BackendNode `json:"distributedNodes,omitempty"` // Distributed nodes for given insertion point.
	IsSVG            bool           `json:"isSVG,omitempty"`            // Whether the node is SVG.
	Parent           *Node          `json:"-"`                          // Parent node.
	Invalidated      chan struct{}  `json:"-"`                          // Invalidated channel.
	State            NodeState      `json:"-"`                          // Node state.
	sync.RWMutex     `json:"-"`     // Read write mutex.
}

// AttributeValue returns the named attribute for the node.
func (n *Node) AttributeValue(name string) string {
	n.RLock()
	defer n.RUnlock()

	for i := 0; i < len(n.Attributes); i += 2 {
		if n.Attributes[i] == name {
			return n.Attributes[i+1]
		}
	}

	return ""
}

// jspathrecursive builds the jspath string.
func (n *Node) jspathrecursive(stopAtDocument, stopAtID bool) string {
	n.RLock()
	defer n.RUnlock()

	p := ""
	pos := ""
	id := n.AttributeValue("id")
	switch {
	case n.Parent == nil:
		return n.LocalName

	case stopAtDocument && n.NodeType == NodeTypeDocument:
		return ""

	case stopAtID && id != "":
		p = "/"
		pos = `[@id='` + id + `']`

	case n.Parent != nil:
		var i int
		var found bool

		n.Parent.RLock()
		for j := 0; j < len(n.Parent.Children); j++ {
			if n.Parent.Children[j].LocalName == n.LocalName {
				i++
			}
			if n.Parent.Children[j].NodeID == n.NodeID {
				found = true
				break
			}
		}
		n.Parent.RUnlock()

		if found {
			if i != 1 {
				pos = fmt.Sprintf(":nth-child(%d)", i)
			}
		}

		p = n.Parent.jspathrecursive(stopAtDocument, stopAtID)
	}

	if n.NodeName == "HTML" {
		return ""
	}
	if n.NodeName == "BODY" {
		return `document.querySelector("body`
	}
	if n.NodeType == NodeTypeDocumentFragment {
		return fmt.Sprintf(`%s").shadowRoot.querySelector("%s%s`, p, n.LocalName, pos)
	}
	if n.Parent.NodeType == NodeTypeDocumentFragment {
		return p + n.LocalName + pos
	}

	return p + " > " + n.LocalName + pos
}

// jspath builds the jspath string.
func (n *Node) jspath(stopAtDocument, stopAtID bool) string {
	return n.jspathrecursive(stopAtDocument, stopAtID) + `")`
}

// PartialJSPathByID returns the partial JSPath for the node, stopping at the
// first parent with an id attribute or at nearest parent document node.
func (n *Node) PartialJSPathByID() string {
	return n.jspath(true, true)
}

// PartialJSPath returns the partial JSPath for the node, stopping at the nearest
// parent document node.
func (n *Node) PartialJSPath() string {
	return n.jspath(true, false)
}

// FullJSPathByID returns the full JSPath for the node, stopping at the top most
// document root or at the closest parent node with an id attribute.
func (n *Node) FullJSPathByID() string {
	return n.jspath(false, true)
}

// FullJSPath returns the full JSPath for the node, stopping only at the top most
// document root.
func (n *Node) FullJSPath() string {
	a := n.jspath(false, false)
	fmt.Printf("%s\n", a)
	return a
}

// xpath builds the xpath string.
func (n *Node) xpath(stopAtDocument, stopAtID bool) string {
	n.RLock()
	defer n.RUnlock()

	p := ""
	pos := ""
	id := n.AttributeValue("id")
	switch {
	case n.Parent == nil:
		return n.LocalName

	case stopAtDocument && n.NodeType == NodeTypeDocument:
		return ""

	case stopAtID && id != "":
		p = "/"
		pos = `[@id='` + id + `']`

	case n.Parent != nil:
		var i int
		var found bool

		n.Parent.RLock()
		for j := 0; j < len(n.Parent.Children); j++ {
			if n.Parent.Children[j].LocalName == n.LocalName {
				i++
			}
			if n.Parent.Children[j].NodeID == n.NodeID {
				found = true
				break
			}
		}
		n.Parent.RUnlock()

		if found {
			pos = "[" + strconv.Itoa(i) + "]"
		}

		p = n.Parent.xpath(stopAtDocument, stopAtID)
	}

	return p + "/" + n.LocalName + pos
}

// PartialXPathByID returns the partial XPath for the node, stopping at the
// first parent with an id attribute or at nearest parent document node.
func (n *Node) PartialXPathByID() string {
	return n.xpath(true, true)
}

// PartialXPath returns the partial XPath for the node, stopping at the nearest
// parent document node.
func (n *Node) PartialXPath() string {
	return n.xpath(true, false)
}

// FullXPathByID returns the full XPath for the node, stopping at the top most
// document root or at the closest parent node with an id attribute.
func (n *Node) FullXPathByID() string {
	return n.xpath(false, true)
}

// FullXPath returns the full XPath for the node, stopping only at the top most
// document root.
func (n *Node) FullXPath() string {
	return n.xpath(false, false)
}

// NodeState is the state of a DOM node.
type NodeState uint8

// NodeState enum values.
const (
	NodeReady NodeState = 1 << (7 - iota)
	NodeVisible
	NodeHighlighted
)

// nodeStateNames are the names of the node states.
var nodeStateNames = map[NodeState]string{
	NodeReady:       "Ready",
	NodeVisible:     "Visible",
	NodeHighlighted: "Highlighted",
}

// String satisfies stringer interface.
func (ns NodeState) String() string {
	var s []string
	for k, v := range nodeStateNames {
		if ns&k != 0 {
			s = append(s, v)
		}
	}
	return "[" + strings.Join(s, " ") + "]"
}

// EmptyNodeID is the "non-existent" node id.
const EmptyNodeID = NodeID(0)

// RGBA a structure holding an RGBA color.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-RGBA
type RGBA struct {
	R int64   `json:"r"`           // The red component, in the [0-255] range.
	G int64   `json:"g"`           // The green component, in the [0-255] range.
	B int64   `json:"b"`           // The blue component, in the [0-255] range.
	A float64 `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}

// NodeType node type.
//
// See: https://developer.mozilla.org/en/docs/Web/API/Node/nodeType
type NodeType int64

// Int64 returns the NodeType as int64 value.
func (t NodeType) Int64() int64 {
	return int64(t)
}

// NodeType values.
const (
	NodeTypeElement               NodeType = 1
	NodeTypeAttribute             NodeType = 2
	NodeTypeText                  NodeType = 3
	NodeTypeCDATA                 NodeType = 4
	NodeTypeEntityReference       NodeType = 5
	NodeTypeEntity                NodeType = 6
	NodeTypeProcessingInstruction NodeType = 7
	NodeTypeComment               NodeType = 8
	NodeTypeDocument              NodeType = 9
	NodeTypeDocumentType          NodeType = 10
	NodeTypeDocumentFragment      NodeType = 11
	NodeTypeNotation              NodeType = 12
)

// String returns the NodeType as string value.
func (t NodeType) String() string {
	switch t {
	case NodeTypeElement:
		return "Element"
	case NodeTypeAttribute:
		return "Attribute"
	case NodeTypeText:
		return "Text"
	case NodeTypeCDATA:
		return "CDATA"
	case NodeTypeEntityReference:
		return "EntityReference"
	case NodeTypeEntity:
		return "Entity"
	case NodeTypeProcessingInstruction:
		return "ProcessingInstruction"
	case NodeTypeComment:
		return "Comment"
	case NodeTypeDocument:
		return "Document"
	case NodeTypeDocumentType:
		return "DocumentType"
	case NodeTypeDocumentFragment:
		return "DocumentFragment"
	case NodeTypeNotation:
		return "Notation"
	}

	return fmt.Sprintf("NodeType(%d)", t)
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t NodeType) MarshalEasyJSON(out *jwriter.Writer) {
	out.Int64(int64(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t NodeType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *NodeType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch NodeType(in.Int64()) {
	case NodeTypeElement:
		*t = NodeTypeElement
	case NodeTypeAttribute:
		*t = NodeTypeAttribute
	case NodeTypeText:
		*t = NodeTypeText
	case NodeTypeCDATA:
		*t = NodeTypeCDATA
	case NodeTypeEntityReference:
		*t = NodeTypeEntityReference
	case NodeTypeEntity:
		*t = NodeTypeEntity
	case NodeTypeProcessingInstruction:
		*t = NodeTypeProcessingInstruction
	case NodeTypeComment:
		*t = NodeTypeComment
	case NodeTypeDocument:
		*t = NodeTypeDocument
	case NodeTypeDocumentType:
		*t = NodeTypeDocumentType
	case NodeTypeDocumentFragment:
		*t = NodeTypeDocumentFragment
	case NodeTypeNotation:
		*t = NodeTypeNotation

	default:
		in.AddError(errors.New("unknown NodeType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *NodeType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// LoaderID unique loader identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#type-LoaderId
type LoaderID string

// String returns the LoaderID as string value.
func (t LoaderID) String() string {
	return string(t)
}

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#type-TimeSinceEpoch
type TimeSinceEpoch time.Time

// Time returns the TimeSinceEpoch as time.Time value.
func (t TimeSinceEpoch) Time() time.Time {
	return time.Time(t)
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TimeSinceEpoch) MarshalEasyJSON(out *jwriter.Writer) {
	v := float64(time.Time(t).UnixNano() / int64(time.Second))

	out.Buffer.EnsureSpace(20)
	out.Buffer.Buf = strconv.AppendFloat(out.Buffer.Buf, v, 'f', -1, 64)
}

// MarshalJSON satisfies json.Marshaler.
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalEasyJSON(in *jlexer.Lexer) {
	*t = TimeSinceEpoch(time.Unix(0, int64(in.Float64()*float64(time.Second))))
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// MonotonicTime monotonically increasing time in seconds since an arbitrary
// point in the past.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#type-MonotonicTime
type MonotonicTime time.Time

// Time returns the MonotonicTime as time.Time value.
func (t MonotonicTime) Time() time.Time {
	return time.Time(t)
}

// MonotonicTimeEpoch is the MonotonicTime time epoch.
var MonotonicTimeEpoch *time.Time

func init() {
	// initialize epoch
	bt := sysutil.BootTime()
	MonotonicTimeEpoch = &bt
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MonotonicTime) MarshalEasyJSON(out *jwriter.Writer) {
	v := float64(time.Time(t).Sub(*MonotonicTimeEpoch)) / float64(time.Second)

	out.Buffer.EnsureSpace(20)
	out.Buffer.Buf = strconv.AppendFloat(out.Buffer.Buf, v, 'f', -1, 64)
}

// MarshalJSON satisfies json.Marshaler.
func (t MonotonicTime) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MonotonicTime) UnmarshalEasyJSON(in *jlexer.Lexer) {
	*t = MonotonicTime(MonotonicTimeEpoch.Add(time.Duration(in.Float64() * float64(time.Second))))
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MonotonicTime) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// FrameID unique frame identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#type-FrameId
type FrameID string

// String returns the FrameID as string value.
func (t FrameID) String() string {
	return string(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *FrameID) UnmarshalEasyJSON(in *jlexer.Lexer) {
	buf := in.Raw()
	if l := len(buf); l > 2 && buf[0] == '"' && buf[l-1] == '"' {
		buf = buf[1 : l-1]
	}

	*t = FrameID(buf)
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *FrameID) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Frame information about the Frame on the page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#type-Frame
type Frame struct {
	ID             FrameID          `json:"id"`                       // Frame unique identifier.
	ParentID       FrameID          `json:"parentId,omitempty"`       // Parent frame identifier.
	LoaderID       LoaderID         `json:"loaderId"`                 // Identifier of the loader associated with this frame.
	Name           string           `json:"name,omitempty"`           // Frame's name as specified in the tag.
	URL            string           `json:"url"`                      // Frame document's URL.
	SecurityOrigin string           `json:"securityOrigin"`           // Frame document's security origin.
	MimeType       string           `json:"mimeType"`                 // Frame document's mimeType as determined by the browser.
	UnreachableURL string           `json:"unreachableUrl,omitempty"` // If the frame failed to load, this contains the URL that could not be loaded.
	State          FrameState       `json:"-"`                        // Frame state.
	Root           *Node            `json:"-"`                        // Frame document root.
	Nodes          map[NodeID]*Node `json:"-"`                        // Frame nodes.
	sync.RWMutex   `json:"-"`       // Read write mutex.
}

// FrameState is the state of a Frame.
type FrameState uint16

// FrameState enum values.
const (
	FrameDOMContentEventFired FrameState = 1 << (15 - iota)
	FrameLoadEventFired
	FrameAttached
	FrameNavigated
	FrameLoading
	FrameScheduledNavigation
)

// frameStateNames are the names of the frame states.
var frameStateNames = map[FrameState]string{
	FrameDOMContentEventFired: "DOMContentEventFired",
	FrameLoadEventFired:       "LoadEventFired",
	FrameAttached:             "Attached",
	FrameNavigated:            "Navigated",
	FrameLoading:              "Loading",
	FrameScheduledNavigation:  "ScheduledNavigation",
}

// String satisfies stringer interface.
func (fs FrameState) String() string {
	var s []string
	for k, v := range frameStateNames {
		if fs&k != 0 {
			s = append(s, v)
		}
	}
	return "[" + strings.Join(s, " ") + "]"
}

// EmptyFrameID is the "non-existent" frame id.
const EmptyFrameID = FrameID("")
