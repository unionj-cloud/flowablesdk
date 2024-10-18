package task

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/runtime/tasks"
	queryUrl  = "/query/tasks"
	detailUrl = baseUrl + "/%s"
)

var (
	ListApi   = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	DetailApi = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	UpdateApi = flowablesdk.NewApi(httpclient.PUT, detailUrl, flowablesdk.ProcessPrefix)
	ActionApi = flowablesdk.NewApi(httpclient.POST, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi = flowablesdk.NewApi(httpclient.DELETE, detailUrl, flowablesdk.ProcessPrefix)
)
