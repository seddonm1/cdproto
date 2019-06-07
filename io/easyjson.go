// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package io

import (
	json "encoding/json"
	runtime "github.com/seddonm1/cdproto/runtime"
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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo(in *jlexer.Lexer, out *ResolveBlobReturns) {
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
		case "uuid":
			out.UUID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo(out *jwriter.Writer, in ResolveBlobReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UUID != "" {
		const prefix string = ",\"uuid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResolveBlobReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResolveBlobReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResolveBlobReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResolveBlobReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo1(in *jlexer.Lexer, out *ResolveBlobParams) {
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
		case "objectId":
			out.ObjectID = runtime.RemoteObjectID(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo1(out *jwriter.Writer, in ResolveBlobParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"objectId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ObjectID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResolveBlobParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResolveBlobParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResolveBlobParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResolveBlobParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo2(in *jlexer.Lexer, out *ReadReturns) {
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
		case "base64Encoded":
			out.Base64encoded = bool(in.Bool())
		case "data":
			out.Data = string(in.String())
		case "eof":
			out.EOF = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo2(out *jwriter.Writer, in ReadReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Base64encoded {
		const prefix string = ",\"base64Encoded\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Base64encoded))
	}
	if in.Data != "" {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Data))
	}
	if in.EOF {
		const prefix string = ",\"eof\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.EOF))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReadReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReadReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReadReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReadReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo3(in *jlexer.Lexer, out *ReadParams) {
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
		case "handle":
			out.Handle = StreamHandle(in.String())
		case "offset":
			out.Offset = int64(in.Int64())
		case "size":
			out.Size = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo3(out *jwriter.Writer, in ReadParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"handle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Handle))
	}
	if in.Offset != 0 {
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Offset))
	}
	if in.Size != 0 {
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Size))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReadParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReadParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReadParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReadParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo4(in *jlexer.Lexer, out *CloseParams) {
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
		case "handle":
			out.Handle = StreamHandle(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo4(out *jwriter.Writer, in CloseParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"handle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Handle))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CloseParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CloseParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoIo4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CloseParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CloseParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoIo4(l, v)
}
