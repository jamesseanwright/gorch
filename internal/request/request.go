package request

type Request interface {
	GetUrl() string
	GetMethod() string
	SetParams(params map[string]string)
}
