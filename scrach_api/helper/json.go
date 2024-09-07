// create json to decode response

package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	// function to read request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)

}


func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	// function to write response body
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	PanicIfError(err)
}