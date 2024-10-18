package form

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl = "/form/form-data"
)

var (
	DetailFromApi = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	SubmitFromApi = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ProcessPrefix)
)
