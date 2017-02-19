package powerbox

// AUTO GENERATED - DO NOT EDIT

import (
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type PowerboxDescriptor struct{ capnp.Struct }

// PowerboxDescriptor_TypeID is the unique identifier for the type PowerboxDescriptor.
const PowerboxDescriptor_TypeID = 0xedcf0fa3bfc71c58

func NewPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor{st}, err
}

func NewRootPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor{st}, err
}

func ReadRootPowerboxDescriptor(msg *capnp.Message) (PowerboxDescriptor, error) {
	root, err := msg.RootPtr()
	return PowerboxDescriptor{root.Struct()}, err
}

func (s PowerboxDescriptor) String() string {
	str, _ := text.Marshal(0xedcf0fa3bfc71c58, s.Struct)
	return str
}

func (s PowerboxDescriptor) Tags() (PowerboxDescriptor_Tag_List, error) {
	p, err := s.Struct.Ptr(0)
	return PowerboxDescriptor_Tag_List{List: p.List()}, err
}

func (s PowerboxDescriptor) HasTags() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDescriptor) SetTags(v PowerboxDescriptor_Tag_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTags sets the tags field to a newly
// allocated PowerboxDescriptor_Tag_List, preferring placement in s's segment.
func (s PowerboxDescriptor) NewTags(n int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := NewPowerboxDescriptor_Tag_List(s.Struct.Segment(), n)
	if err != nil {
		return PowerboxDescriptor_Tag_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s PowerboxDescriptor) Quality() PowerboxDescriptor_MatchQuality {
	return PowerboxDescriptor_MatchQuality(s.Struct.Uint16(0))
}

func (s PowerboxDescriptor) SetQuality(v PowerboxDescriptor_MatchQuality) {
	s.Struct.SetUint16(0, uint16(v))
}

// PowerboxDescriptor_List is a list of PowerboxDescriptor.
type PowerboxDescriptor_List struct{ capnp.List }

// NewPowerboxDescriptor creates a new list of PowerboxDescriptor.
func NewPowerboxDescriptor_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PowerboxDescriptor_List{l}, err
}

func (s PowerboxDescriptor_List) At(i int) PowerboxDescriptor {
	return PowerboxDescriptor{s.List.Struct(i)}
}

func (s PowerboxDescriptor_List) Set(i int, v PowerboxDescriptor) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDescriptor_Promise is a wrapper for a PowerboxDescriptor promised by a client call.
type PowerboxDescriptor_Promise struct{ *capnp.Pipeline }

func (p PowerboxDescriptor_Promise) Struct() (PowerboxDescriptor, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDescriptor{s}, err
}

type PowerboxDescriptor_Tag struct{ capnp.Struct }

// PowerboxDescriptor_Tag_TypeID is the unique identifier for the type PowerboxDescriptor_Tag.
const PowerboxDescriptor_Tag_TypeID = 0xbe3d16a8df03c418

func NewPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor_Tag{st}, err
}

func NewRootPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor_Tag{st}, err
}

func ReadRootPowerboxDescriptor_Tag(msg *capnp.Message) (PowerboxDescriptor_Tag, error) {
	root, err := msg.RootPtr()
	return PowerboxDescriptor_Tag{root.Struct()}, err
}

func (s PowerboxDescriptor_Tag) String() string {
	str, _ := text.Marshal(0xbe3d16a8df03c418, s.Struct)
	return str
}

func (s PowerboxDescriptor_Tag) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s PowerboxDescriptor_Tag) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s PowerboxDescriptor_Tag) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s PowerboxDescriptor_Tag) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDescriptor_Tag) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s PowerboxDescriptor_Tag) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s PowerboxDescriptor_Tag) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// PowerboxDescriptor_Tag_List is a list of PowerboxDescriptor_Tag.
type PowerboxDescriptor_Tag_List struct{ capnp.List }

// NewPowerboxDescriptor_Tag creates a new list of PowerboxDescriptor_Tag.
func NewPowerboxDescriptor_Tag_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PowerboxDescriptor_Tag_List{l}, err
}

func (s PowerboxDescriptor_Tag_List) At(i int) PowerboxDescriptor_Tag {
	return PowerboxDescriptor_Tag{s.List.Struct(i)}
}

func (s PowerboxDescriptor_Tag_List) Set(i int, v PowerboxDescriptor_Tag) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDescriptor_Tag_Promise is a wrapper for a PowerboxDescriptor_Tag promised by a client call.
type PowerboxDescriptor_Tag_Promise struct{ *capnp.Pipeline }

