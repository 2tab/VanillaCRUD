package utils

import "net/http"

func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	_, err := res.Write(resMessage)
	if err != nil {
		return
	}

}
