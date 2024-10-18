package form_definition

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/form-repository/form-definitions"
	detailUrl = baseUrl + "/%s"
	modelUrl  = detailUrl + "/model"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.FormPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.FormPrefix)
	ModelApi  = flowablesdk.NewApi(httpclient.GET, modelUrl, flowablesdk.FormPrefix)
)
