package request

type Request interface {
	Url() string
	Method() string
	Body() string
	DeserialisationTarget() interface{}
	SetParams(params map[string]string)
	SetDeserialisationTarget(target interface{})
}
