package process_instance_extension

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl = "/runtime/process-instances"
)

var (
	StartApi = flowablesdk.NewApi(httpclient.POST, baseUrl, flowablesdk.ExtensionPrefix)
)
