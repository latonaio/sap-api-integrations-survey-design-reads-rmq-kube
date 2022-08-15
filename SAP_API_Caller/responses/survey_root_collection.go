package responses

type SurveyRootCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