func (p PowerboxDescriptor_Tag_Promise) Struct() (PowerboxDescriptor_Tag, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDescriptor_Tag{s}, err
}

func (p PowerboxDescriptor_Tag_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type PowerboxDescriptor_MatchQuality uint16

// Values of PowerboxDescriptor_MatchQuality.
const (
	PowerboxDescriptor_MatchQuality_preferred    PowerboxDescriptor_MatchQuality = 1
	PowerboxDescriptor_MatchQuality_acceptable   PowerboxDescriptor_MatchQuality = 0
	PowerboxDescriptor_MatchQuality_unacceptable PowerboxDescriptor_MatchQuality = 2
)

// String returns the enum's constant name.
func (c PowerboxDescriptor_MatchQuality) String() string {
	switch c {
	case PowerboxDescriptor_MatchQuality_preferred:
		return "preferred"
	case PowerboxDescriptor_MatchQuality_acceptable:
		return "acceptable"
	case PowerboxDescriptor_MatchQuality_unacceptable:
		return "unacceptable"

	default:
		return ""
	}
}

// PowerboxDescriptor_MatchQualityFromString returns the enum value with a name,
// or the zero value if there's no such value.
func PowerboxDescriptor_MatchQualityFromString(c string) PowerboxDescriptor_MatchQuality {
	switch c {
	case "preferred":
		return PowerboxDescriptor_MatchQuality_preferred
	case "acceptable":
		return PowerboxDescriptor_MatchQuality_acceptable
	case "unacceptable":
		return PowerboxDescriptor_MatchQuality_unacceptable

	default:
		return 0
	}
}

type PowerboxDescriptor_MatchQuality_List struct{ capnp.List }

func NewPowerboxDescriptor_MatchQuality_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_MatchQuality_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return PowerboxDescriptor_MatchQuality_List{l.List}, err
}

func (l PowerboxDescriptor_MatchQuality_List) At(i int) PowerboxDescriptor_MatchQuality {
	ul := capnp.UInt16List{List: l.List}
	return PowerboxDescriptor_MatchQuality(ul.At(i))
}

func (l PowerboxDescriptor_MatchQuality_List) Set(i int, v PowerboxDescriptor_MatchQuality) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type PowerboxDisplayInfo struct{ capnp.Struct }

// PowerboxDisplayInfo_TypeID is the unique identifier for the type PowerboxDisplayInfo.
const PowerboxDisplayInfo_TypeID = 0xa553a209bee32bec

func NewPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return PowerboxDisplayInfo{st}, err
}

func NewRootPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return PowerboxDisplayInfo{st}, err
}

func ReadRootPowerboxDisplayInfo(msg *capnp.Message) (PowerboxDisplayInfo, error) {
	root, err := msg.RootPtr()
	return PowerboxDisplayInfo{root.Struct()}, err
}

func (s PowerboxDisplayInfo) String() string {
	str, _ := text.Marshal(0xa553a209bee32bec, s.Struct)
	return str
}

func (s PowerboxDisplayInfo) Title() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(0)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PowerboxDisplayInfo) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetTitle(v util.LocalizedText) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) VerbPhrase() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(1)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PowerboxDisplayInfo) HasVerbPhrase() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetVerbPhrase(v util.LocalizedText) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewVerbPhrase() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) Description() (util.LocalizedText, error) {
	p, err := s.Struct.Ptr(2)
	return util.LocalizedText{Struct: p.Struct()}, err
}

