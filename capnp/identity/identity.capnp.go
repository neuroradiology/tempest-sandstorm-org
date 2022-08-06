// Code generated by capnpc-go. DO NOT EDIT.

package identity

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	util "zenhack.net/go/sandstorm/capnp/util"
)

type Identity capnp.Client

// Identity_TypeID is the unique identifier for the type Identity.
const Identity_TypeID = 0xc084987aa951dd18

func (c Identity) GetProfile(ctx context.Context, params func(Identity_getProfile_Params) error) (Identity_getProfile_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc084987aa951dd18,
			MethodID:      0,
			InterfaceName: "identity.capnp:Identity",
			MethodName:    "getProfile",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Identity_getProfile_Params(s)) }
	}
	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Identity_getProfile_Results_Future{Future: ans.Future()}, release
}

func (c Identity) AddRef() Identity {
	return Identity(capnp.Client(c).AddRef())
}

func (c Identity) Release() {
	capnp.Client(c).Release()
}

func (c Identity) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Identity) DecodeFromPtr(p capnp.Ptr) Identity {
	return Identity(capnp.Client{}.DecodeFromPtr(p))
}

func (c Identity) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// A Identity_Server is a Identity with a local implementation.
type Identity_Server interface {
	GetProfile(context.Context, Identity_getProfile) error
}

// Identity_NewServer creates a new Server from an implementation of Identity_Server.
func Identity_NewServer(s Identity_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Identity_Methods(nil, s), s, c)
}

// Identity_ServerToClient creates a new Client from an implementation of Identity_Server.
// The caller is responsible for calling Release on the returned Client.
func Identity_ServerToClient(s Identity_Server) Identity {
	return Identity(capnp.NewClient(Identity_NewServer(s)))
}

// Identity_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Identity_Methods(methods []server.Method, s Identity_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc084987aa951dd18,
			MethodID:      0,
			InterfaceName: "identity.capnp:Identity",
			MethodName:    "getProfile",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.GetProfile(ctx, Identity_getProfile{call})
		},
	})

	return methods
}

// Identity_getProfile holds the state for a server call to Identity.getProfile.
// See server.Call for documentation.
type Identity_getProfile struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Identity_getProfile) Args() Identity_getProfile_Params {
	return Identity_getProfile_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Identity_getProfile) AllocResults() (Identity_getProfile_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_getProfile_Results(r), err
}

// Identity_List is a list of Identity.
type Identity_List = capnp.CapList[Identity]

// NewIdentity creates a new list of Identity.
func NewIdentity_List(s *capnp.Segment, sz int32) (Identity_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Identity](l), err
}

type Identity_PowerboxTag capnp.Struct

// Identity_PowerboxTag_TypeID is the unique identifier for the type Identity_PowerboxTag.
const Identity_PowerboxTag_TypeID = 0xf35052cfb1d3bbe8

func NewIdentity_PowerboxTag(s *capnp.Segment) (Identity_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_PowerboxTag(st), err
}

func NewRootIdentity_PowerboxTag(s *capnp.Segment) (Identity_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_PowerboxTag(st), err
}

func ReadRootIdentity_PowerboxTag(msg *capnp.Message) (Identity_PowerboxTag, error) {
	root, err := msg.Root()
	return Identity_PowerboxTag(root.Struct()), err
}

func (s Identity_PowerboxTag) String() string {
	str, _ := text.Marshal(0xf35052cfb1d3bbe8, capnp.Struct(s))
	return str
}

func (s Identity_PowerboxTag) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Identity_PowerboxTag) DecodeFromPtr(p capnp.Ptr) Identity_PowerboxTag {
	return Identity_PowerboxTag(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Identity_PowerboxTag) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Identity_PowerboxTag) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Identity_PowerboxTag) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Identity_PowerboxTag) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Identity_PowerboxTag) Permissions() (capnp.BitList, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return capnp.BitList(p.List()), err
}

func (s Identity_PowerboxTag) HasPermissions() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Identity_PowerboxTag) SetPermissions(v capnp.BitList) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s Identity_PowerboxTag) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}

