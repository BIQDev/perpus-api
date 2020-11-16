package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type BIQHTTPResponse struct {
	Status  string      `bson:"status" json:"status"`
	Message interface{} `bson:"message" json:"message"`
	Data    interface{} `bson:"data" json:"data"`
}

func WriteResponse( w http.ResponseWriter, code int, status string, message interface{}, data interface{} ) {
	var res BIQHTTPResponse
	res.Status = status
	res.Message = message
	res.Data = data

	result, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error Marshal Response: %v \n", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(result)
}
