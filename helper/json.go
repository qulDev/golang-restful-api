package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriterResponseToBody(w http.ResponseWriter, response any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(response)

	PanicIfError(err)
}
