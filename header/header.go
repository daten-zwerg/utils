package header

import "fmt"

type HeaderField string

const (
	Accept             HeaderField = "Accept"
	AcceptCharset                  = "Accept-Charset"
	AcceptEncoding                 = "Accept-Encoding"
	AcceptLanguage                 = "Accept-Language"
	Authorization                  = "Authorization"
	CacheControl                   = "Cache-Control"
	Connection                     = "Connection"
	Cookie                         = "Cookie"
	ContentLength                  = "Content-Length"
	ContentMD5                     = "Content-MD5"
	ContentType                    = "Content-Type"
	Date                           = "Date"
	Expect                         = "Expect"
	Forwarded                      = "Forwarded"
	From                           = "From"
	Host                           = "Host"
	IfMatch                        = "If-Match"
	IfModifiedSince                = "If-Modified-Since"
	IfNoneMatch                    = "If-None-Match"
	IfRange                        = "If-Range"
	IfUnmodifiedSince              = "If-Unmodified-Since"
	MaxForwards                    = "Max-Forwards"
	Pragma                         = "Pragma"
	ProxyAuthorization             = "Proxy-Authorization"
	Range                          = "Range"
	Referer                        = "Referer"
	TE                             = "TE"
	TransferEncoding               = "Transfer-Encoding"
	Upgrade                        = "Upgrade"
	UserAgent                      = "User-Agent"
	Via                            = "Via"
	Warning                        = "Warning"
)

type Header struct {
	fields map[HeaderField]string
}

func (h *Header) Add(f HeaderField, v string) {
	h.fields[f] = v
}

func (h *Header) String() string {
	var str string
	for k, v := range h.fields {
		str += fmt.Sprintf("%s: %s\n", k, v)
	}
	return str
}

func NewHeader() *Header {
	h := Header{
		fields: make(map[HeaderField]string),
	}
	return &h
}
