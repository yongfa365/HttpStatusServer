package main

func GetIndexContent(port string) string {
	return `

<!DOCTYPE html>
<html>
<head>
<title>HttpStatusServer</title>
<style type="text/css">
/* ------------------------------ *
 * Reset Reloaded (by Eric Meyer) *
 * ------------------------------ */
html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, font, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td {
margin: 0;
padding: 0;
border: 0;
outline: 0;
font-weight: inherit;
font-style: inherit;
font-size: 100%;
font-family: inherit;
vertical-align: baseline;
}

body {
line-height: 1;
color: black;
background: white;
}

ol, ul {
list-style: none;
}

table {
border-collapse: separate;
border-spacing: 0;
}

caption, th, td {
text-align: left;
font-weight: normal;
}

blockquote:before, blockquote:after, q:before, q:after {
content: "";
}

blockquote, q {
quotes: """";
}



html {
margin: 20px;
}

body {
width: 778px;
margin: 0 auto;
padding: 15px 40px;
font-family: Helvetica, Arial, sans-serif;
font-size: 16px;
line-height: 1.75;
color: #111;
}

section, header, article, footer {
display: block;
}

h1, h2, h3 {
font-family: Georgia, Times, serif;
}

strong {
font-weight: bold;
}

#header {
margin-bottom: 30px;
}

#title {
font-size: 32px;
color: #3A0;
}

#view p {
margin: 15px 0;
}

#view a:link,
#view a:active,
#view a:visited {
color: #3a0;
}

#view dl {
width: 100%;
overflow: hidden;
margin-left: 30px;
}

#view dl dt {
float: left;
clear: left;
}

#view dl dd {
float: left;
margin-left: 10px;
}

#view dl dd.note {
clear: left;
margin-left: 0;
color: #999;
}

code {
display: block;
white-space: pre;
font-family: Consolas, monospace;
margin-left: 30px;
}


footer {
margin: 50px 0 30px;
padding-top: 50px;
border-top: 1px solid #CCC;
font-size: 11px;
color: #777;
text-transform: uppercase;
}

footer a:link,
footer a:active,
footer a:visited {
color: #777;
}

p.img {
text-align: center;
}

em {
font-style: italic;
font-size: 12px;
}
</style>


</head>
<body>
<div id='content'>
<header>
<h1 id='title'>127.0.0.1:` + port + `</h1>
</header>
<div id='view'>
<p>This is a super simple service for generating different HTTP codes.</p>
<p>It's useful for testing how your own scripts deal with varying responses.</p>
<p>
Just add the status code you want to the URL, like this:
<a href="/200">http://127.0.0.1:` + port + `/200</a>
</p>
<p>We'll return a response like this:</p>
<code>HTTP/1.1 {status code} {status description}
Content-Type: text/plain or application/json
Content-Length: {something}
{any custom response headers}

{status code} {status description}
{list of any custom response headers we added}</code>

<p>
To get a JSON response back, you need to ensure that the Accept header contains 'application/json'. Then we'll JSON encode the response and send the Content-Type header accordingly.
</p>

<p>
If you want a delay on the response add a query string of sleep (the time in ms, max 5 minutes*), like this:
<a href="/200?sleep=5000">127.0.0.1:` + port + `/200?sleep=5000</a>
<br />
<em>*When using the hosted instance the timeout is actually 230 seconds, which is the max timeout allowed by an Azure WebApp (see <a href="https://social.msdn.microsoft.com/Forums/en-US/05f254a6-9b34-4eb2-a5f7-2a82fb40135f/time-out-after-230-seconds?forum=windowsazurewebsitespreview" target="_blank">this thread post</a>). If you host it yourself in IIS/IIS Express you won't have that limit.</em>
</p>

<p>Here are all the codes we support (and any special notes):</p>
<dl>
<dt><a href="/100">100</a></dt>
<dd>
Continue        </dd>
<dt><a href="/101">101</a></dt>
<dd>
Switching Protocols        </dd>
<dt><a href="/102">102</a></dt>
<dd>
Processing        </dd>
<dt><a href="/103">103</a></dt>
<dd>
Early Hints        </dd>
<dt><a href="/200">200</a></dt>
<dd>
OK        </dd>
<dt><a href="/201">201</a></dt>
<dd>
Created        </dd>
<dt><a href="/202">202</a></dt>
<dd>
Accepted        </dd>
<dt><a href="/203">203</a></dt>
<dd>
Non-Authoritative Information        </dd>
<dt><a href="/204">204</a></dt>
<dd>
No Content        </dd>
<dt><a href="/205">205</a></dt>
<dd>
Reset Content        </dd>
<dt><a href="/206">206</a></dt>
<dd>
Partial Content        </dd>
<dt><a href="/300">300</a></dt>
<dd>
Multiple Choices        </dd>
<dt><a href="/301">301</a></dt>
<dd>
Moved Permanently        </dd>
<dt><a href="/302">302</a></dt>
<dd>
Found        </dd>
<dt><a href="/303">303</a></dt>
<dd>
See Other        </dd>
<dt><a href="/304">304</a></dt>
<dd>
Not Modified        </dd>
<dt><a href="/305">305</a></dt>
<dd>
Use Proxy        </dd>
<dt><a href="/306">306</a></dt>
<dd>
Unused        </dd>
<dt><a href="/307">307</a></dt>
<dd>
Temporary Redirect        </dd>
<dt><a href="/308">308</a></dt>
<dd>
Permanent Redirect        </dd>
<dt><a href="/400">400</a></dt>
<dd>
Bad Request        </dd>
<dt><a href="/401">401</a></dt>
<dd>
Unauthorized        </dd>
<dt><a href="/402">402</a></dt>
<dd>
Payment Required        </dd>
<dt><a href="/403">403</a></dt>
<dd>
Forbidden        </dd>
<dt><a href="/404">404</a></dt>
<dd>
Not Found        </dd>
<dt><a href="/405">405</a></dt>
<dd>
Method Not Allowed        </dd>
<dt><a href="/406">406</a></dt>
<dd>
Not Acceptable        </dd>
<dt><a href="/407">407</a></dt>
<dd>
Proxy Authentication Required        </dd>
<dt><a href="/408">408</a></dt>
<dd>
Request Timeout        </dd>
<dt><a href="/409">409</a></dt>
<dd>
Conflict        </dd>
<dt><a href="/410">410</a></dt>
<dd>
Gone        </dd>
<dt><a href="/411">411</a></dt>
<dd>
Length Required        </dd>
<dt><a href="/412">412</a></dt>
<dd>
Precondition Failed        </dd>
<dt><a href="/413">413</a></dt>
<dd>
Request Entity Too Large        </dd>
<dt><a href="/414">414</a></dt>
<dd>
Request-URI Too Long        </dd>
<dt><a href="/415">415</a></dt>
<dd>
Unsupported Media Type        </dd>
<dt><a href="/416">416</a></dt>
<dd>
Requested Range Not Satisfiable        </dd>
<dt><a href="/417">417</a></dt>
<dd>
Expectation Failed        </dd>
<dt><a href="/418">418</a></dt>
<dd>
<a href="https://www.ietf.org/rfc/rfc2324.txt" target="_blank" title="I'm a teapot">I'm a teapot</a>
</dd>
<dt><a href="/421">421</a></dt>
<dd>
Misdirected Request        </dd>
<dt><a href="/422">422</a></dt>
<dd>
Unprocessable Entity        </dd>
<dt><a href="/423">423</a></dt>
<dd>
Locked        </dd>
<dt><a href="/425">425</a></dt>
<dd>
Too Early        </dd>
<dt><a href="/426">426</a></dt>
<dd>
Upgrade Required        </dd>
<dt><a href="/428">428</a></dt>
<dd>
Precondition Required        </dd>
<dt><a href="/429">429</a></dt>
<dd>
Too Many Requests        </dd>
<dt><a href="/431">431</a></dt>
<dd>
Request Header Fields Too Large        </dd>
<dt><a href="/451">451</a></dt>
<dd>
Unavailable For Legal Reasons        </dd>
<dt><a href="/500">500</a></dt>
<dd>
Internal Server Error        </dd>
<dt><a href="/501">501</a></dt>
<dd>
Not Implemented        </dd>
<dt><a href="/502">502</a></dt>
<dd>
Bad Gateway        </dd>
<dt><a href="/503">503</a></dt>
<dd>
Service Unavailable        </dd>
<dt><a href="/504">504</a></dt>
<dd>
Gateway Timeout        </dd>
<dt><a href="/505">505</a></dt>
<dd>
HTTP Version Not Supported        </dd>
<dt><a href="/506">506</a></dt>
<dd>
Variant Also Negotiates        </dd>
<dt><a href="/507">507</a></dt>
<dd>
Insufficient Storage        </dd>
<dt><a href="/511">511</a></dt>
<dd>
Network Authentication Required        </dd>
<dt><a href="/520">520</a></dt>
<dd>
Web server is returning an unknown error        </dd>
<dt><a href="/522">522</a></dt>
<dd>
Connection timed out        </dd>
<dt><a href="/524">524</a></dt>
<dd>
A timeout occurred        </dd>
</dl>

</div>

<div style="clear:both"></div>
<p>If you send any other three digit number that's not in that list, we'll return it too. Or, <a href="https://github.com/yongfa365/HttpStatusServer">send us a pull request</a> to add full support for a new code.</p>
<p>Enjoy!</p>

Created by
https://github.com/yongfa365/HttpStatusServer
<strong>We don't capture or store any data about the requests you make.</strong>
</body>
</html>

`
}
