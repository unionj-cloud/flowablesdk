package execution

import (
	"github.com/MasterJoyHunan/flowablesdk/common"
	"github.com/MasterJoyHunan/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest
	common.WithTenant
	Id                           string                           `json:"id,omitempty"`
	ActivityId                   string                           `json:"activityId,omitempty"`
	ProcessDefinitionKey         string                           `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionId          string                           `json:"processDefinitionId,omitempty"`
	ProcessInstanceId            string                           `json:"processInstanceId,omitempty"`
	MessageEventSubscriptionName string                           `json:"messageEventSubscriptionName,omitempty"`
	SignalEventSubscriptionName  string                           `json:"signalEventSubscriptionName,omitempty"`
	ParentId                     string                           `json:"parentId,omitempty"`
	Variables                    []variable.VariableSearchRequest `json:"variable,omitempty"`
	ProcessInstanceVariables     []variable.VariableSearchRequest `json:"processInstanceVariables,omitempty"`
}

type ExecuteRequest struct {
	Action      string                    `json:"action,omitempty"` // signal,messageEventReceived
	MessageName string                    `json:"messageName,omitempty"`
	Variables   []variable.SimpleVariable `json:"variables,omitempty"`
}