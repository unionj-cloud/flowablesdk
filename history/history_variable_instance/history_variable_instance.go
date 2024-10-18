package history_variable_instance

import (
	"encoding/json"

	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/common"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/variable"
)

type HistoryVariableInstance struct {
	Id                 string            `json:"id"`
	ProcessInstanceId  string            `json:"processInstanceId"`
	ProcessInstanceUrl string            `json:"processInstanceUrl"`
	TaskId             string            `json:"taskId"`
	Variable           variable.Variable `json:"variable"`
}

// List 获取流程历史变量
func List(req ListRequest) (list []HistoryVariableInstance, count int, err error) {
	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	var commonData common.ListCommonResponse
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return
	}

	count = commonData.Total
	err = json.Unmarshal(commonData.Data, &list)
	return
}

// BinaryData 获取单个流程历史变量二进制文件
func BinaryData(varInstanceId string) (data []byte, err error) {
	request := flowablesdk.GetRequest(BinaryDataApi, varInstanceId)
	data, err = request.DoHttpRequest()
	return
}
