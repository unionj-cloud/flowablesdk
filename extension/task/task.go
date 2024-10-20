package task_extension

import (
	"encoding/json"
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/task"
)

func Complete(taskId string, req CompleteRequest) (resp task.Task, err error) {
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
