package utility

import (
	"encoding/json"
	"net/http"
)

func DecodeBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var obj T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&obj)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return nil, err
	}

	return &obj, nil
}
