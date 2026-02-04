package http

import (
	"net/http"
	"net/textproto"

	"github.com/chaos-io/core/go/chaos/core"
)

const HeaderTypeName = "Header"
const HeaderTypeFullName = "http.Header"

const (
	AcceptHeaderName                  = "Accept"
	AcceptCharsetHeaderName           = "Accept-Charset"
	AcceptEncodingHeaderName          = "Accept-Encoding"
	AcceptLanguageHeaderName          = "Accept-Language"
	AcceptRangesHeaderName            = "Accept-Ranges"
	CacheControlHeaderName            = "Cache-Control"
	CcHeaderName                      = "Cc"
	ConnectionHeaderName              = "Connection"
	ContentIdHeaderName               = "Content-Id"
	ContentLanguageHeaderName         = "Content-Language"
	ContentLengthHeaderName           = "Content-Length"
	ContentTransferEncodingHeaderName = "Content-Transfer-Encoding"
	ContentTypeHeaderName             = "Content-Type"
	CookieHeaderName                  = "Cookie"
	DateHeaderName                    = "Date"
	EtagHeaderName                    = "Etag"
	ExpiresHeaderName                 = "Expires"
	FromHeaderName                    = "From"
	HostHeaderName                    = "Host"
	IfModifiedSinceHeaderName         = "If-Modified-Since"
	IfNoneMatchHeaderName             = "If-None-Match"
	InReplyToHeaderName               = "In-Reply-To"
	LastModifiedHeaderName            = "Last-Modified"
	LocationHeaderName                = "Location"
	MessageIdHeaderName               = "Message-Id"
	MimeVersionHeaderName             = "Mime-Version"
	PragmaHeaderName                  = "Pragma"
	ReceivedHeaderName                = "Received"
	ReturnPathHeaderName              = "Return-Path"
	ServerHeaderName                  = "Server"
	SetCookieHeaderName               = "Set-Cookie"
	ToHeaderName                      = "To"
	ViaHeaderName                     = "Via"
	XForwardedForHeaderName           = "X-Forwarded-For"
	XPoweredByHeaderName              = "X-Powered-By"
)

func NewHeader() *Header {
	return &Header{Vals: make(map[string]*core.StringValues)}
}

func NewHeaderFrom(header http.Header) *Header {
	return NewHeader().SyncFrom(header)
}

func (x *Header) Add(key, value string) *Header {
	if x != nil {
		key = textproto.CanonicalMIMEHeaderKey(key)
		if v, ok := x.Vals[key]; ok {
			v.Append(value)
		} else {
			x.Vals[key] = &core.StringValues{}
		}
	}
	return x
}

func (x *Header) Set(key, value string) *Header {
	if x != nil {
		key = textproto.CanonicalMIMEHeaderKey(key)
		x.Vals[key] = core.NewStringValues(value)
	}
	return x
}

func (x *Header) Get(key string) string {
	if x != nil {
		key = textproto.CanonicalMIMEHeaderKey(key)
		if v, ok := x.Vals[key]; ok {
			vals := v.GetVals()
			if len(vals) > 0 {
				return vals[0]
			}
		}
	}
	return ""
}

func (x *Header) Values(key string) *core.StringValues {
	if x != nil {
		key = textproto.CanonicalMIMEHeaderKey(key)
		if v, ok := x.Vals[key]; ok {
			return v
		}
	}
	return nil
}

func (x *Header) SetValues(key string, values ...string) *Header {
	if x != nil {
		key = textproto.CanonicalMIMEHeaderKey(key)
		x.Vals[key] = core.NewStringValues(values...)
	}
	return x
}

func (x *Header) Del(key string) *Header {
	if x != nil {
		key = textproto.CanonicalMIMEHeaderKey(key)
		delete(x.Vals, key)
	}
	return x
}

func (x *Header) Clone() *Header {
	if x == nil {
		return nil
	}

	newHeader := NewHeader()
	for k, v := range x.Vals {
		if v == nil {
			continue
		}
		newHeader.Vals[k] = core.NewStringValues(v.Vals...)
	}

	return newHeader
}

func (x *Header) SyncFrom(header http.Header) *Header {
	if x != nil {
		for k, vals := range header {
			if vals != nil {
				for _, v := range vals {
					x.Add(k, v)
				}
			}
		}
	}
	return x
}

func (x *Header) SyncTo(header http.Header) http.Header {
	if x != nil {
		for k, vals := range x.Vals {
			if vals != nil {
				for _, v := range vals.Vals {
					header.Set(k, v)
				}
			}
		}
	}
	return header
}
