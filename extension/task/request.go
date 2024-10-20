package task_extension

import "github.com/unionj-cloud/flowablesdk/variable"

type CompleteRequest struct {
	Variables []variable.VariableRequest `json:"variables,omitempty"`
}
