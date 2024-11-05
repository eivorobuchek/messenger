package profile

import (
	"context"
	"encoding/json"
	"net/http"
	profileModels "user_profile_service/internal/app/usecases/profile"
)

func (c *ControllerProfile) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var profile profileModels.EditProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := profile.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.UsecaseProfile

	err := h.ProfileUsecase.(context.Background(), profile)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

