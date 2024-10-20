package flowablesdk

import "github.com/unionj-cloud/flowablesdk/pkg/httpclient"

type Api struct {
	Method httpclient.Method
	Url    string
	Prefix int
}

const (
	ProcessPrefix = iota
	FormPrefix
	ExtensionPrefix
)

func NewApi(m httpclient.Method, url string, prefix int) *Api {
	return &Api{
		Method: m,
		Url:    url,
		Prefix: prefix,
	}
}
