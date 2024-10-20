package process_instance_extension

import (
	"encoding/json"
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/task"
)

// Start 启动一个流程实例
func Start(req StartProcessRequest) (resp task.Task, err error) {
	request := flowablesdk.GetRequest(StartApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}