// Identity_PowerboxTag_List is a list of Identity_PowerboxTag.
type Identity_PowerboxTag_List = capnp.StructList[Identity_PowerboxTag]

// NewIdentity_PowerboxTag creates a new list of Identity_PowerboxTag.
func NewIdentity_PowerboxTag_List(s *capnp.Segment, sz int32) (Identity_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Identity_PowerboxTag](l), err
}

// Identity_PowerboxTag_Future is a wrapper for a Identity_PowerboxTag promised by a client call.
type Identity_PowerboxTag_Future struct{ *capnp.Future }

func (p Identity_PowerboxTag_Future) Struct() (Identity_PowerboxTag, error) {
	s, err := p.Future.Struct()
	return Identity_PowerboxTag(s), err
}

type Identity_getProfile_Params capnp.Struct

// Identity_getProfile_Params_TypeID is the unique identifier for the type Identity_getProfile_Params.
const Identity_getProfile_Params_TypeID = 0xf32d79a7b575d94c

func NewIdentity_getProfile_Params(s *capnp.Segment) (Identity_getProfile_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Identity_getProfile_Params(st), err
}

func NewRootIdentity_getProfile_Params(s *capnp.Segment) (Identity_getProfile_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Identity_getProfile_Params(st), err
}

func ReadRootIdentity_getProfile_Params(msg *capnp.Message) (Identity_getProfile_Params, error) {
	root, err := msg.Root()
	return Identity_getProfile_Params(root.Struct()), err
}

func (s Identity_getProfile_Params) String() string {
	str, _ := text.Marshal(0xf32d79a7b575d94c, capnp.Struct(s))
	return str
}

func (s Identity_getProfile_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Identity_getProfile_Params) DecodeFromPtr(p capnp.Ptr) Identity_getProfile_Params {
	return Identity_getProfile_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Identity_getProfile_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Identity_getProfile_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Identity_getProfile_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Identity_getProfile_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Identity_getProfile_Params_List is a list of Identity_getProfile_Params.
type Identity_getProfile_Params_List = capnp.StructList[Identity_getProfile_Params]

// NewIdentity_getProfile_Params creates a new list of Identity_getProfile_Params.
func NewIdentity_getProfile_Params_List(s *capnp.Segment, sz int32) (Identity_getProfile_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Identity_getProfile_Params](l), err
}

// Identity_getProfile_Params_Future is a wrapper for a Identity_getProfile_Params promised by a client call.
type Identity_getProfile_Params_Future struct{ *capnp.Future }

func (p Identity_getProfile_Params_Future) Struct() (Identity_getProfile_Params, error) {
	s, err := p.Future.Struct()
	return Identity_getProfile_Params(s), err
}

type Identity_getProfile_Results capnp.Struct

// Identity_getProfile_Results_TypeID is the unique identifier for the type Identity_getProfile_Results.
const Identity_getProfile_Results_TypeID = 0xcd7272b855c92f6d

func NewIdentity_getProfile_Results(s *capnp.Segment) (Identity_getProfile_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_getProfile_Results(st), err
}

func NewRootIdentity_getProfile_Results(s *capnp.Segment) (Identity_getProfile_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Identity_getProfile_Results(st), err
}

func ReadRootIdentity_getProfile_Results(msg *capnp.Message) (Identity_getProfile_Results, error) {
	root, err := msg.Root()
	return Identity_getProfile_Results(root.Struct()), err
}

func (s Identity_getProfile_Results) String() string {
	str, _ := text.Marshal(0xcd7272b855c92f6d, capnp.Struct(s))
	return str
}

func (s Identity_getProfile_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Identity_getProfile_Results) DecodeFromPtr(p capnp.Ptr) Identity_getProfile_Results {
	return Identity_getProfile_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Identity_getProfile_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Identity_getProfile_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Identity_getProfile_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Identity_getProfile_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Identity_getProfile_Results) Profile() (Profile, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Profile(p.Struct()), err
}

