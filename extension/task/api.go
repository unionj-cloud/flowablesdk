package task_extension

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl     = "/runtime/tasks"
	detailUrl   = baseUrl + "/%s"
	completeUrl = detailUrl + "/complete"
)

var (
	CompleteApi = flowablesdk.NewApi(httpclient.POST, completeUrl, flowablesdk.ExtensionPrefix)
)
