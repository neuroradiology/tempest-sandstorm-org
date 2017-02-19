package bridgeproxy

// AUTO GENERATED - DO NOT EDIT

import (
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type ProxyClaimRequestRequest struct{ capnp.Struct }

// ProxyClaimRequestRequest_TypeID is the unique identifier for the type ProxyClaimRequestRequest.
const ProxyClaimRequestRequest_TypeID = 0xe85d79ab9bde4783

func NewProxyClaimRequestRequest(s *capnp.Segment) (ProxyClaimRequestRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ProxyClaimRequestRequest{st}, err
}

func NewRootProxyClaimRequestRequest(s *capnp.Segment) (ProxyClaimRequestRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ProxyClaimRequestRequest{st}, err
}

func ReadRootProxyClaimRequestRequest(msg *capnp.Message) (ProxyClaimRequestRequest, error) {
	root, err := msg.RootPtr()
	return ProxyClaimRequestRequest{root.Struct()}, err
}

func (s ProxyClaimRequestRequest) String() string {
	str, _ := text.Marshal(0xe85d79ab9bde4783, s.Struct)
	return str
}

func (s ProxyClaimRequestRequest) RequestToken() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ProxyClaimRequestRequest) HasRequestToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ProxyClaimRequestRequest) RequestTokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ProxyClaimRequestRequest) SetRequestToken(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s ProxyClaimRequestRequest) RequiredPermissions() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s ProxyClaimRequestRequest) HasRequiredPermissions() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ProxyClaimRequestRequest) SetRequiredPermissions(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s ProxyClaimRequestRequest) NewRequiredPermissions(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s ProxyClaimRequestRequest) Label() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s ProxyClaimRequestRequest) HasLabel() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ProxyClaimRequestRequest) SetLabel(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewLabel sets the label field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s ProxyClaimRequestRequest) NewLabel() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// ProxyClaimRequestRequest_List is a list of ProxyClaimRequestRequest.
type ProxyClaimRequestRequest_List struct{ capnp.List }

// NewProxyClaimRequestRequest creates a new list of ProxyClaimRequestRequest.
func NewProxyClaimRequestRequest_List(s *capnp.Segment, sz int32) (ProxyClaimRequestRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return ProxyClaimRequestRequest_List{l}, err
}

func (s ProxyClaimRequestRequest_List) At(i int) ProxyClaimRequestRequest {
	return ProxyClaimRequestRequest{s.List.Struct(i)}
}

func (s ProxyClaimRequestRequest_List) Set(i int, v ProxyClaimRequestRequest) error {
	return s.List.SetStruct(i, v.Struct)
}

// ProxyClaimRequestRequest_Promise is a wrapper for a ProxyClaimRequestRequest promised by a client call.
type ProxyClaimRequestRequest_Promise struct{ *capnp.Pipeline }

func (p ProxyClaimRequestRequest_Promise) Struct() (ProxyClaimRequestRequest, error) {
	s, err := p.Pipeline.Struct()
	return ProxyClaimRequestRequest{s}, err
}

func (p ProxyClaimRequestRequest_Promise) Label() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type ProxyClaimRequestResponse struct{ capnp.Struct }

// ProxyClaimRequestResponse_TypeID is the unique identifier for the type ProxyClaimRequestResponse.
const ProxyClaimRequestResponse_TypeID = 0x88b48ab0daf5d8d0

func NewProxyClaimRequestResponse(s *capnp.Segment) (ProxyClaimRequestResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ProxyClaimRequestResponse{st}, err
}

func NewRootProxyClaimRequestResponse(s *capnp.Segment) (ProxyClaimRequestResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ProxyClaimRequestResponse{st}, err
}

func ReadRootProxyClaimRequestResponse(msg *capnp.Message) (ProxyClaimRequestResponse, error) {
	root, err := msg.RootPtr()
	return ProxyClaimRequestResponse{root.Struct()}, err
}

