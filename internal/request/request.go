package request

type Request interface {
	GetUrl() string
	GetMethod() string
	GetBody() string
	SetParams(params map[string]string)
}
