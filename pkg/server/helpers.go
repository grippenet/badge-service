package server

import(
	"strings"
	"errors"
	"github.com/influenzanet/study-service/pkg/types"
)

var ErrItemResponseNotFound = errors.New("item not found")
var ErrResponseNotFound = errors.New("response object is not found")
var ErrMissingSurveyItemResponse = errors.New("missing survey item response")
var ErrMissingSuryveItem = errors.New("missing survey item")

func findSurveyItemResponse(responses []types.SurveyItemResponse, key string) (responseOfInterest *types.SurveyItemResponse, err error) {
	for _, response := range responses {
		if response.Key == key {
			return &response, nil
		}
	}
	return nil, ErrItemResponseNotFound
}

// Method to retrive one level of the nested response object
func findResponseObject(surveyItem *types.SurveyItemResponse, responseKey string) (responseItem *types.ResponseItem, err error) {
	if surveyItem == nil {
		return responseItem, ErrMissingSuryveItem
	}
	if surveyItem.Response == nil {
		return responseItem, ErrMissingSurveyItemResponse
	}
	for i, k := range strings.Split(responseKey, ".") {
		if i == 0 {
			if surveyItem.Response.Key != k {
				// item not found:
				return responseItem, ErrResponseNotFound
			}
			responseItem = surveyItem.Response
			continue
		}

		found := false
		for _, item := range responseItem.Items {
			if item.Key == k {
				found = true
				responseItem = item
				break
			}
		}
		if !found {
			// item not found:
			return responseItem, ErrResponseNotFound
		}
	}
	return responseItem, nil
}