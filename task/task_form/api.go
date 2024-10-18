package task_form

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const baseUrl = "/runtime/tasks/%s/form"

var ListApi = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
