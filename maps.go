package main

type StatusCodeResult struct {
	Description    string
	IncludeHeaders map[string]string
	ExcludeBody    bool
	Link           string
}

var StatusCodes = make(map[string]StatusCodeResult)

func InitStatusCodes() {
	StatusCodes["100"] = StatusCodeResult{
		Description: "Continue",
		ExcludeBody: true,
	}
	StatusCodes["101"] = StatusCodeResult{
		Description: "Switching Protocols",
		ExcludeBody: true,
	}
	StatusCodes["102"] = StatusCodeResult{
		Description: "Processing",
		ExcludeBody: true,
	}
	StatusCodes["103"] = StatusCodeResult{
		Description: "Early Hints",
		ExcludeBody: true,
		IncludeHeaders: map[string]string{
			"Link": "</Content/main.css>; rel=preload",
		},
	}
	StatusCodes["200"] = StatusCodeResult{
		Description: "OK",
	}
	StatusCodes["201"] = StatusCodeResult{
		Description: "Created",
	}
	StatusCodes["202"] = StatusCodeResult{
		Description: "Accepted",
	}
	StatusCodes["203"] = StatusCodeResult{
		Description: "Non-Authoritative Information",
	}
	StatusCodes["204"] = StatusCodeResult{
		Description: "No Content",
		ExcludeBody: true,
	}
	StatusCodes["205"] = StatusCodeResult{
		Description: "Reset Content",
		ExcludeBody: true,
	}
	StatusCodes["206"] = StatusCodeResult{
		Description: "Partial Content",
		IncludeHeaders: map[string]string{
			"Content-Range": "0-30",
		},
	}
	StatusCodes["300"] = StatusCodeResult{
		Description: "Multiple Choices",
	}
	StatusCodes["301"] = StatusCodeResult{
		Description: "Moved Permanently",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	}
	StatusCodes["302"] = StatusCodeResult{
		Description: "Found",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	}
	StatusCodes["303"] = StatusCodeResult{
		Description: "See Other",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	}
	StatusCodes["304"] = StatusCodeResult{
		Description: "Not Modified",
		ExcludeBody: true,
	}
	StatusCodes["305"] = StatusCodeResult{
		Description: "Use Proxy",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	}
	StatusCodes["306"] = StatusCodeResult{
		Description: "Unused",
	}
	StatusCodes["307"] = StatusCodeResult{
		Description: "Temporary Redirect",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	}
	StatusCodes["308"] = StatusCodeResult{
		Description: "Permanent Redirect",
		IncludeHeaders: map[string]string{
			"Location": "http://127.0.0.1:5555",
		},
	}
	StatusCodes["400"] = StatusCodeResult{
		Description: "Bad Request",
	}
	StatusCodes["401"] = StatusCodeResult{
		Description: "Unauthorized",
		IncludeHeaders: map[string]string{
			"WWW-Authenticate": "Basic realm=\"Fake Realm\"",
		},
	}
	StatusCodes["402"] = StatusCodeResult{
		Description: "Payment Required",
	}
	StatusCodes["403"] = StatusCodeResult{
		Description: "Forbidden",
	}
	StatusCodes["404"] = StatusCodeResult{
		Description: "Not Found",
	}
	StatusCodes["405"] = StatusCodeResult{
		Description: "Method Not Allowed",
	}
	StatusCodes["406"] = StatusCodeResult{
		Description: "Not Acceptable",
	}
	StatusCodes["407"] = StatusCodeResult{
		Description: "Proxy Authentication Required",
		IncludeHeaders: map[string]string{
			"Proxy-Authenticate": "Basic realm=\"Fake Realm\"",
		},
	}
	StatusCodes["408"] = StatusCodeResult{
		Description: "Request Timeout",
	}
	StatusCodes["409"] = StatusCodeResult{
		Description: "Conflict",
	}
	StatusCodes["410"] = StatusCodeResult{
		Description: "Gone",
	}
	StatusCodes["411"] = StatusCodeResult{
		Description: "Length Required",
	}
	StatusCodes["412"] = StatusCodeResult{
		Description: "Precondition Failed",
	}
	StatusCodes["413"] = StatusCodeResult{
		Description: "Request Entity Too Large",
	}
	StatusCodes["414"] = StatusCodeResult{
		Description: "Request-URI Too Long",
	}
	StatusCodes["415"] = StatusCodeResult{
		Description: "Unsupported Media Type",
	}
	StatusCodes["416"] = StatusCodeResult{
		Description: "Requested Range Not Satisfiable",
	}
	StatusCodes["417"] = StatusCodeResult{
		Description: "Expectation Failed",
	}
	StatusCodes["418"] = StatusCodeResult{
		Description: "I'm a teapot",
		Link:        "https://www.ietf.org/rfc/rfc2324.txt",
	}
	StatusCodes["421"] = StatusCodeResult{
		Description: "Misdirected Request",
	}
	StatusCodes["422"] = StatusCodeResult{
		Description: "Unprocessable Entity",
	}
	StatusCodes["423"] = StatusCodeResult{
		Description: "Locked",
	}
	StatusCodes["425"] = StatusCodeResult{
		Description: "Too Early",
	}
	StatusCodes["426"] = StatusCodeResult{
		Description: "Upgrade Required",
	}
	StatusCodes["428"] = StatusCodeResult{
		Description: "Precondition Required",
	}
	StatusCodes["429"] = StatusCodeResult{
		Description: "Too Many Requests",
		IncludeHeaders: map[string]string{
			"Retry-After": "5",
		},
	}
	StatusCodes["431"] = StatusCodeResult{
		Description: "Request Header Fields Too Large",
	}
	StatusCodes["451"] = StatusCodeResult{
		Description: "Unavailable For Legal Reasons",
	}
	StatusCodes["500"] = StatusCodeResult{
		Description: "Internal Server Error",
	}
	StatusCodes["501"] = StatusCodeResult{
		Description: "Not Implemented",
	}
	StatusCodes["502"] = StatusCodeResult{
		Description: "Bad Gateway",
	}
	StatusCodes["503"] = StatusCodeResult{
		Description: "Service Unavailable",
	}
	StatusCodes["504"] = StatusCodeResult{
		Description: "Gateway Timeout",
	}
	StatusCodes["505"] = StatusCodeResult{
		Description: "HTTP Version Not Supported",
	}
	StatusCodes["506"] = StatusCodeResult{
		Description: "Variant Also Negotiates",
	}
	StatusCodes["507"] = StatusCodeResult{
		Description: "Insufficient Storage",
	}
	StatusCodes["511"] = StatusCodeResult{
		Description: "Network Authentication Required",
	}
	StatusCodes["520"] = StatusCodeResult{
		Description: "Web server is returning an unknown error",
	}
	StatusCodes["522"] = StatusCodeResult{
		Description: "Connection timed out",
	}
	StatusCodes["524"] = StatusCodeResult{
		Description: "A timeout occurred",
	}
}