func (s ProxyClaimRequestResponse) String() string {
	str, _ := text.Marshal(0x88b48ab0daf5d8d0, s.Struct)
	return str
}

func (s ProxyClaimRequestResponse) Cap() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ProxyClaimRequestResponse) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ProxyClaimRequestResponse) CapBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ProxyClaimRequestResponse) SetCap(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// ProxyClaimRequestResponse_List is a list of ProxyClaimRequestResponse.
type ProxyClaimRequestResponse_List struct{ capnp.List }

// NewProxyClaimRequestResponse creates a new list of ProxyClaimRequestResponse.
func NewProxyClaimRequestResponse_List(s *capnp.Segment, sz int32) (ProxyClaimRequestResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ProxyClaimRequestResponse_List{l}, err
}

func (s ProxyClaimRequestResponse_List) At(i int) ProxyClaimRequestResponse {
	return ProxyClaimRequestResponse{s.List.Struct(i)}
}

func (s ProxyClaimRequestResponse_List) Set(i int, v ProxyClaimRequestResponse) error {
	return s.List.SetStruct(i, v.Struct)
}

// ProxyClaimRequestResponse_Promise is a wrapper for a ProxyClaimRequestResponse promised by a client call.
type ProxyClaimRequestResponse_Promise struct{ *capnp.Pipeline }

func (p ProxyClaimRequestResponse_Promise) Struct() (ProxyClaimRequestResponse, error) {
	s, err := p.Pipeline.Struct()
	return ProxyClaimRequestResponse{s}, err
}

const schema_bf72526e76ecd73b = "x\xda\x8c\x8e\xbfK\xebP\x00\x85\xcf\xb97}yC" +
	"\x7f\x85t\x7f\xf0\xe6\xf7\x04u\x10tPP\x11\x07!" +
	"\xb7\xb8vH\x9b\x8b\xa4\xa6I\x9a\xb4bqt\x11\xfd" +
	"\x13\xdc\\\x1dt\x10\x1c\xfd\x0b\xc4n\x8a\xa0\xa0\x93\x83" +
	"\xab\xa3D\xaeB\x07'\xb7\xc3\xc7\xc7\xc7\xa9\x8f\x97\xc4" +
	"t\xe9F\x02\xea_\xe9W1\xbe{\xbb??\xbc8" +
	"\x80\xe3\xb2X\xb8}\xdd\x89\x9b\xd9\x15J\xb4\x81\xd9k" +
	"v\xe9>\x99\xe9>p\x11,\xf6\xd7\x1e\x8fOG\xad" +
	"\x97o\xb24\xf2;\xdbt\x1da\xe4\x8a8CP\xb4" +
	"\xb30\xd8\xd2\xffS+KvGS\x1d?\x8d\xd3y" +
	"\xcf\xec\xe5\xc8\x0f{M\xdd\x1f\xea|\xd0\xd4yZK" +
	"\xe2\\{\xa4\xb2\xa4\x05X\x04\x9c\xca_@\xfd\x96T" +
	"\x0dA\xbb\xe3\xa7,C\xb0\x0c\xfe4\xda\x1f\xda:\x1f" +
	"\x98fy\xd2\\\xed\x02jERy\x82\x0e\xd9\xa0\x81" +
	"\x1b'\x80\xf2$U$\xe8\x08\xd1\xa0\x00\x9cp\x06P" +
	"\x81\xa4J\x05\x8b\xec\xab\xb9\x89Z\xb2\xad\xe3\xc9\x11\x83" +
	"\xc3L\x07\xf4t\xd6\x0b\xf3<\xb4\x938g\x15\xf4$" +
	"?\xa5*\xf8'\xf2\xdb:b\xbdX\x7f\xde\x9b\xcb." +
	"[G\x00Y\x07?\x02\x00\x00\xff\xff\xachg\x80"

func init() {
	schemas.Register(schema_bf72526e76ecd73b,
		0x88b48ab0daf5d8d0,
		0xe85d79ab9bde4783)
}
