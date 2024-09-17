// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
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

func easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel(in *jlexer.Lexer, out *Content) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			out.Title = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "url":
			out.Url = string(in.String())
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
func easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel(out *jwriter.Writer, in Content) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.Url))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Content) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Content) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Content) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Content) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel(l, v)
}
func easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel1(in *jlexer.Lexer, out *BannerGetRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tag_ids":
			out.TagIds = int(in.Int())
		case "feature_id":
			out.FeatureId = int(in.Int())
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
func easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel1(out *jwriter.Writer, in BannerGetRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tag_ids\":"
		out.RawString(prefix[1:])
		out.Int(int(in.TagIds))
	}
	{
		const prefix string = ",\"feature_id\":"
		out.RawString(prefix)
		out.Int(int(in.FeatureId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BannerGetRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BannerGetRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BannerGetRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BannerGetRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel1(l, v)
}
func easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel2(in *jlexer.Lexer, out *BannerFilterResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "banner_id":
			out.BannerId = int(in.Int())
		case "tag_ids":
			if in.IsNull() {
				in.Skip()
				out.TagIds = nil
			} else {
				in.Delim('[')
				if out.TagIds == nil {
					if !in.IsDelim(']') {
						out.TagIds = make([]int, 0, 8)
					} else {
						out.TagIds = []int{}
					}
				} else {
					out.TagIds = (out.TagIds)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int
					v1 = int(in.Int())
					out.TagIds = append(out.TagIds, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "feature_id":
			out.FeatureId = int(in.Int())
		case "content":
			(out.Content).UnmarshalEasyJSON(in)
		case "is_active":
			out.IsActive = bool(in.Bool())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel2(out *jwriter.Writer, in BannerFilterResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"banner_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.BannerId))
	}
	{
		const prefix string = ",\"tag_ids\":"
		out.RawString(prefix)
		if in.TagIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.TagIds {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"feature_id\":"
		out.RawString(prefix)
		out.Int(int(in.FeatureId))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(in.Content).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"is_active\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsActive))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BannerFilterResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BannerFilterResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BannerFilterResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BannerFilterResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel2(l, v)
}
func easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel3(in *jlexer.Lexer, out *BannerFilterGetRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tag_ids":
			out.TagIds = int(in.Int())
		case "feature_id":
			out.FeatureId = int(in.Int())
		case "limit":
			out.Limit = int(in.Int())
		case "offset":
			out.Offset = int(in.Int())
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
func easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel3(out *jwriter.Writer, in BannerFilterGetRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tag_ids\":"
		out.RawString(prefix[1:])
		out.Int(int(in.TagIds))
	}
	{
		const prefix string = ",\"feature_id\":"
		out.RawString(prefix)
		out.Int(int(in.FeatureId))
	}
	{
		const prefix string = ",\"limit\":"
		out.RawString(prefix)
		out.Int(int(in.Limit))
	}
	{
		const prefix string = ",\"offset\":"
		out.RawString(prefix)
		out.Int(int(in.Offset))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BannerFilterGetRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BannerFilterGetRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BannerFilterGetRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BannerFilterGetRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel3(l, v)
}
func easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel4(in *jlexer.Lexer, out *BannerCreateOrUpdateRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tag_ids":
			if in.IsNull() {
				in.Skip()
				out.TagIds = nil
			} else {
				in.Delim('[')
				if out.TagIds == nil {
					if !in.IsDelim(']') {
						out.TagIds = make([]int, 0, 8)
					} else {
						out.TagIds = []int{}
					}
				} else {
					out.TagIds = (out.TagIds)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int
					v4 = int(in.Int())
					out.TagIds = append(out.TagIds, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "feature_id":
			out.FeatureId = int(in.Int())
		case "content":
			(out.Content).UnmarshalEasyJSON(in)
		case "is_active":
			out.IsActive = bool(in.Bool())
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
func easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel4(out *jwriter.Writer, in BannerCreateOrUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tag_ids\":"
		out.RawString(prefix[1:])
		if in.TagIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.TagIds {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"feature_id\":"
		out.RawString(prefix)
		out.Int(int(in.FeatureId))
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(in.Content).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"is_active\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsActive))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BannerCreateOrUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BannerCreateOrUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BannerCreateOrUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BannerCreateOrUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel4(l, v)
}
func easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel5(in *jlexer.Lexer, out *Banner) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tag_ids":
			if in.IsNull() {
				in.Skip()
				out.TagIds = nil
			} else {
				in.Delim('[')
				if out.TagIds == nil {
					if !in.IsDelim(']') {
						out.TagIds = make([]int, 0, 8)
					} else {
						out.TagIds = []int{}
					}
				} else {
					out.TagIds = (out.TagIds)[:0]
				}
				for !in.IsDelim(']') {
					var v7 int
					v7 = int(in.Int())
					out.TagIds = append(out.TagIds, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "feature_ids":
			if in.IsNull() {
				in.Skip()
				out.Feature_id = nil
			} else {
				in.Delim('[')
				if out.Feature_id == nil {
					if !in.IsDelim(']') {
						out.Feature_id = make([]int, 0, 8)
					} else {
						out.Feature_id = []int{}
					}
				} else {
					out.Feature_id = (out.Feature_id)[:0]
				}
				for !in.IsDelim(']') {
					var v8 int
					v8 = int(in.Int())
					out.Feature_id = append(out.Feature_id, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "content":
			(out.Content).UnmarshalEasyJSON(in)
		case "is_active":
			out.Is_active = bool(in.Bool())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "banner_id":
			out.Banner_id = int(in.Int())
		case "article":
			out.Article = int(in.Int())
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
func easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel5(out *jwriter.Writer, in Banner) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tag_ids\":"
		out.RawString(prefix[1:])
		if in.TagIds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.TagIds {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"feature_ids\":"
		out.RawString(prefix)
		if in.Feature_id == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Feature_id {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		(in.Content).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"is_active\":"
		out.RawString(prefix)
		out.Bool(bool(in.Is_active))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"banner_id\":"
		out.RawString(prefix)
		out.Int(int(in.Banner_id))
	}
	{
		const prefix string = ",\"article\":"
		out.RawString(prefix)
		out.Int(int(in.Article))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Banner) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Banner) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC16456caEncodeAvitoTestBackendGolangInternalModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Banner) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Banner) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC16456caDecodeAvitoTestBackendGolangInternalModel5(l, v)
}