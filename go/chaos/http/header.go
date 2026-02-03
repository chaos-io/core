package http

import (
	core "github.com/chaos-io/core/go/chaos/core"
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

// func NewHeaderFrom(header http.Header) *Header {
// 	return NewHeader().SyncFrom(header)
// }