func (s PowerboxDisplayInfo) HasDescription() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s PowerboxDisplayInfo) SetDescription(v util.LocalizedText) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(s.Struct.Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// PowerboxDisplayInfo_List is a list of PowerboxDisplayInfo.
type PowerboxDisplayInfo_List struct{ capnp.List }

// NewPowerboxDisplayInfo creates a new list of PowerboxDisplayInfo.
func NewPowerboxDisplayInfo_List(s *capnp.Segment, sz int32) (PowerboxDisplayInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return PowerboxDisplayInfo_List{l}, err
}

func (s PowerboxDisplayInfo_List) At(i int) PowerboxDisplayInfo {
	return PowerboxDisplayInfo{s.List.Struct(i)}
}

func (s PowerboxDisplayInfo_List) Set(i int, v PowerboxDisplayInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// PowerboxDisplayInfo_Promise is a wrapper for a PowerboxDisplayInfo promised by a client call.
type PowerboxDisplayInfo_Promise struct{ *capnp.Pipeline }

func (p PowerboxDisplayInfo_Promise) Struct() (PowerboxDisplayInfo, error) {
	s, err := p.Pipeline.Struct()
	return PowerboxDisplayInfo{s}, err
}

func (p PowerboxDisplayInfo_Promise) Title() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p PowerboxDisplayInfo_Promise) VerbPhrase() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p PowerboxDisplayInfo_Promise) Description() util.LocalizedText_Promise {
	return util.LocalizedText_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

const schema_f6c200ab14cd53e4 = "x\xdat\x92Kk\x13Q\x1c\xc5\xcf\xb971\x0d\xa8" +
	"\xd3\xdb\xa9\xa2A\x09\x04\xc5\x17\x86>Q\x84b\x90\x16" +
	"\x0c(\xe6\xda,\xa4\xd4\xc5M2M\xa7\x84d\x9cL" +
	"\x1a\xa3\x8b~\x00W~\x05E\x10\\\xb8t!T\x10" +
	"\xc1\x95\x88~\x00\x8b/p\xa1[\x17nF&\x0f\x13" +
	"\xa4]\xce\xff\x9c\xb9\xe7\xfc\x7f\xfc\xa7\xe6\x99\x13\xd3\xf1" +
	"\xad\x18\xa0\xe7\xe2\xfb\xc2\x9f\xe7\xben'\x1f-?\x81" +
	"\xb2\xb8\xf3m\xf9\xdd\xe4\xb3\xd7\xbf\xe32\x01\xcc\x9e\x14" +
	")\xda\xf3\"\x01\xd8\xd3\xe29\x18~\xfa\xb1\xfaq\xe9" +
	"O\xea%TF\x84\xb7\x8e\xbd}\xf5\xd8z\xff\x0b\xe0" +
	"\xec\x07q\x8f\xf6\xf7\xae\xf3\xb3\xb8\x01\x86G\xde\xc8\x9d" +
	"\xa7\x87\x17\xb6\xa13\xe4\xd0\x1ag\xf4.\xe5\x0cm\x15" +
	"E\xd8\x07d\x1b#\xba\xb68\xec\xd0\xf5\xbar\x82v" +
	"G\x9e\x02\xec\x87\xb2\x8dJ\xe85\xda\x8e_j\xdc\x95" +
	"\xd9\xb2\xf1\xea\xde\xa5B\xff{\xd1mz5\xd3\xc9\xd7" +
	"\xe5Z\xa3@\xea\xfd2\x06\xc4\x08\xa8\xa5\x19@\xe7$" +
	"\xf55AEN2\x1a\xe6W\x00}UR\x17\x05\x95" +
	"\x10\x93\x14\x80\xd2%@\x17$\xf5\xaa`:p\x83\x9a" +
	"\xc3\xf10\xff\xe5\xfe\x05\xff\xc5\xed\x07\x009\x0e\x86\x9b" +
	"\x8e_*\xac\xfb\x06\xb2\xb9\x9b\\q\x9ae\xdf\xf5\x02" +
	"$\xdcF}\x17}\xd0?\xf6\x7f\xff\xfe\x7f\x0d?{" +
	"\xdd\x04\xe5u}\xb9ejn\xd0\xe9\xeeBA\xaa\xe3" +
	"+\xa0:z\x13\xa0P\x876\x80\xd0\x94\xcb\x8e\x17\x98" +
	"\x12d\xcd\x09=\xdfYs|\xdf\x01+a\xab\xde\x93" +
	"`\x99R$\xed\x85l\x10\x99\xf6\xb3ES\x8d\x92\xc6" +
	"\xfeQ;\x93\x02\xf4\x09I=%8\x80v>\"y" +
	"ZR\xcf\x09J\xb7\xc2$\x04\x93`z\xd3\xd4Z\x0e" +
	"' 81\xb2\xe1\x9eq\xf4\xf5\x18GnD%3" +
	"#\xa7\x15\xdfH\x14M5\xecAh\x19X\x11\x85\xd1" +
	"bg\xfb\xc5r\xfdb\xa4Z\xb8\x02\xe8\x8b\x92zQ" +
	"\xd0\x0aL\xb5\xc9\x83`A\x92\xe3\xc3\x140\x1an\xdd" +
	"\xe9a\xa55\x8c\x04i\x81\x7f\x03\x00\x00\xff\xff\xae&" +
	"\xda\xdd"

func init() {
	schemas.Register(schema_f6c200ab14cd53e4,
		0xa553a209bee32bec,
		0xbb1afa45d25ce8de,
		0xbe3d16a8df03c418,
		0xedcf0fa3bfc71c58)
}