func (s Identity_getProfile_Results) HasProfile() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Identity_getProfile_Results) SetProfile(v Profile) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewProfile sets the profile field to a newly
// allocated Profile struct, preferring placement in s's segment.
func (s Identity_getProfile_Results) NewProfile() (Profile, error) {
	ss, err := NewProfile(capnp.Struct(s).Segment())
	if err != nil {
		return Profile{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

// Identity_getProfile_Results_List is a list of Identity_getProfile_Results.
type Identity_getProfile_Results_List = capnp.StructList[Identity_getProfile_Results]

// NewIdentity_getProfile_Results creates a new list of Identity_getProfile_Results.
func NewIdentity_getProfile_Results_List(s *capnp.Segment, sz int32) (Identity_getProfile_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Identity_getProfile_Results](l), err
}

// Identity_getProfile_Results_Future is a wrapper for a Identity_getProfile_Results promised by a client call.
type Identity_getProfile_Results_Future struct{ *capnp.Future }

func (p Identity_getProfile_Results_Future) Struct() (Identity_getProfile_Results, error) {
	s, err := p.Future.Struct()
	return Identity_getProfile_Results(s), err
}

func (p Identity_getProfile_Results_Future) Profile() Profile_Future {
	return Profile_Future{Future: p.Future.Field(0, nil)}
}

type Profile capnp.Struct

// Profile_TypeID is the unique identifier for the type Profile.
const Profile_TypeID = 0xd3d0c34d7201fcef

func NewProfile(s *capnp.Segment) (Profile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Profile(st), err
}

func NewRootProfile(s *capnp.Segment) (Profile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Profile(st), err
}

func ReadRootProfile(msg *capnp.Message) (Profile, error) {
	root, err := msg.Root()
	return Profile(root.Struct()), err
}

func (s Profile) String() string {
	str, _ := text.Marshal(0xd3d0c34d7201fcef, capnp.Struct(s))
	return str
}

func (s Profile) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Profile) DecodeFromPtr(p capnp.Ptr) Profile {
	return Profile(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Profile) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Profile) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Profile) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Profile) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Profile) DisplayName() (util.LocalizedText, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return util.LocalizedText(p.Struct()), err
}

func (s Profile) HasDisplayName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Profile) SetDisplayName(v util.LocalizedText) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewDisplayName sets the displayName field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s Profile) NewDisplayName() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(capnp.Struct(s).Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Profile) PreferredHandle() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Profile) HasPreferredHandle() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Profile) PreferredHandleBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Profile) SetPreferredHandle(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

func (s Profile) Picture() util.StaticAsset {
	p, _ := capnp.Struct(s).Ptr(2)
	return util.StaticAsset(p.Interface().Client())
}

func (s Profile) HasPicture() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Profile) SetPicture(v util.StaticAsset) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(2, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(2, in.ToPtr())
}

func (s Profile) Pronouns() Profile_Pronouns {
	return Profile_Pronouns(capnp.Struct(s).Uint16(0))
}

func (s Profile) SetPronouns(v Profile_Pronouns) {
	capnp.Struct(s).SetUint16(0, uint16(v))
}

// Profile_List is a list of Profile.
type Profile_List = capnp.StructList[Profile]

// NewProfile creates a new list of Profile.
func NewProfile_List(s *capnp.Segment, sz int32) (Profile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return capnp.StructList[Profile](l), err
}

// Profile_Future is a wrapper for a Profile promised by a client call.
type Profile_Future struct{ *capnp.Future }

func (p Profile_Future) Struct() (Profile, error) {
	s, err := p.Future.Struct()
	return Profile(s), err
}

func (p Profile_Future) DisplayName() util.LocalizedText_Future {
	return util.LocalizedText_Future{Future: p.Future.Field(0, nil)}
}

func (p Profile_Future) Picture() util.StaticAsset {
	return util.StaticAsset(p.Future.Field(2, nil).Client())
}

type Profile_Pronouns uint16

// Profile_Pronouns_TypeID is the unique identifier for the type Profile_Pronouns.
const Profile_Pronouns_TypeID = 0x84752dcf8539ab01

// Values of Profile_Pronouns.
const (
	Profile_Pronouns_neutral Profile_Pronouns = 0
	Profile_Pronouns_male    Profile_Pronouns = 1
	Profile_Pronouns_female  Profile_Pronouns = 2
	Profile_Pronouns_robot   Profile_Pronouns = 3
)

