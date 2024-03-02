package websession

import (
	"net/http"
	"strings"

	"capnproto.org/go/capnp/v3"
	"zenhack.net/go/util"

	websession "sandstorm.org/go/tempest/capnp/web-session"
)

var (
	// ContextHeaderFilter is a header filter for WebSession.Context.additionalHeaders
	ContextHeaderFilter = util.Must(ParseHeaderWhitelist(websession.Context_headerWhitelist))

	// ResponseHeaderFilter is a header filter for WebSession.Response.additionalHeaders
	ResponseHeaderFilter = util.Must(ParseHeaderWhitelist(websession.Response_headerWhitelist))
)

// A HeaderFilter filters headers based on an allow list.
type HeaderFilter struct {
	// Headers matching keys in exact are allowed.
	exact map[string]struct{}

	// Headers starting with elements of prefixes are allowed.
	prefixes []string
}

// NewHeaderFilter a new empty header filter (i.e. one that drops
// all headers).
func NewHeaderFilter() HeaderFilter {
	return HeaderFilter{
		exact: make(map[string]struct{}),
	}
}

// ParseHeaderWhiteList parses the `headerWhitelist` fields defined
// in web-session.capnp, and returns a corresponding HeaderFilter.
func ParseHeaderWhitelist(l capnp.TextList) (HeaderFilter, error) {
	ret := NewHeaderFilter()
	for i := 0; i < l.Len(); i++ {
		rule, err := l.At(i)
		if err != nil {
			return HeaderFilter{}, err
		}
		ret.AddAllow(http.CanonicalHeaderKey(rule))
	}
	return ret, nil
}

// AddAllow adds an allow rule to the filter. If the rule ends
// in '*' it is a prefix, otherwise it is an exact value.
func (hf *HeaderFilter) AddAllow(rule string) {
	if strings.HasSuffix(rule, "*") {
		hf.prefixes = append(hf.prefixes, rule[:len(rule)-1])
		return
	}
	hf.exact[rule] = struct{}{}
}

// Allows returns true iff the header filter allows the named header.
func (hf HeaderFilter) Allows(name string) bool {
	if _, ok := hf.exact[name]; ok {
		return true
	}
	for _, p := range hf.prefixes {
		if strings.HasPrefix(name, p) {
			return true
		}
	}

	return false
}

// Copy copies all headers allowed by the filter
func (hf HeaderFilter) Copy(dst, src http.Header) {
	for k, v := range src {
		if hf.Allows(k) {
			dst[k] = v
		}
	}
}
