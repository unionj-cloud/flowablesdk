package history_task_instance

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl       = "/history/historic-task-instances"
	queryUrl      = "/query/historic-task-instances"
	detailUrl     = baseUrl + "/%s"
	identityUrl   = detailUrl + "/identitylinks"
	binaryDataUrl = detailUrl + "/variable/%s/data"
)

var (
	ListApi         = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	DetailApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	DeleteApi       = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	ListIdentityApi = flowablesdk.NewApi(httpclient.GET, identityUrl, flowablesdk.ProcessPrefix)
	BinaryDataApi   = flowablesdk.NewApi(httpclient.GET, binaryDataUrl, flowablesdk.ProcessPrefix)
)
