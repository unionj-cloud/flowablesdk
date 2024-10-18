package process_definition

import (
	"github.com/unionj-cloud/flowablesdk"
	"github.com/unionj-cloud/flowablesdk/pkg/httpclient"
)

const (
	baseUrl            = "/repository/process-definitions"
	detailUrl          = baseUrl + "/%s"
	ResourceContentUrl = detailUrl + "/resourcedata"
	ModelUrl           = detailUrl + "/model"
	CandidateUrl       = detailUrl + "/identitylinks"
	CandidateDetailUrl = CandidateUrl + "/%s/%s"
)

var (
	ListApi            = flowablesdk.NewApi(httpclient.GET, baseUrl, flowablesdk.ProcessPrefix)
	DetailApi          = flowablesdk.NewApi(httpclient.GET, detailUrl, flowablesdk.ProcessPrefix)
	UpdateApi          = flowablesdk.NewApi(httpclient.PUT, detailUrl, flowablesdk.ProcessPrefix)
	ResourceContentApi = flowablesdk.NewApi(httpclient.GET, ResourceContentUrl, flowablesdk.ProcessPrefix)
	ModelApi           = flowablesdk.NewApi(httpclient.GET, ModelUrl, flowablesdk.ProcessPrefix)
	ListCandidateApi   = flowablesdk.NewApi(httpclient.GET, CandidateUrl, flowablesdk.ProcessPrefix)
	AddCandidateApi    = flowablesdk.NewApi(httpclient.POST, CandidateUrl, flowablesdk.ProcessPrefix)
	DeleteCandidateApi = flowablesdk.NewApi(httpclient.DELETE, CandidateDetailUrl, flowablesdk.ProcessPrefix)
	CandidateDetailApi = flowablesdk.NewApi(httpclient.GET, CandidateDetailUrl, flowablesdk.ProcessPrefix)
)
