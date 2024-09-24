package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zaki031/shortner/database"
)

var ctx = context.Background()

func GetUrl(w http.ResponseWriter, r *http.Request) {
	client := database.Client
	vars := mux.Vars(r)
	slug := vars["slug"]
	val, err := client.Get(ctx, slug).Result()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println(val)
	json.NewEncoder(w).Encode(val)

}
