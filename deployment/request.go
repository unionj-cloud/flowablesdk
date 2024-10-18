package deployment

import "github.com/unionj-cloud/flowablesdk/common"

type ListRequest struct {
	common.ListCommonRequest // order allow id,name,deployTime,tenantId
	Name                     string
	NameLike                 string
	Category                 string
	CategoryNotEquals        string
	common.WithTenant
}

type CreateRequest struct {
	FileName string // must end by .bpmn20.xml
	Xml      string
}
