package task_extension

import (
	"encoding/json"
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/task"
)

func Complete(taskId string, req task.ActionRequest) (resp task.Task, err error) {
	request := flowablesdk.GetRequest(CompleteApi, taskId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	if len(data) == 0 {
		return task.Task{}, nil
	}
	err = json.Unmarshal(data, &resp)
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

func Move(taskId string, req MoveRequest) (resp task.Task, err error) {
	request := flowablesdk.GetRequest(MoveApi, taskId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	if len(data) == 0 {
		return task.Task{}, nil
	}
	err = json.Unmarshal(data, &resp)
	return
}

func CompleteByUser(taskId string, req task.ActionRequest, userId string) (resp task.Task, err error) {
	request := flowablesdk.GetRequest(CompleteApi, taskId)
	request.With(
		httpclient.WithJson(req),
		httpclient.WithHeader(httpclient.FlowableUserId, userId),
	)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	if len(data) == 0 {
		return task.Task{}, nil
	}
	err = json.Unmarshal(data, &resp)
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

func MoveByUser(taskId string, req MoveRequest, userId string) (resp task.Task, err error) {
	request := flowablesdk.GetRequest(MoveApi, taskId)
	request.With(
		httpclient.WithJson(req),
		httpclient.WithHeader(httpclient.FlowableUserId, userId),
	)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	if len(data) == 0 {
		return task.Task{}, nil
	}
	err = json.Unmarshal(data, &resp)
	return
}
