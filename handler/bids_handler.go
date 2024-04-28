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

func QueryBids() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		results, serviceErr := bidsService.QueryBids()

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(results)
	}
}

func DeleteBid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		serviceErr := bidsService.DeleteBid(id)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func FetchBid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		bid, serviceErr := bidsService.GetBidById(id)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(bid)
	}
}
