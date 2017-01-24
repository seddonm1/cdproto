// AUTOGENERATED FILE: easyjson marshaler/unmarshalers.

package profiler

import (
	json "encoding/json"
	debugger "github.com/knq/chromedp/cdp/debugger"
	runtime "github.com/knq/chromedp/cdp/runtime"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler(in *jlexer.Lexer, out *StopReturns) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "profile":
			if in.IsNull() {
				in.Skip()
				out.Profile = nil
			} else {
				if out.Profile == nil {
					out.Profile = new(Profile)
				}
				(*out.Profile).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler(out *jwriter.Writer, in StopReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Profile != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"profile\":")
		if in.Profile == nil {
			out.RawString("null")
		} else {
			(*in.Profile).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler1(in *jlexer.Lexer, out *StopParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler1(out *jwriter.Writer, in StopParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StopParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StopParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StopParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StopParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler2(in *jlexer.Lexer, out *StartParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler2(out *jwriter.Writer, in StartParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StartParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StartParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StartParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StartParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler3(in *jlexer.Lexer, out *SetSamplingIntervalParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "interval":
			out.Interval = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler3(out *jwriter.Writer, in SetSamplingIntervalParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"interval\":")
	out.Int64(int64(in.Interval))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetSamplingIntervalParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetSamplingIntervalParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetSamplingIntervalParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetSamplingIntervalParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler4(in *jlexer.Lexer, out *ProfileNode) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "callFrame":
			if in.IsNull() {
				in.Skip()
				out.CallFrame = nil
			} else {
				if out.CallFrame == nil {
					out.CallFrame = new(runtime.CallFrame)
				}
				(*out.CallFrame).UnmarshalEasyJSON(in)
			}
		case "hitCount":
			out.HitCount = int64(in.Int64())
		case "children":
			if in.IsNull() {
				in.Skip()
				out.Children = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Children = make([]int64, 0, 8)
				} else {
					out.Children = []int64{}
				}
				for !in.IsDelim(']') {
					var v1 int64
					v1 = int64(in.Int64())
					out.Children = append(out.Children, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "deoptReason":
			out.DeoptReason = string(in.String())
		case "positionTicks":
			if in.IsNull() {
				in.Skip()
				out.PositionTicks = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.PositionTicks = make([]*PositionTickInfo, 0, 8)
				} else {
					out.PositionTicks = []*PositionTickInfo{}
				}
				for !in.IsDelim(']') {
					var v2 *PositionTickInfo
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(PositionTickInfo)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.PositionTicks = append(out.PositionTicks, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler4(out *jwriter.Writer, in ProfileNode) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Int64(int64(in.ID))
	}
	if in.CallFrame != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"callFrame\":")
		if in.CallFrame == nil {
			out.RawString("null")
		} else {
			(*in.CallFrame).MarshalEasyJSON(out)
		}
	}
	if in.HitCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"hitCount\":")
		out.Int64(int64(in.HitCount))
	}
	if len(in.Children) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"children\":")
		if in.Children == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Children {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v4))
			}
			out.RawByte(']')
		}
	}
	if in.DeoptReason != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"deoptReason\":")
		out.String(string(in.DeoptReason))
	}
	if len(in.PositionTicks) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"positionTicks\":")
		if in.PositionTicks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.PositionTicks {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler5(in *jlexer.Lexer, out *Profile) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nodes":
			if in.IsNull() {
				in.Skip()
				out.Nodes = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Nodes = make([]*ProfileNode, 0, 8)
				} else {
					out.Nodes = []*ProfileNode{}
				}
				for !in.IsDelim(']') {
					var v7 *ProfileNode
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(ProfileNode)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Nodes = append(out.Nodes, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "startTime":
			out.StartTime = float64(in.Float64())
		case "endTime":
			out.EndTime = float64(in.Float64())
		case "samples":
			if in.IsNull() {
				in.Skip()
				out.Samples = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Samples = make([]int64, 0, 8)
				} else {
					out.Samples = []int64{}
				}
				for !in.IsDelim(']') {
					var v8 int64
					v8 = int64(in.Int64())
					out.Samples = append(out.Samples, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "timeDeltas":
			if in.IsNull() {
				in.Skip()
				out.TimeDeltas = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.TimeDeltas = make([]int64, 0, 8)
				} else {
					out.TimeDeltas = []int64{}
				}
				for !in.IsDelim(']') {
					var v9 int64
					v9 = int64(in.Int64())
					out.TimeDeltas = append(out.TimeDeltas, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler5(out *jwriter.Writer, in Profile) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Nodes) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"nodes\":")
		if in.Nodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Nodes {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.StartTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"startTime\":")
		out.Float64(float64(in.StartTime))
	}
	if in.EndTime != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"endTime\":")
		out.Float64(float64(in.EndTime))
	}
	if len(in.Samples) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"samples\":")
		if in.Samples == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Samples {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v13))
			}
			out.RawByte(']')
		}
	}
	if len(in.TimeDeltas) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"timeDeltas\":")
		if in.TimeDeltas == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.TimeDeltas {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Profile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Profile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Profile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Profile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler6(in *jlexer.Lexer, out *PositionTickInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "line":
			out.Line = int64(in.Int64())
		case "ticks":
			out.Ticks = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler6(out *jwriter.Writer, in PositionTickInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Line != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"line\":")
		out.Int64(int64(in.Line))
	}
	if in.Ticks != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"ticks\":")
		out.Int64(int64(in.Ticks))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionTickInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionTickInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionTickInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionTickInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler7(in *jlexer.Lexer, out *EventConsoleProfileStarted) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "location":
			if in.IsNull() {
				in.Skip()
				out.Location = nil
			} else {
				if out.Location == nil {
					out.Location = new(debugger.Location)
				}
				(*out.Location).UnmarshalEasyJSON(in)
			}
		case "title":
			out.Title = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler7(out *jwriter.Writer, in EventConsoleProfileStarted) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Location != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location\":")
		if in.Location == nil {
			out.RawString("null")
		} else {
			(*in.Location).MarshalEasyJSON(out)
		}
	}
	if in.Title != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"title\":")
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventConsoleProfileStarted) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventConsoleProfileStarted) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventConsoleProfileStarted) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventConsoleProfileStarted) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler7(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler8(in *jlexer.Lexer, out *EventConsoleProfileFinished) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "location":
			if in.IsNull() {
				in.Skip()
				out.Location = nil
			} else {
				if out.Location == nil {
					out.Location = new(debugger.Location)
				}
				(*out.Location).UnmarshalEasyJSON(in)
			}
		case "profile":
			if in.IsNull() {
				in.Skip()
				out.Profile = nil
			} else {
				if out.Profile == nil {
					out.Profile = new(Profile)
				}
				(*out.Profile).UnmarshalEasyJSON(in)
			}
		case "title":
			out.Title = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler8(out *jwriter.Writer, in EventConsoleProfileFinished) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Location != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"location\":")
		if in.Location == nil {
			out.RawString("null")
		} else {
			(*in.Location).MarshalEasyJSON(out)
		}
	}
	if in.Profile != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"profile\":")
		if in.Profile == nil {
			out.RawString("null")
		} else {
			(*in.Profile).MarshalEasyJSON(out)
		}
	}
	if in.Title != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"title\":")
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventConsoleProfileFinished) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventConsoleProfileFinished) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventConsoleProfileFinished) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventConsoleProfileFinished) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler8(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler9(in *jlexer.Lexer, out *EnableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler9(out *jwriter.Writer, in EnableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler9(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler10(in *jlexer.Lexer, out *DisableParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler10(out *jwriter.Writer, in DisableParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DisableParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DisableParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpProfiler10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DisableParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DisableParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpProfiler10(l, v)
}
