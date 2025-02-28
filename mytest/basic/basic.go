package basic

// pcurl represents a curl-like HTTP client with various options for making HTTP requests.
// :quickclop
type pcurl struct {
	Method        string   `clop:"-X; --request" usage:"Specify request command to use"`
	Get           bool     `clop:"-G; --get" usage:"Put the post data in the URL and use GET"`
	Header        []string `clop:"-H; --header" usage:"Pass custom header(s) to server"`
	Data          string   `clop:"-d; --data"   usage:"HTTP POST data"`
	DataRaw       string   `clop:"--data-raw" usage:"HTTP POST data, '@' allowed"`
	Form          []string `clop:"-F; --form" usage:"Specify multipart MIME data"`
	URL2          string   `clop:"args=url2" usage:"url2"`
	URL           string   `clop:"--url" usage:"URL to work with"`
	Location      bool     `clop:"-L; --location" usage:"Follow redirects"`
	DataUrlencode []string `clop:"--data-urlencode" usage:"HTTP POST data url encoded"`

	Compressed bool `clop:"--compressed" usage:"Request compressed response"`
	// 在响应包里面打印http header, 仅做字段赋值
	Include  bool `clop:"-i;--include" usage:"Include the HTTP response headers in the output. The HTTP response headers can include things like server name, cookies, date of the document, HTTP version and more."`
	Insecure bool `clop:"-k; --insecure" usage:"Allow insecure server connections when using SSL"`
}
