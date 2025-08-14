package utils

import (
	"log"
	"net/http"
)

func HandleError(err error, msg string, w http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		http.Error(w, msg, http.StatusBadRequest)
	}
}
