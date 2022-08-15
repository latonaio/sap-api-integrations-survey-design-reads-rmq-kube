package sap_api_output_formatter

type SurveyDesign struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	APISchema        string `json:"api_schema"`
	SurveyDesignCode string `json:"survey_design_code"`
	Deleted          bool   `json:"deleted"`
}

type SurveyRootCollection struct {
	ObjectID                         string `json:"ObjectID"`
	ID                               string `json:"ID"`
	Version                          string `json:"Version"`
	Name                             string `json:"Name"`
	NameLanguageCode                 string `json:"NameLanguageCode"`
	NameLanguageCodeText             string `json:"NameLanguageCodeText"`
	TypeCode                         string `json:"TypeCode"`
	TypeCodeText                     string `json:"TypeCodeText"`
	CategoryCode                     string `json:"CategoryCode"`
	CategoryCodeText                 string `json:"CategoryCodeText"`
	ValidFromDate                    string `json:"ValidFromDate"`
	ValidToDate                      string `json:"ValidToDate"`
	LifeCycleStatusCode              string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText          string `json:"LifeCycleStatusCodeText"`
	PreviousAnswersIndicator         bool   `json:"PreviousAnswersIndicator"`
	PreviousRuntimeProductsIndicator bool   `json:"PreviousRuntimeProductsIndicator"`
	MatrixViewIndicator              bool   `json:"MatrixViewIndicator"`
	PaginationIndicator              bool   `json:"PaginationIndicator"`
	IncludeProductListIndicator      bool   `json:"IncludeProductListIndicator"`
	SalesOrganisationID              string `json:"SalesOrganisationID"`
	ServiceOrganisationID            string `json:"ServiceOrganisationID"`
	EntityLastChangedOn              string `json:"EntityLastChangedOn"`
	ETag                             string `json:"ETag"`
	TotalMaximumScoreValue           int    `json:"TotalMaximumScoreValue"`
	TotalMinimumScoreValue           int    `json:"TotalMinimumScoreValue"`
	SurveyCreationDateTime           string `json:"SurveyCreationDateTime"`
}
