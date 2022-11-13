package utils

import (
	"fmt"
	"net/http"
)

func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	_, err := res.Write(resMessage)
	if err != nil {
		return
	}

}

func HandleMessage(success bool, message string) []byte {
	responseString := fmt.Sprintf(`{
		"success" : %s,
		"message" : %s
}`, success, message)
	return []byte(responseString)
}
