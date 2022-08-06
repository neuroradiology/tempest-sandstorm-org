// Code generated by capnpc-go. DO NOT EDIT.

package appidreplacementstest

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	appidreplacements "zenhack.net/go/sandstorm/capnp/appidreplacements"
)

// Constants defined in appid-replacements-test.capnp.
const (
	TestIds_unusedApp = "6pm4ujs8f5f5wugc87uhuhvhs57he09u10rv8qd2jgdup9f69yzh"
	TestIds_app1      = "vjvekechd398fn1t1kn1dgdnmaekqq9jkjv3zsgzymc4z913ref0"
	TestIds_app2      = "wq95qmutckc0yfmecv0ky96cqxgp156up8sv81yxvmery58q87jh"
	TestIds_app3      = "302t6c6kf8hjer1kh3469d4ch10d936g7wkwtxcs12pwh9u5axqh"
	TestIds_app4      = "5ddk4uqnstnsqvp3thc2tyed41c7wp4x5ygt20zrh3u0tnv5jqd0"
	TestIds_app5      = "jkz6yhywhp4uk5sgkc5ugwnee57a5h5wu4rfmujtahny5r8g3ych"
	TestIds_app6      = "adk6syfj42fpp3xhgqrrheqgfxkhaw8e1t11vug44ys6pzaxqugh"
	TestIds_unusedPkg = "7300e3448dd2b53e075d0a8481c2bc06"
	TestIds_pkg1      = "b5bb9d8014a0f9b1d61e21e796d78dcc"
	TestIds_pkg2      = "8613a11b8ac365cb36775a6b8ca6176c"
	TestIds_pkg3      = "77c4f45aee83e376d31a5680cdb841a2"
)

// Constants defined in appid-replacements-test.capnp.
var (
	TestAppIdReplacementList = appidreplacements.AppIdReplacement_List(capnp.MustUnmarshalRoot(x_bee445adfb01a777[0:712]).List())
)

type TestIds capnp.Struct

// TestIds_TypeID is the unique identifier for the type TestIds.
const TestIds_TypeID = 0x9440399ec56efc9b

func NewTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TestIds(st), err
}

func NewRootTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TestIds(st), err
}

func ReadRootTestIds(msg *capnp.Message) (TestIds, error) {
	root, err := msg.Root()
	return TestIds(root.Struct()), err
}

func (s TestIds) String() string {
	str, _ := text.Marshal(0x9440399ec56efc9b, capnp.Struct(s))
	return str
}

func (s TestIds) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (TestIds) DecodeFromPtr(p capnp.Ptr) TestIds {
	return TestIds(capnp.Struct{}.DecodeFromPtr(p))
}

func (s TestIds) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s TestIds) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s TestIds) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s TestIds) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// TestIds_List is a list of TestIds.
type TestIds_List = capnp.StructList[TestIds]

// NewTestIds creates a new list of TestIds.
func NewTestIds_List(s *capnp.Segment, sz int32) (TestIds_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[TestIds](l), err
}

// TestIds_Future is a wrapper for a TestIds promised by a client call.
type TestIds_Future struct{ *capnp.Future }

func (p TestIds_Future) Struct() (TestIds, error) {
	s, err := p.Future.Struct()
	return TestIds(s), err
}

