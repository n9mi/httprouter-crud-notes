package helper

import (
	"encoding/json"
	"net/http"
)

func EncodeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}

func DecodeJson(r *http.Request, data interface{}) {
	r.Header.Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		panic(err)
	}
}
