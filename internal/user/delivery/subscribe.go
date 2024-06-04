package delivery

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (d *UserDelivery) Subscribe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	subscribedToID, err := strconv.Atoi(vars["subscribedToID"])
	if err != nil {
		http.Error(w, "Invalid subscribed to ID", http.StatusBadRequest)
		return
	}
	err = d.service.Subscribe(r.Context(), userID, subscribedToID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