const schema_bee445adfb01a777 = "x\xda\xac\xd5[\x88\x1b\xd5\x1f\x07\xf0sv2\x9bR" +
	"\xbaM\xf7\x1f\xfa\x7f*],(\x12\xdc:g\xce\x9c" +
	"\xb9\x88\xd2\xdd\x87E\x16-\xecn*\xd4R/\x93\x99" +
	"33\xcd\x98\xe1L\xe6\x92L^J\xabE\x0b\x0a\x0a" +
	"E\xbc\xa0TP\x8a\xa9\x8b\xf6\xa1\xb4\xd4V|\xb0J" +
	"\xa1>\x88\x8a\x17((\xf4I\xf4AD\xad\xb4\x109" +
	"Iuc\x82\xba\x09\xbe\xfc\x18\xc2\x87\xf3\xcd\xf9\x9d\xdf" +
	"\x99\x91\x1e\x99\x98\xcb\xa1\xa9]y0\xb1\xbcG\x9c\xec" +
	"\xfc\xf2\xee\xa1\x07\xa2\x83W\x1e\x07\xd3\xb7\xe5:/\xdf" +
	"\x08.\xbej\xcc\x1d\x03\x00\xe2#\xc2>X|Q\xc8" +
	"\x03P>&\x08\xb0|\\\x98\x80\x00\xac\x91\xe9\xed\xb0" +
	"\xd38\x01\xaf\xaf.\\}\x0f\xe4\xf2\x00\xe0UA\x86" +
	"\xc5\x0b\xc2\x19 t~\x9e9y\xf7\xb5S\xf9\xe7\x87" +
	"V\xbd5\xb7\x0f\x16\x0d\xce\xcbJN\x80\xe5\xb9\\w" +
	"\xd5\xd3\xf7e\xc2\x91\x97\x84\xd7\x87\xfc\xee\xdcaX4" +
	"\xbb~?\xf7^\xcf?H\x827\xd2\xa9;\xdf\x1e\xf2" +
	"\x09_\xffh\xd7?\xc1\xfd\xb3=\xef\xad~\xd7<s" +
	"\xfe\xc9a\xff\x0a\xf7\xa7\xba\xfe-\xee\xcf\xf6\xfc\xc9)" +
	"if\xf5\x85w\xce\x0f\xf9\x8b\xdc\x7f\xd5\xf5\x9fq\xff" +
	"M\xcf\x7f\xf1\xf4b4\xe3\x1f~\x7f\xc8\xff\xc0=\x14" +
	"\xb9\xbf\xc1\xfd\x06\xb1\xeb7|\xfd\xdc\xb6\xff=\xfa\xf9" +
	"\x87C~J<\x0c\x8b\xb7t\xfd6Q\x80\xe5\xdb{" +
	"~\xc7\xd5\xef7_\xdf\xda\xb9\x04\xa6\xb7\xe7\xd6\xba\x0e" +
	" \x9e\x15\xcf\xc1\xe2B\xd7\xcfq\xbf\xbf\xe7\x7f\xdaq" +
	"yO\"\x1c\xfath\xfd7'\xf7\xc1\xe2\x85I\xee" +
	"\xcfN\x0a\xb0\xfc\xc1d\xd7?\xfc\x89y\xb0\xfc\xd1\xbd" +
	"\xbf\x0e\xf9K\xdc_\xe9\xfa/\xb9\xbf\xda\xf3\x1b|v" +
	"bo\xaa]\x1b\xf2?r/\xe6\xf3\x00\xac\xe4\x05X" +
	"\xde\x94\x9f\x80`o\xc7d\xec\x80=[\xa79\xf6\x98" +
	"i\xd1\x1a\x0d\xe2h6\xa6Q\xbc\xd32Y\xc0\xee\xda" +
	"C\xa3x\xd1\x8ev\x9aL`d\x09B\xb8\x09L\xc0" +
	"M\x00 \xd8\x86\x9d\xaa\xdfR3/k\x88\x1eS\x12" +
	"\x9fD\xaeo\x91\xc4m\x04\x94\x12\xcd$\x1ei$J" +
	"\xdd\xa9%\xd5\xd8\xf4\x82\x8c\xd4u\xb7\x803\xcb\xfb3" +
	"R\xf8\xbb\xc8\x99n\xe6\xf2\x1d\x10\xf6M\xdf\xf2J\xdf" +
	"h-\x97\xfa\xfa\xb2\xbb\xd47\xd4\x8b\xa5\xbe\x89Z(" +
	"\xf5]\xa2\xf9R\xdf\xec\xdcS\xea;hc\xa5\xafk" +
	"F\xa9odH\xa9\xef\xbcP\xa9\x93\x04ID\xedy" +
	"\x06 +\x98\x8c!^d^0/\x0a/\x84\x17\xf5" +
	"\xa6\\\xf2\x01t\x0b\xccw\x11/2/x\x94\x96\xe3" +
	"\xc1\x96cI\x8eUK\xf5EG\xf7\xaa\xb4\x8e|\x0f" +
	"+\xaaa+\x96\x87$\xdb\xc0\xaa\xab5\xfcF\xdc\xb4" +
	"\"$\xb3\x86g$\xa4`6\xc3\xb5\x96\x8b\xff\x1c\xf9" +
	"\xc7\xf6\x18\x18\xccUYMI\xaa\x91.:\xc4!\x8d" +
	"\xc4\xb5t-\xf1\x12/\xf5\"\xa2yT2\x12$\xd5" +
	"S=\xb4\xe5\xaak'\xccp\xd4\x82\x91\xb5\xbcQ\xb6" +
	"\x8a\x06#\xd3jJ}jy\xa2\x8d\x0d\xdd\x09P\x8c" +
	"\xfc\x00\xd9\xae\x1d\xd4L\xea\x87\xa1Q\xf5\xab)nE" +
	"n+\xabYJ\xcb@\xb8P\xa7\x8e4J\xa42\x18" +
	"Il\xdbW\x920\x10\xa38\x88\xc2\x94\xe1\xd8\xb3\xe4" +
	"8\xa3\xb6\x82,\xad\xc1\x94&\xc9\xdcX\x96Zu\x0f" +
	"'R\x1c\xa4\xa4P\x0d\xed\x91\"\xd5\xc1H\xd3\xf6\xd5" +
	"(s\xaa\xa2\";\x8c\xe1\xa6\xe7\x86\xf5\xbaGC\xd7" +
	"i\xfa\x9e\xd9\xd0)\x8a\x11J\x13WQ\xb2He-" +
	"\xb3Y\x08\x13w\xdd\x8de\xbe\xe0\xca\x7f\x89\xdc\x08;" +
	"\xba\x8a\xb0\x89PE\xd0M\x0b\xab\xc4\xaa`U\xd3\x88" +
	"\xa9Vt\xcbT\x91\xa6Z\x00\x8c60K\xbe\x0b\x06" +
	"C4,I\x14+\x8a\xa0\xdb\xb6\\!\x98J\x1a\xb1" +
	"%SWtd\xc9\x15KR\xff=\x84?\xce3\xb6" +
	"h\xaf\xd0\x9b`W\x10\xdf\x7f \x8ay\xd6f\x00\x97" +
	"\x04\x08\xb7tZ\xc7\xcf5\xbd\xcb\xcf<\x05\x00\xff\x11" +
	"LCw~\x8b\x80f\xdb\x10-\xb4\xe1\xf4C\xffG" +
	"G\xdb\x10\xbd\xd6\x86\x00\xa0\xd3m\x88>\xee>}\xdb" +
	"\x86\xe8\xb76,N\xc3\xcd\xe3\xceY#4HXK" +
	"b\xd1\xf2-)sj\xd4J%?3T+l\xba" +
	"\x0c\x115az\x94\xea(k\xa65Z\xcf\x88\x1e\x16" +
	"t\xad\xea!q#D[7\xc2N\x85T*\x86\xad" +
	"K\x02RL\xc91*\xc8V\x11\x95\x11\xd5\x0c\xd5\xd6" +
	"t\xdb\xe2\x87\xb0\x9e\x83\x1a\xefo\x8c\xfb\x1e\x19\xf3\x82" +
	"\x8c\xf9\xa1\x18\xf7\xfb2\xd6\x95\xeaM\xadf)\x8eB" +
	"L\x81R\x1dS\xac\xa96F&Qu\xc9\xb2+\xba" +
	"\x82L\xb9oj\xd7q\xf5\xf0\xd0\xad\xf8\xef\xd6\xe7o" +
	"\x13y\xf0m2\xe64\x8c\xb0%4\xb8\xa5u\x0c\xf2" +
	"\xef\x01\x00\x00\xff\xff\xeaj\x85~"

