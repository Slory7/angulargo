package httpclient

type IHttpClient interface {
	HttpGetShort(baseURL string, relativeURL string, timeout int32) (result HttpResult, err error)
	HttpGet(baseURL string, relativeURL string, token TokenInfo, xRequestID string, isSecure bool, timeout int32) (result HttpResult, err error)
	HttpPostShort(baseURL string, relativeURL string, postJSONData string, timeout int32) (result HttpResult, err error)
	HttpPost(baseURL string, relativeURL string, postJSONData string, token TokenInfo, xRequestID string, isSecure bool, timeout int32) (result HttpResult, err error)
	HttpPostForm(baseURL string, relativeURL string, postData map[string]string, token TokenInfo, xRequestID string, isSecure bool, timeout int32) (result HttpResult, err error)
	HttpDelete(baseURL string, relativeURL string, token TokenInfo, xRequestID string, isSecure bool, timeout int32) (result HttpResult, err error)
	HttpPut(baseURL string, relativeURL string, postJSONData string, token TokenInfo, xRequestID string, isSecure bool, timeout int32) (result HttpResult, err error)
	HttpSend(baseURL string, relativeURL string, URLParams map[string]string, headers map[string]string, contentType string, method string, postData string, token TokenInfo, xRequestID string, isSecure bool, timeout int32) (result HttpResult, err error)
}
