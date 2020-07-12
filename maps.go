package main

type StatusCodeResult struct {
	Description    string
	IncludeHeaders map[string]string
	ExcludeBody    bool
	Link           string
}

var StatusCodes = map[string]StatusCodeResult{
	"100": StatusCodeResult{
		Description: "Continue",
		ExcludeBody: true,
	},
	"101": StatusCodeResult{
		Description: "Switching Protocols",
		ExcludeBody: true,
	},
	"102": StatusCodeResult{
		Description: "Processing",
		ExcludeBody: true,
	},
	"103": StatusCodeResult{
		Description: "Early Hints",
		ExcludeBody: true,
		IncludeHeaders: map[string]string{
			"Link": "</Content/main.css>; rel=preload",
		},
	},
	"200": StatusCodeResult{
		Description: "OK",
	},
	"201": StatusCodeResult{
		Description: "Created",
	},
	"202": StatusCodeResult{
		Description: "Accepted",
	},
	"203": StatusCodeResult{
		Description: "Non-Authoritative Information",
	},
	"204": StatusCodeResult{
		Description: "No Content",
		ExcludeBody: true,
	},
	"205": StatusCodeResult{
		Description: "Reset Content",
		ExcludeBody: true,
	},
	"206": StatusCodeResult{
		Description: "Partial Content",
		IncludeHeaders: map[string]string{
			"Content-Range": "0-30",
		},
	},
	"300": StatusCodeResult{
		Description: "Multiple Choices",
	},
	"301": StatusCodeResult{
		Description: "Moved Permanently",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"302": StatusCodeResult{
		Description: "Found",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"303": StatusCodeResult{
		Description: "See Other",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"304": StatusCodeResult{
		Description: "Not Modified",
		ExcludeBody: true,
	},
	"305": StatusCodeResult{
		Description: "Use Proxy",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"306": StatusCodeResult{
		Description: "Unused",
	},
	"307": StatusCodeResult{
		Description: "Temporary Redirect",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"308": StatusCodeResult{
		Description: "Permanent Redirect",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	},
	"400": StatusCodeResult{
		Description: "Bad Request",
	},
	"401": StatusCodeResult{
		Description: "Unauthorized",
		IncludeHeaders: map[string]string{
			"WWW-Authenticate": "Basic realm=\"Fake Realm\"",
		},
	},
	"402": StatusCodeResult{
		Description: "Payment Required",
	},
	"403": StatusCodeResult{
		Description: "Forbidden",
	},
	"404": StatusCodeResult{
		Description: "Not Found",
	},
	"405": StatusCodeResult{
		Description: "Method Not Allowed",
	},
	"406": StatusCodeResult{
		Description: "Not Acceptable",
	},
	"407": StatusCodeResult{
		Description: "Proxy Authentication Required",
		IncludeHeaders: map[string]string{
			"Proxy-Authenticate": "Basic realm=\"Fake Realm\"",
		},
	},
	"408": StatusCodeResult{
		Description: "Request Timeout",
	},
	"409": StatusCodeResult{
		Description: "Conflict",
	},
	"410": StatusCodeResult{
		Description: "Gone",
	},
	"411": StatusCodeResult{
		Description: "Length Required",
	},
	"412": StatusCodeResult{
		Description: "Precondition Failed",
	},
	"413": StatusCodeResult{
		Description: "Request Entity Too Large",
	},
	"414": StatusCodeResult{
		Description: "Request-URI Too Long",
	},
	"415": StatusCodeResult{
		Description: "Unsupported Media Type",
	},
	"416": StatusCodeResult{
		Description: "Requested Range Not Satisfiable",
	},
	"417": StatusCodeResult{
		Description: "Expectation Failed",
	},
	"418": StatusCodeResult{
		Description: "I'm a teapot",
		Link:        "https://www.ietf.org/rfc/rfc2324.txt",
	},
	"421": StatusCodeResult{
		Description: "Misdirected Request",
	},
	"422": StatusCodeResult{
		Description: "Unprocessable Entity",
	},
	"423": StatusCodeResult{
		Description: "Locked",
	},
	"425": StatusCodeResult{
		Description: "Too Early",
	},
	"426": StatusCodeResult{
		Description: "Upgrade Required",
	},
	"428": StatusCodeResult{
		Description: "Precondition Required",
	},
	"429": StatusCodeResult{
		Description: "Too Many Requests",
		IncludeHeaders: map[string]string{
			"Retry-After": "5",
		},
	},
	"431": StatusCodeResult{
		Description: "Request Header Fields Too Large",
	},
	"451": StatusCodeResult{
		Description: "Unavailable For Legal Reasons",
	},
	"500": StatusCodeResult{
		Description: "Internal Server Error",
	},
	"501": StatusCodeResult{
		Description: "Not Implemented",
	},
	"502": StatusCodeResult{
		Description: "Bad Gateway",
	},
	"503": StatusCodeResult{
		Description: "Service Unavailable",
	},
	"504": StatusCodeResult{
		Description: "Gateway Timeout",
	},
	"505": StatusCodeResult{
		Description: "HTTP Version Not Supported",
	},
	"506": StatusCodeResult{
		Description: "Variant Also Negotiates",
	},
	"507": StatusCodeResult{
		Description: "Insufficient Storage",
	},
	"511": StatusCodeResult{
		Description: "Network Authentication Required",
	},
	"520": StatusCodeResult{
		Description: "Web server is returning an unknown error",
	},
	"522": StatusCodeResult{
		Description: "Connection timed out",
	},
	"524": StatusCodeResult{
		Description: "A timeout occurred",
	},
}