// String returns the enum's constant name.
func (c Profile_Pronouns) String() string {
	switch c {
	case Profile_Pronouns_neutral:
		return "neutral"
	case Profile_Pronouns_male:
		return "male"
	case Profile_Pronouns_female:
		return "female"
	case Profile_Pronouns_robot:
		return "robot"

	default:
		return ""
	}
}

// Profile_PronounsFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Profile_PronounsFromString(c string) Profile_Pronouns {
	switch c {
	case "neutral":
		return Profile_Pronouns_neutral
	case "male":
		return Profile_Pronouns_male
	case "female":
		return Profile_Pronouns_female
	case "robot":
		return Profile_Pronouns_robot

	default:
		return 0
	}
}

type Profile_Pronouns_List = capnp.EnumList[Profile_Pronouns]

func NewProfile_Pronouns_List(s *capnp.Segment, sz int32) (Profile_Pronouns_List, error) {
	return capnp.NewEnumList[Profile_Pronouns](s, sz)
}

type UserInfo capnp.Struct

// UserInfo_TypeID is the unique identifier for the type UserInfo.
const UserInfo_TypeID = 0x94b9d1efb35d11d3

func NewUserInfo(s *capnp.Segment) (UserInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7})
	return UserInfo(st), err
}

func NewRootUserInfo(s *capnp.Segment) (UserInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7})
	return UserInfo(st), err
}

func ReadRootUserInfo(msg *capnp.Message) (UserInfo, error) {
	root, err := msg.Root()
	return UserInfo(root.Struct()), err
}

func (s UserInfo) String() string {
	str, _ := text.Marshal(0x94b9d1efb35d11d3, capnp.Struct(s))
	return str
}

func (s UserInfo) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (UserInfo) DecodeFromPtr(p capnp.Ptr) UserInfo {
	return UserInfo(capnp.Struct{}.DecodeFromPtr(p))
}

func (s UserInfo) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s UserInfo) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s UserInfo) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s UserInfo) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s UserInfo) DisplayName() (util.LocalizedText, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return util.LocalizedText(p.Struct()), err
}

func (s UserInfo) HasDisplayName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s UserInfo) SetDisplayName(v util.LocalizedText) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewDisplayName sets the displayName field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s UserInfo) NewDisplayName() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(capnp.Struct(s).Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s UserInfo) PreferredHandle() (string, error) {
	p, err := capnp.Struct(s).Ptr(4)
	return p.Text(), err
}

func (s UserInfo) HasPreferredHandle() bool {
	return capnp.Struct(s).HasPtr(4)
}

func (s UserInfo) PreferredHandleBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(4)
	return p.TextBytes(), err
}

func (s UserInfo) SetPreferredHandle(v string) error {
	return capnp.Struct(s).SetText(4, v)
}

func (s UserInfo) PictureUrl() (string, error) {
	p, err := capnp.Struct(s).Ptr(5)
	return p.Text(), err
}

func (s UserInfo) HasPictureUrl() bool {
	return capnp.Struct(s).HasPtr(5)
}

func (s UserInfo) PictureUrlBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(5)
	return p.TextBytes(), err
}

func (s UserInfo) SetPictureUrl(v string) error {
	return capnp.Struct(s).SetText(5, v)
}

func (s UserInfo) Pronouns() Profile_Pronouns {
	return Profile_Pronouns(capnp.Struct(s).Uint16(0))
}

func (s UserInfo) SetPronouns(v Profile_Pronouns) {
	capnp.Struct(s).SetUint16(0, uint16(v))
}

func (s UserInfo) DeprecatedPermissionsBlob() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return []byte(p.Data()), err
}

func (s UserInfo) HasDeprecatedPermissionsBlob() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s UserInfo) SetDeprecatedPermissionsBlob(v []byte) error {
	return capnp.Struct(s).SetData(1, v)
}

func (s UserInfo) Permissions() (capnp.BitList, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return capnp.BitList(p.List()), err
}

func (s UserInfo) HasPermissions() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s UserInfo) SetPermissions(v capnp.BitList) error {
	return capnp.Struct(s).SetPtr(3, v.ToPtr())
}

