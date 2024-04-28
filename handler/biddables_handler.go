package handler

import (
	"bideey/model"
	"bideey/service"
	"encoding/json"
	"net/http"
)

var biddablesService *service.BiddablesService = service.NewBidabblesService()

func CreateBiddable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var biddable model.Biddable
		if err := json.NewDecoder(r.Body).Decode(&biddable); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, serviceErr := biddablesService.CreateBiddable(&biddable)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func UpdateBiddable() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var biddable model.Biddable
		if err := json.NewDecoder(r.Body).Decode(&biddable); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, serviceErr := biddablesService.UpdateBiddable(&biddable)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func QueryBiddables() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		results, serviceErr := biddablesService.QueryBiddables()

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(results)
	}
}

func FetchBiddable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		biddable, serviceErr := biddablesService.GetBiddableById(id)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(biddable)
	}
}

func DeleteBiddable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		serviceErr := biddablesService.DeleteBiddable(id)

		if serviceErr != nil {
			http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
			return
		}
	}
}
