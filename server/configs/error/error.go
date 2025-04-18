package httperror

const (
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	NOT_FOUND             = "NOT_FOUND"
	URL_NOT_FOUND         = "URL_NOT_FOUND"
	MOVED_PERMANENTLY     = "MOVED_PERMANENTLY"
	USE_PROXY             = "USE_PROXY"
)

/* { [string 错误码]: [int http状态码] } */
var ErrorCodeMap = map[string]int{
	INTERNAL_SERVER_ERROR: 500,
	NOT_FOUND:             404,
	URL_NOT_FOUND:         404,
	MOVED_PERMANENTLY:     301,
	USE_PROXY:             305,
}
