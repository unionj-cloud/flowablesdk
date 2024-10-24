package task_extension

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl     = "/runtime/tasks"
	detailUrl   = baseUrl + "/%s"
	completeUrl = detailUrl + "/complete"
	assignUrl   = detailUrl + "/assign"
)

var (
	CompleteApi = flowablesdk.NewApi(httpclient.POST, completeUrl, flowablesdk.ExtensionPrefix)
	AssignApi   = flowablesdk.NewApi(httpclient.POST, assignUrl, flowablesdk.ExtensionPrefix)
)
