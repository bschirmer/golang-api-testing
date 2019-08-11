package utils

import (
	"encoding/json"
	"strings"

	"github.com/bs/a-jumbo-backend-test/models"
)

func ProcessResponse(code int32, resType string, messages []string) []byte {
	response := new(models.ApiResponse)

	response.Code = code
	response.Type = resType
	if len(messages) == 1 {
		response.Message = messages[0]
	} else {
		response.Message = strings.Join(messages, " , ")
	}
	//Marshal or convert user object back to json and write to response
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return responseJson
}
