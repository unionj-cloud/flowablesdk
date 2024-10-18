package history_detail

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl       = "/history/historic-detail"
	queryUrl      = "/query/historic-detail"
	binaryDataUrl = baseUrl + "/%s/data"
)

var (
	ListApi       = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
	BinaryDataApi = flowablesdk.NewApi(httpclient.GET, binaryDataUrl, flowablesdk.ProcessPrefix)
)
