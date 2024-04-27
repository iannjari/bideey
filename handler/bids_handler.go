package handler

import (
	"bideey/model"
	"bideey/service"
	"encoding/json"
	"net/http"
)

var bidsService *service.BidsService = service.NewBidsService()

func CreateBid() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var bid model.Bid
		if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, serviceErr := bidsService.CreateBid(&bid)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func UpdateBid() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var bid model.Bid
		if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, serviceErr := bidsService.UpdateBid(&bid)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func DeleteBid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		serviceErr := service.DeleteBid(id)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}
	}
}
