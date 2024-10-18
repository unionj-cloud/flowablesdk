package task

import (
	"github.com/unionj-cloud/flowablesdk/common"
	"github.com/unionj-cloud/flowablesdk/pkg/timefmt"
	"github.com/unionj-cloud/flowablesdk/variable"
)

type ListRequest struct {
	common.ListCommonRequest                                  // order allow id,name,description,dueDate,createTime,priority,executionId,processInstanceId,tenantId
	Name                           string                     `json:"name,omitempty"`
	NameLike                       string                     `json:"nameLike,omitempty"`
	Description                    string                     `json:"description,omitempty"`
	DescriptionLike                string                     `json:"descriptionLike,omitempty"`
	Priority                       int                        `json:"priority,omitempty"`
	MinimumPriority                int                        `json:"minimumPriority,omitempty"`
	MaximumPriority                int                        `json:"maximumPriority,omitempty"`
	Assignee                       string                     `json:"assignee,omitempty"`
	AssigneeLike                   string                     `json:"assigneeLike,omitempty"`
	Owner                          string                     `json:"owner,omitempty"`
	OwnerLike                      string                     `json:"ownerLike,omitempty"`
	Unassigned                     *bool                      `json:"unassigned,omitempty"`
	DelegationState                string                     `json:"delegationState,omitempty"`
	CandidateUser                  string                     `json:"candidateUser,omitempty"`
	CandidateGroup                 string                     `json:"candidateGroup,omitempty"`
	CandidateGroupIn               []string                   `json:"candidateGroupIn,omitempty"`
	InvolvedUser                   string                     `json:"involvedUser,omitempty"`
	ProcessInstanceId              string                     `json:"processInstanceId,omitempty"`
	TaskDefinitionKey              string                     `json:"taskDefinitionKey,omitempty"`
	TaskDefinitionKeyLike          string                     `json:"taskDefinitionKeyLike,omitempty"`
	ProcessInstanceBusinessKey     string                     `json:"processInstanceBusinessKey,omitempty"`
	ProcessInstanceBusinessKeyLike string                     `json:"processInstanceBusinessKeyLike,omitempty"`
	ProcessDefinitionId            string                     `json:"processDefinitionId,omitempty"`
	ProcessDefinitionKey           string                     `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionKeyLike       string                     `json:"processDefinitionKeyLike,omitempty"`
	ProcessDefinitionName          string                     `json:"processDefinitionName,omitempty"`
	ProcessDefinitionNameLike      string                     `json:"processDefinitionNameLike,omitempty"`
	ExecutionId                    string                     `json:"executionId,omitempty"`
	CreatedOn                      *timefmt.Time              `json:"createdOn,omitempty"`
	CreatedBefore                  *timefmt.Time              `json:"createdBefore,omitempty"`
	CreatedAfter                   *timefmt.Time              `json:"createdAfter,omitempty"`
	DueOn                          *timefmt.Time              `json:"dueOn,omitempty"`
	DueBefore                      *timefmt.Time              `json:"dueBefore,omitempty"`
	DueAfter                       *timefmt.Time              `json:"dueAfter,omitempty"`
	WithoutDueDate                 *timefmt.Time              `json:"withoutDueDate,omitempty"`
	ExcludeSubTasks                *bool                      `json:"excludeSubTasks,omitempty"`
	Active                         *bool                      `json:"active,omitempty"`
	IncludeTaskLocalVariables      *bool                      `json:"includeTaskLocalVariables,omitempty"`
	IncludeProcessVariables        *bool                      `json:"includeProcessVariables,omitempty"`
	CandidateOrAssigned            string                     `json:"candidateOrAssigned,omitempty"`
	Category                       string                     `json:"category,omitempty"`
	TaskVariables                  []variable.VariableRequest `json:"taskVariables,omitempty"`
	ProcessInstanceVariables       []variable.VariableRequest `json:"processInstanceVariables,omitempty"`
	common.WithTenant
}

type UpdateRequest struct {
	Assignee        string `json:"assignee,omitempty"`
	DelegationState string `json:"delegationState,omitempty"`
	Description     string `json:"description,omitempty"`
	DueDate         string `json:"dueDate,omitempty"`
	Name            string `json:"name,omitempty"`
	Owner           string `json:"owner,omitempty"`
	ParentTaskId    string `json:"parentTaskId,omitempty"`
	Priority        int    `json:"priority,omitempty"`
}

type ActionRequest struct {
	Action             string                     `json:"action"`
	Assignee           string                     `json:"assignee,omitempty"`
	FormDefinitionId   string                     `json:"formDefinitionId,omitempty"`
	Outcome            string                     `json:"outcome,omitempty"`
	Variables          []variable.VariableRequest `json:"variables,omitempty"`
	TransientVariables []variable.VariableRequest `json:"transientVariables,omitempty"`
}

type DeleteRequest struct {
	CascadeHistory bool
	DeleteReason   string
}