// NewPermissions sets the permissions field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s UserInfo) NewPermissions(n int32) (capnp.BitList, error) {
	l, err := capnp.NewBitList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = capnp.Struct(s).SetPtr(3, l.ToPtr())
	return l, err
}

func (s UserInfo) IdentityId() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return []byte(p.Data()), err
}

func (s UserInfo) HasIdentityId() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s UserInfo) SetIdentityId(v []byte) error {
	return capnp.Struct(s).SetData(2, v)
}

func (s UserInfo) Identity() Identity {
	p, _ := capnp.Struct(s).Ptr(6)
	return Identity(p.Interface().Client())
}

func (s UserInfo) HasIdentity() bool {
	return capnp.Struct(s).HasPtr(6)
}

func (s UserInfo) SetIdentity(v Identity) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(6, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(6, in.ToPtr())
}

// UserInfo_List is a list of UserInfo.
type UserInfo_List = capnp.StructList[UserInfo]

// NewUserInfo creates a new list of UserInfo.
func NewUserInfo_List(s *capnp.Segment, sz int32) (UserInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 7}, sz)
	return capnp.StructList[UserInfo](l), err
}

// UserInfo_Future is a wrapper for a UserInfo promised by a client call.
type UserInfo_Future struct{ *capnp.Future }

func (p UserInfo_Future) Struct() (UserInfo, error) {
	s, err := p.Future.Struct()
	return UserInfo(s), err
}

func (p UserInfo_Future) DisplayName() util.LocalizedText_Future {
	return util.LocalizedText_Future{Future: p.Future.Field(0, nil)}
}

func (p UserInfo_Future) Identity() Identity {
	return Identity(p.Future.Field(6, nil).Client())
}

