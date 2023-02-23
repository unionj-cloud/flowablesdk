package history_variable_instance

import (
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest
	common.WithTenant
	ActivityId          string                           `json:"activityId,omitempty"`
	ActivityInstanceId  string                           `json:"activityInstanceId,omitempty"`
	ActivityName        string                           `json:"activityName,omitempty"`
	ActivityType        string                           `json:"activityType,omitempty"`
	ExecutionId         string                           `json:"executionId,omitempty"`
	Finished            *bool                            `json:"finished,omitempty"`
	TaskAssignee        string                           `json:"taskAssignee,omitempty"`
	ProcessInstanceId   string                           `json:"processInstanceId,omitempty"`
	ProcessDefinitionId string                           `json:"processDefinitionId,omitempty"`
	Variables           []variable.VariableSearchRequest `json:"variable,omitempty"`
}
