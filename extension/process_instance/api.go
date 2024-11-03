package process_instance_extension

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl   = "/runtime/process-instances"
	detailUrl = baseUrl + "/%s"
	moveUrl   = detailUrl + "/move"
)

var (
	MoveApi = flowablesdk.NewApi(httpclient.POST, moveUrl, flowablesdk.ExtensionPrefix)
)
