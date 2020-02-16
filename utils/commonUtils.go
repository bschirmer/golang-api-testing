package utils

import (
	"encoding/json"
	"strings"

	"golang-api-testing/models"
)

func ProcessBadResponse(code int32, resType string, messages []string) []byte {
	response := new(models.ApiResponse)

	response.Code = code
	response.Type = resType
	if len(messages) == 1 {
		response.Message = messages[0]
	} else {
		response.Message = strings.Join(messages, " , ")
	}
	response.Body = "" // this is because we are using the same struct for GOOD and BAD responses
	//Marshal or convert object back to json and write to response
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return responseJson
}

func ProcessResponse(code int32, resType string, object interface{}) []byte {
	response := new(models.ApiResponse)

	response.Code = code
	response.Type = resType
	response.Body = object

	//Marshal or convert object back to json and write to response
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return responseJson
}
