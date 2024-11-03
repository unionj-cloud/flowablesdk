package task_extension

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/task"
)

func Complete(taskId string, req task.ActionRequest) (err error) {
	request := flowablesdk.GetRequest(CompleteApi, taskId)
	request.With(httpclient.WithJson(req))
	_, err = request.DoHttpRequest()
	return
}

func Assign(taskId string, req task.ActionRequest) (err error) {
	request := flowablesdk.GetRequest(AssignApi, taskId)
	request.With(httpclient.WithJson(req))
	_, err = request.DoHttpRequest()
	return
}

func Recall(taskId string) (err error) {
	request := flowablesdk.GetRequest(RecallApi, taskId)
	_, err = request.DoHttpRequest()
	return
}

func CompleteByUser(taskId string, req task.ActionRequest, userId string) (err error) {
	request := flowablesdk.GetRequest(CompleteApi, taskId)
	request.With(
		httpclient.WithJson(req),
		httpclient.WithHeader(httpclient.FlowableUserId, userId),
	)
	_, err = request.DoHttpRequest()
	return
}

func AssignByUser(taskId string, req task.ActionRequest, userId string) (err error) {
	request := flowablesdk.GetRequest(AssignApi, taskId)
	request.With(
		httpclient.WithJson(req),
		httpclient.WithHeader(httpclient.FlowableUserId, userId),
	)
	_, err = request.DoHttpRequest()
	return
}

func RecallByUser(taskId string, userId string) (err error) {
	request := flowablesdk.GetRequest(RecallApi, taskId)
	request.With(
		httpclient.WithHeader(httpclient.FlowableUserId, userId),
	)
	_, err = request.DoHttpRequest()
	return
}
