package history_task_instance

import (
	"encoding/json"

	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/common"
	"github.com/unionj-cloud/flowablesdk/identity"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/variable"
)

type HistoryProcessInstanceTask struct {
	Id                   string              `json:"id"`
	ProcessDefinitionId  string              `json:"processDefinitionId"`
	ProcessDefinitionUrl string              `json:"processDefinitionUrl"`
	ProcessInstanceId    string              `json:"processInstanceId"`
	ProcessInstanceUrl   string              `json:"processInstanceUrl"`
	ExecutionId          string              `json:"executionId"`
	Name                 string              `json:"name"`
	Description          string              `json:"description"`
	DeleteReason         string              `json:"deleteReason"`
	Owner                string              `json:"owner"`
	Assignee             string              `json:"assignee"`
	StartTime            string              `json:"startTime"`
	EndTime              string              `json:"endTime"`
	DurationInMillis     int                 `json:"durationInMillis"`
	WorkTimeInMillis     int                 `json:"workTimeInMillis"`
	ClaimTime            string              `json:"claimTime"`
	TaskDefinitionKey    string              `json:"taskDefinitionKey"`
	FormKey              string              `json:"formKey"`
	Priority             int                 `json:"priority"`
	DueDate              string              `json:"dueDate"`
	ParentTaskId         string              `json:"parentTaskId"`
	Url                  string              `json:"url"`
	Variables            []variable.Variable `json:"variables"`
	TenantId             string              `json:"tenantId"`
	Category             string              `json:"category"`
}

// List 获取历史任务
func List(req ListRequest) (list []HistoryProcessInstanceTask, count int, err error) {
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

// Detail 获取单个历史任务
func Detail(taskId string) (resp HistoryProcessInstanceTask, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个历史任务
func Delete(taskId string) error {
	request := flowablesdk.GetRequest(DeleteApi, taskId)
	_, err := request.DoHttpRequest()
	return err
}

// ListIdentity 单个历史任务相关用户
func ListIdentity(taskId string) (resp []identity.Identity, err error) {
	request := flowablesdk.GetRequest(ListIdentityApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// BinaryData 获取单个历史任务二进制文件
func BinaryData(taskId, variableName string) (data []byte, err error) {
	request := flowablesdk.GetRequest(BinaryDataApi, taskId, variableName)
	data, err = request.DoHttpRequest()
	return
}
