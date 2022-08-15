package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-survey-design-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSurveyRootCollection(raw []byte, l *logger.Logger) ([]SurveyRootCollection, error) {
	pm := &responses.SurveyRootCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SurveyRootCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	surveyRootCollection := make([]SurveyRootCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		surveyRootCollection = append(surveyRootCollection, SurveyRootCollection{
			ObjectID:                         data.ObjectID,
			ID:                               data.ID,
			Version:                          data.Version,
			Name:                             data.Name,
			NameLanguageCode:                 data.NameLanguageCode,
			NameLanguageCodeText:             data.NameLanguageCodeText,
			TypeCode:                         data.TypeCode,
			TypeCodeText:                     data.TypeCodeText,
			CategoryCode:                     data.CategoryCode,
			CategoryCodeText:                 data.CategoryCodeText,
			ValidFromDate:                    data.ValidFromDate,
			ValidToDate:                      data.ValidToDate,
			LifeCycleStatusCode:              data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:          data.LifeCycleStatusCodeText,
			PreviousAnswersIndicator:         data.PreviousAnswersIndicator,
			PreviousRuntimeProductsIndicator: data.PreviousRuntimeProductsIndicator,
			MatrixViewIndicator:              data.MatrixViewIndicator,
			PaginationIndicator:              data.PaginationIndicator,
			IncludeProductListIndicator:      data.IncludeProductListIndicator,
			SalesOrganisationID:              data.SalesOrganisationID,
			ServiceOrganisationID:            data.ServiceOrganisationID,
			EntityLastChangedOn:              data.EntityLastChangedOn,
			ETag:                             data.ETag,
			TotalMaximumScoreValue:           data.TotalMaximumScoreValue,
			TotalMinimumScoreValue:           data.TotalMinimumScoreValue,
			SurveyCreationDateTime:           data.SurveyCreationDateTime,
		})
	}

	return surveyRootCollection, nil
}