const schema_c822108a5c3d7d25 = "x\xda\x8c\x94O\x88\x1cE\x18\xc5\xdf\xab\xea\x9eI`" +
	"\xd7\x99\xa2\x07\xb2\x01!\x7f\x88\xa0b\xd6?9\x88\x01" +
	"\x99\xb0\x17\xddE\xa5k\xc7=\xb8\x18\xb1g\xa7v\x1d" +
	"\xe8\x99\x1e\xab{0\xab\xc8J\xc8\x1eT\x02\x1e<\xe8" +
	"\xd5\x8b\x04\xf1 \"\xba\xe6\"\x91\xa09\x84$\xeei" +
	"AA\xbd$\x82\x12EO\"\xb4T\xcf\xbf%b\xc8" +
	"i\x86\xaf^\xbf\xef\xeb\xef\xfd\xba\x1e\xda'Nx\x0f" +
	"O_/A\xe8\xd0/\xe5\xfc\xe8\xb1\xcd+G\xfbg" +
	"\xa0\xf6\x89\xfc\xe6?\xb4O\x7f}u\x1b\xe0\xb1\xe7\xa4" +
	"``d\x19\x08\"\xb9\x06\xe6\xdb\xea\xe4\xa77\xafm" +
	"\xbd\x0b]!\xf3{^\x7f\xfc\xf9\xb7\xaa\x87\xbf\x85_" +
	"v\x92My!8\xeb\xc4\xc7\xde\x94\xdf\x10\xccg~" +
	"\xd0\xe7^}\xef\xccWP\x159\x11\x83\xc1\xa6\x7f!" +
	"8\xeb\xef\x03\x82\xf7\xfd'\x82\x8b~\x19\xc8;\x0f^" +
	"Z\xfa\xc2\xda\xcbP3\x04|:\x9f\x8f\xfde\x82\xc1" +
	"\x96_\x07'\x83\xdd\xd2\xbb\x18o\xc7\xff2\xf8\xa9p" +
	"\xfc\xcd\xbf\x0e\xe6O\xed\xf4?\xfbp\xfd\xe8\x9f\x037" +
	"\xcf\x99].-\x12^~\xe3\xfc\xf6'W\x16\xc3\xe2" +
	"d2\xe1\xa0\xdfV\xe90\x83K%gx\xb1T\xc7" +
	";y\xbbe\xbaY;[\x97\xb3+Q\xaf\xdb;\x1e" +
	"\xdad\xb5\x1d\x9b\xd9\xd0&\xdd\xa4\xdfM\x11\x92\xbaJ" +
	"\x01\xa8\xfb\xe6\x00R\x1d\xba\x1f\xa0P\xfb\x8f\x03\x94J" +
	"=\x02ltM?\xb3Q\\\xe9D\xb1\xa9\xaf\x1a\xf7" +
	"s\xc0&\xcd$\x1b\xdb\x8b\xa1\xfdRj\xec|w5" +
	")l\x0fJ\x0f\xf0\x08\xa8kM@_\x95\xd4\xdf\x0b" +
	"*\x8f5\xba\xe2\xcew\x80\xfeQR\xff*\xa8J\xa2" +
	"V\x0c\xf1\xcb2\xa0oH\xea\xbf\x04\x95/k\x94\x80" +
	"\xfa\xc3=\xfe\xbbdc\x8a\x82\x8a^\x8d\x1e\x10\xec\xe5" +
	"i\xa0\xb1\x87\x92\x8d\x1a\x05)k\xf4\x81@q\x01h" +
	"T]\xf9n'\x17~\x8d% \xd8\xcfe\xa01\xe3" +
	"\xeaG\\\xbd\\\xaa\xb9\x95\x05\x87\x0a\xfdAW\x7f\x80" +
	"\x82y\xab\x9d\xf6\xe2h\xfd\x19\x94\xa3\x8ea5\x9f\xff" +
	"\xf9\xb5G\xed\xe7'\xdfv\xcb\xa9\x82y\xcb\xf4\xacY" +
	"\x892aZ\xa1\xb1\x9dv\x9a\xb6\x93n:\x17'l" +
	"r\x1a\x82\xd3\xe0x-\x90\xf3\xadq\xb17T\xa3\x9c" +
	"tS\xde\x05\x86\x92$\x84\xfb\x9b\xf7\xacY5\xd6\x1a" +
	"\xb6\x9e\x8c\xba\xad\xd8\x80S\x10\x9c*N\x86I\x01\xac" +
	"LH\x07Yq\xa7\xed\x95\xaco\xcd\x12\xa4\x8d\xc7\x8f" +
	"\x8c\xdb\x03T\x13D\\\xba\xbbNG\x99\xcd\x8f\xd4\xda" +
	"#w\xd1\xc5f\x1e&\xaf\x18\xdbLN\xa1\xfcl\xb4" +
	"\xa6=\xe9\x03c09\xe2]\xa9e\x08\xb5\xb7\x9c\xaf" +
	"\x99\xac\xc0\x0b26'\x18r\xd2\xca\xbb\xa5\xd5\xecH" +
	"\x1a\x9b#\x8b&\xed\xc72K\xb57\xc6ez\x0e\xd0" +
	"{$uMp\xa37\xd0\xb1\xba\xfb\xc3.\x92\x18\xb9" +
	"s\xc4v} -\xdecr%p!\x0f';\xd4" +
	"\xd5q\x9b\xc8a\xf5\xa2\xa4\x8e\x1dUC*\xdb\xa7\x01" +
	"\xfd\x92\xa4\xce\x1c;C*_v\x03\xc5\x92\xfa\xd4\x80" +
	"3\x07e\x7f\x01\xd0\x99\xa4~\xe3\x0e\xa0\xf9\xdfx7" +
	"\x86\x01R\xe5\xf5\x0f\xce\x1dh\xbep\xfe\xefQN\xb7" +
	"\x0f\xfeNV\x1bF6\xea0\xfd\xcf%0\xd6\x8e\xf2" +
	"u\xf1\xba/vW\x04n7S\x92\xfa^q{r" +
	"\xff\x0d\x00\x00\xff\xff\xee\x14\x7fa"

func init() {
	schemas.Register(schema_c822108a5c3d7d25,
		0x84752dcf8539ab01,
		0x94b9d1efb35d11d3,
		0xc084987aa951dd18,
		0xcd7272b855c92f6d,
		0xd3d0c34d7201fcef,
		0xf32d79a7b575d94c,
		0xf35052cfb1d3bbe8)
}
