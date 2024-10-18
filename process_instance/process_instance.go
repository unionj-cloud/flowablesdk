package process_instance

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/candidate"
	"github.com/unionj-cloud/flowablesdk/common"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
	"github.com/unionj-cloud/flowablesdk/variable"
)

type ProcessInstance struct {
	Id                           string              `json:"id"`
	Url                          string              `json:"url"`
	BusinessKey                  string              `json:"businessKey"`
	Suspended                    bool                `json:"suspended"`
	Ended                        bool                `json:"ended"`
	ProcessDefinitionId          string              `json:"processDefinitionId"`
	ProcessDefinitionUrl         string              `json:"processDefinitionUrl"`
	ProcessDefinitionName        string              `json:"processDefinitionName"`
	ProcessDefinitionDescription string              `json:"processDefinitionDescription"`
	ActivityId                   string              `json:"activityId"`
	StartUserId                  string              `json:"startUserId"`
	StartTime                    time.Time           `json:"startTime"`
	Variables                    []variable.Variable `json:"variables"`
	CallbackId                   string              `json:"callbackId"`
	CallbackType                 string              `json:"callbackType"`
	ReferenceId                  string              `json:"referenceId"`
	ReferenceType                string              `json:"referenceType"`
	TenantId                     string              `json:"tenantId"`
}

// List 返回所有流程实例
func List(req ListRequest) (list []ProcessInstance, count int, err error) {
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

// Detail 返回单个流程实例
func Detail(instanceId string) (resp ProcessInstance, err error) {
	request := flowablesdk.GetRequest(DetailApi, instanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Update 更新一个流程实例
func Update(instanceId string, req UpdateRequest) (resp ProcessInstance, err error) {
	request := flowablesdk.GetRequest(UpdateApi, instanceId)

	if len(req.Action) > 0 {
		if !(req.Action == "activate" || req.Action == "suspend") {
			err = errors.New("action is not allow")
			return
		}
	}

	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// Delete 删除单个流程实例
func Delete(instanceId string) error {
	request := flowablesdk.GetRequest(DeleteApi, instanceId)
	_, err := request.DoHttpRequest()
	return err
}

// Start 启动一个流程实例
func Start(req StartProcessRequest) (resp ProcessInstance, err error) {
	request := flowablesdk.GetRequest(StartApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// ListCandidate 获取单个流程实例所有相关人员
func ListCandidate(instanceId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListCandidateApi, instanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddCandidate 单个流程实例添加参与用户
func AddCandidate(instanceId string, req candidate.Candidate) (resp candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(AddCandidateApi, instanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// DeleteCandidate 单个流程实例删除参与用户
func DeleteCandidate(instanceId string, req candidate.Candidate) error {
	request := flowablesdk.GetRequest(DeleteCandidateApi, instanceId, req.User, req.Type)
	_, err := request.DoHttpRequest()
	return err
}

// ListVariables 获取单个流程实例所有变量
func ListVariables(instanceId string) (resp []variable.Variable, err error) {
	request := flowablesdk.GetRequest(ListVariablesApi, instanceId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddVariables 单个流程实例添加变量
func AddVariables(instanceId string, req []variable.VariableRequest) (resp []variable.Variable, err error) {
	request := flowablesdk.GetRequest(AddVariablesApi, instanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// UpdateVariables 单个流程实例编辑变量(多个)
func UpdateVariables(instanceId string, req []variable.VariableRequest) (resp []variable.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariablesApi, instanceId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// UpdateVariable 单个流程实例编辑变量(单个)
func UpdateVariable(instanceId, variable string, req variable.VariableRequest) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariableApi, instanceId, variable)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// VariableDetail 获取单个流程实例所有变量
func VariableDetail(instanceId, variable string) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(VariableDetailApi, instanceId, variable)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// AddFileVariable 单个流程实例添加文件变量
func AddFileVariable(instanceId string, req variable.FileVariableRequest) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(AddVariablesApi, instanceId)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, map[string]string{
		"name": req.VariableName,
		"type": "binary",
	}))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}

// UpdateFileVariable 单个流程实例编辑文件变量
func UpdateFileVariable(instanceId string, req variable.FileVariableRequest) (resp variable.Variable, err error) {
	request := flowablesdk.GetRequest(UpdateVariableApi, instanceId, req.VariableName)
	request.With(httpclient.WithMultipartFrom(&httpclient.UploadFile{
		Field:    "file",
		FileName: req.FileName,
		File:     req.Value,
	}, map[string]string{
		"name": req.VariableName,
		"type": "binary",
	}))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &resp)
	return
}
