package delivery

import (
	"awesomeProject/internal/user/delivery/dto"
	"encoding/json"
	"net/http"
)

func (d *UserDelivery) AddUser(w http.ResponseWriter, r *http.Request) {
	var addUserDTO dto.AddUserDTO
	if err := json.NewDecoder(r.Body).Decode(&addUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := addUserDTO.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := dto.ConvertToUser(addUserDTO)

	addedUser, err := d.service.AddUser(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedUser)
}
