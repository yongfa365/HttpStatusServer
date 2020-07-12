package main

type StatusCodeResult struct {
	Description    string
	IncludeHeaders map[string]string
	ExcludeBody    bool
	Link           string
}

var StatusCodes = map[string]StatusCodeResult{
	"100": {
		Description: "Continue",
		ExcludeBody: true,
	},
	"101": {
		Description: "Switching Protocols",
		ExcludeBody: true,
	},
	"102": {
		Description: "Processing",
		ExcludeBody: true,
	},
	"103": {
		Description: "Early Hints",
		ExcludeBody: true,
		IncludeHeaders: map[string]string{
			"Link": "</Content/main.css>; rel=preload",
		},
	},
	"200": {
		Description: "OK",
	},
	"201": {
		Description: "Created",
	},
	"202": {
		Description: "Accepted",
	},
	"203": {
		Description: "Non-Authoritative Information",
	},
	"204": {
		Description: "No Content",
		ExcludeBody: true,
	},
	"205": {
		Description: "Reset Content",
		ExcludeBody: true,
	},
	"206": {
		Description: "Partial Content",
		IncludeHeaders: map[string]string{
			"Content-Range": "0-30",
		},
	},
	"300": {
		Description: "Multiple Choices",
	},
	"301": {
		Description: "Moved Permanently",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"302": {
		Description: "Found",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"303": {
		Description: "See Other",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"304": {
		Description: "Not Modified",
		ExcludeBody: true,
	},
	"305": {
		Description: "Use Proxy",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"306": {
		Description: "Unused",
	},
	"307": {
		Description: "Temporary Redirect",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"308": {
		Description: "Permanent Redirect",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"400": {
		Description: "Bad Request",
	},
	"401": {
		Description: "Unauthorized",
		IncludeHeaders: map[string]string{
			"WWW-Authenticate": "Basic realm=\"Fake Realm\"",
		},
	},
	"402": {
		Description: "Payment Required",
	},
	"403": {
		Description: "Forbidden",
	},
	"404": {
		Description: "Not Found",
	},
	"405": {
		Description: "Method Not Allowed",
	},
	"406": {
		Description: "Not Acceptable",
	},
	"407": {
		Description: "Proxy Authentication Required",
		IncludeHeaders: map[string]string{
			"Proxy-Authenticate": "Basic realm=\"Fake Realm\"",
		},
	},
	"408": {
		Description: "Request Timeout",
	},
	"409": {
		Description: "Conflict",
	},
	"410": {
		Description: "Gone",
	},
	"411": {
		Description: "Length Required",
	},
	"412": {
		Description: "Precondition Failed",
	},
	"413": {
		Description: "Request Entity Too Large",
	},
	"414": {
		Description: "Request-URI Too Long",
	},
	"415": {
		Description: "Unsupported Media Type",
	},
	"416": {
		Description: "Requested Range Not Satisfiable",
	},
	"417": {
		Description: "Expectation Failed",
	},
	"418": {
		Description: "I'm a teapot",
		Link:        "https://www.ietf.org/rfc/rfc2324.txt",
	},
	"421": {
		Description: "Misdirected Request",
	},
	"422": {
		Description: "Unprocessable Entity",
	},
	"423": {
		Description: "Locked",
	},
	"425": {
		Description: "Too Early",
	},
	"426": {
		Description: "Upgrade Required",
	},
	"428": {
		Description: "Precondition Required",
	},
	"429": {
		Description: "Too Many Requests",
		IncludeHeaders: map[string]string{
			"Retry-After": "5",
		},
	},
	"431": {
		Description: "Request Header Fields Too Large",
	},
	"451": {
		Description: "Unavailable For Legal Reasons",
	},
	"500": {
		Description: "Internal Server Error",
	},
	"501": {
		Description: "Not Implemented",
	},
	"502": {
		Description: "Bad Gateway",
	},
	"503": {
		Description: "Service Unavailable",
	},
	"504": {
		Description: "Gateway Timeout",
	},
	"505": {
		Description: "HTTP Version Not Supported",
	},
	"506": {
		Description: "Variant Also Negotiates",
	},
	"507": {
		Description: "Insufficient Storage",
	},
	"511": {
		Description: "Network Authentication Required",
	},
	"520": {
		Description: "Web server is returning an unknown error",
	},
	"522": {
		Description: "Connection timed out",
	},
	"524": {
		Description: "A timeout occurred",
	},
}
