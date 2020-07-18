//go:generate goversioninfo
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const SLEEP_MIN = 0
const SLEEP_MAX = 300000 // 5 mins in milliseconds
var ROOT_URL string
var Index_Content string

var port string

func init() {
	flag.StringVar(&port, "port", "55555", "http port")
	flag.Parse()

	ROOT_URL = "http://127.0.0.1:" + port + "/"
	Index_Content = GetIndexContent(port)
}

func main() {
	http.HandleFunc("/", doIt)
	println(ROOT_URL + " is ready. --by https://github.com/yongfa365/HttpStatusServer")
	go OpenBrowser(ROOT_URL)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func doIt(rs http.ResponseWriter, rq *http.Request) {
	var message = rq.URL.Query().Get("message")
	var sleep, _ = strconv.Atoi(rq.URL.Query().Get("sleep"))

	var statusCode = ""
	if len(rq.URL.Path) > 3 {
		statusCode = rq.URL.Path[1:4]
	}

	if statusCode == "" {
		_, _ = fmt.Fprintf(rs, Index_Content)
		return
	}

	DoSleep(sleep)

	var statusData StatusCodeResult
	if rs2, ok := StatusCodes[statusCode]; ok {
		statusData = rs2
	} else {
		if message == "" {
			message = "Unknown Code"
		}
		statusData = StatusCodeResult{Description: message}
	}

	ExecuteResult(statusCode, statusData, rq, rs)
}

func ExecuteResult(statusCode string, statusData StatusCodeResult, rq *http.Request, rs http.ResponseWriter) {
	var body = statusCode + " " + statusData.Description
	if statusData.ExcludeBody {
		//remove Content-Length and Content-Type when there isn't any body
		if rs.Header().Get("Content-Length") != "" {
			rs.Header().Del("Content-Length")
		}

		if rs.Header().Get("Content-Type") != "" {
			rs.Header().Del("Content-Type")
		}
		body = ""
	} else {
		var acceptTypes = rq.Header.Get("Accept") //TODO：这些header可能有大小写问题
		if acceptTypes != "" {
			if strings.Contains(acceptTypes, "application/json") {
				//Set the body to be the status code and description with a JSON object type response
				body = "{\"code\": " + statusCode + ", \"description\": \"" + statusData.Description + "\"}"
				rs.Header().Set("Content-Type", "application/json")
				rs.Header().Set("Content-Type", "application/json")
			} else {
				//Set the body to be the status code and description with a plain content type response
				body = statusCode + " " + statusData.Description
				rs.Header().Set("Content-Type", "text/plain")

			}
		}
	}

	if len(statusData.IncludeHeaders) > 0 {
		for k, v := range statusData.IncludeHeaders {
			rs.Header().Set(k, v)
		}
	}
	if code, ok := strconv.Atoi(statusCode); ok == nil {
		rs.WriteHeader(code)
	}

	_, _ = fmt.Fprintf(rs, body)
}

func DoSleep(sleep int) {
	var sleepData = SanitizeSleepParameter(sleep, SLEEP_MIN, SLEEP_MAX)

	if sleepData > 0 {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}

func SanitizeSleepParameter(sleep, min, max int) int {
	var sleepData = sleep

	// range check - minimum should be 0
	if sleepData < min {
		sleepData = min
	}

	// range check- maximum should be 300000 (5 mins)
	if sleepData > max {
		sleepData = max
	}

	return sleepData
}

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
