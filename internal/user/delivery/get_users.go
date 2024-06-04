package delivery

import (
	"encoding/json"
	"net/http"
)

func (d *UserDelivery) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := d.service.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
