package request

type Request interface {
	Url() string
	Method() string
	Body() string
	SetParams(params map[string]string)
	SetHeaders(params map[string]string)
}
