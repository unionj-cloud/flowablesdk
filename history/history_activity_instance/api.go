package history_activity_instance

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	queryUrl = "/query/historic-activity-instances"
)

var (
	ListApi = flowablesdk.NewApi(httpclient.POST, queryUrl, flowablesdk.ProcessPrefix)
)
