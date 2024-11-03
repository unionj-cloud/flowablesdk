package process_instance_extension

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

func MoveByUser(processInstanceId string, req MoveRequest, userId string) (err error) {
	request := flowablesdk.GetRequest(MoveApi, processInstanceId)
	request.With(
		httpclient.WithJson(req),
		httpclient.WithHeader(httpclient.FlowableUserId, userId),
	)
	_, err = request.DoHttpRequest()
	return
}