func init() {
	schemas.Register(schema_bee445adfb01a777,
		0x83dd7f735581bbf6,
		0x9440399ec56efc9b,
		0x9607b1f83cab1ff5,
		0xa4039a8503794bb5,
		0xaf2f0d76a56e3559,
		0xaf87bcb778eaad68,
		0xbcb098ad1f300dab,
		0xc0826b1f73498cd7,
		0xc6d560121c91da08,
		0xc9ff15fb0eece422,
		0xd381037554cc22f3,
		0xf747c7537f61d15e,
		0xf8377658a7706b08)
}

var x_bee445adfb01a777 = []byte{
	0, 0, 0, 0, 88, 0, 0, 0,
	1, 0, 0, 0, 103, 0, 0, 0,
	16, 0, 0, 0, 0, 0, 3, 0,
	45, 0, 0, 0, 170, 1, 0, 0,
	69, 0, 0, 0, 170, 1, 0, 0,
	93, 0, 0, 0, 22, 0, 0, 0,
	137, 0, 0, 0, 170, 1, 0, 0,
	161, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	181, 0, 0, 0, 170, 1, 0, 0,
	205, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	225, 0, 0, 0, 170, 1, 0, 0,
	249, 0, 0, 0, 170, 1, 0, 0,
	17, 1, 0, 0, 14, 0, 0, 0,
	118, 106, 118, 101, 107, 101, 99, 104,
	100, 51, 57, 56, 102, 110, 49, 116,
	49, 107, 110, 49, 100, 103, 100, 110,
	109, 97, 101, 107, 113, 113, 57, 106,
	107, 106, 118, 51, 122, 115, 103, 122,
	121, 109, 99, 52, 122, 57, 49, 51,
	114, 101, 102, 48, 0, 0, 0, 0,
	119, 113, 57, 53, 113, 109, 117, 116,
	99, 107, 99, 48, 121, 102, 109, 101,
	99, 118, 48, 107, 121, 57, 54, 99,
	113, 120, 103, 112, 49, 53, 54, 117,
	112, 56, 115, 118, 56, 49, 121, 120,
	118, 109, 101, 114, 121, 53, 56, 113,
	56, 55, 106, 104, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 1, 0, 0,
	21, 0, 0, 0, 10, 1, 0, 0,
	98, 53, 98, 98, 57, 100, 56, 48,
	49, 52, 97, 48, 102, 57, 98, 49,
	100, 54, 49, 101, 50, 49, 101, 55,
	57, 54, 100, 55, 56, 100, 99, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	56, 54, 49, 51, 97, 49, 49, 98,
	56, 97, 99, 51, 54, 53, 99, 98,
	51, 54, 55, 55, 53, 97, 54, 98,
	56, 99, 97, 54, 49, 55, 54, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	119, 113, 57, 53, 113, 109, 117, 116,
	99, 107, 99, 48, 121, 102, 109, 101,
	99, 118, 48, 107, 121, 57, 54, 99,
	113, 120, 103, 112, 49, 53, 54, 117,
	112, 56, 115, 118, 56, 49, 121, 120,
	118, 109, 101, 114, 121, 53, 56, 113,
	56, 55, 106, 104, 0, 0, 0, 0,
	51, 48, 50, 116, 54, 99, 54, 107,
	102, 56, 104, 106, 101, 114, 49, 107,
	104, 51, 52, 54, 57, 100, 52, 99,
	104, 49, 48, 100, 57, 51, 54, 103,
	55, 119, 107, 119, 116, 120, 99, 115,
	49, 50, 112, 119, 104, 57, 117, 53,
	97, 120, 113, 104, 0, 0, 0, 0,
	53, 100, 100, 107, 52, 117, 113, 110,
	115, 116, 110, 115, 113, 118, 112, 51,
	116, 104, 99, 50, 116, 121, 101, 100,
	52, 49, 99, 55, 119, 112, 52, 120,
	53, 121, 103, 116, 50, 48, 122, 114,
	104, 51, 117, 48, 116, 110, 118, 53,
	106, 113, 100, 48, 0, 0, 0, 0,
	106, 107, 122, 54, 121, 104, 121, 119,
	104, 112, 52, 117, 107, 53, 115, 103,
	107, 99, 53, 117, 103, 119, 110, 101,
	101, 53, 55, 97, 53, 104, 53, 119,
	117, 52, 114, 102, 109, 117, 106, 116,
	97, 104, 110, 121, 53, 114, 56, 103,
	51, 121, 99, 104, 0, 0, 0, 0,
	106, 107, 122, 54, 121, 104, 121, 119,
	104, 112, 52, 117, 107, 53, 115, 103,
	107, 99, 53, 117, 103, 119, 110, 101,
	101, 53, 55, 97, 53, 104, 53, 119,
	117, 52, 114, 102, 109, 117, 106, 116,
	97, 104, 110, 121, 53, 114, 56, 103,
	51, 121, 99, 104, 0, 0, 0, 0,
	97, 100, 107, 54, 115, 121, 102, 106,
	52, 50, 102, 112, 112, 51, 120, 104,
	103, 113, 114, 114, 104, 101, 113, 103,
	102, 120, 107, 104, 97, 119, 56, 101,
	49, 116, 49, 49, 118, 117, 103, 52,
	52, 121, 115, 54, 112, 122, 97, 120,
	113, 117, 103, 104, 0, 0, 0, 0,
	1, 0, 0, 0, 10, 1, 0, 0,
	55, 55, 99, 52, 102, 52, 53, 97,
	101, 101, 56, 51, 101, 51, 55, 54,
	100, 51, 49, 97, 53, 54, 56, 48,
	99, 100, 98, 56, 52, 49, 97, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
}
