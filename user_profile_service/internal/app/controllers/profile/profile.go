package profile

import (
	"context"
	"encoding/json"
	"net/http"
	"user_profile_service/internal/app/models"
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

	err = c.UsecaseProfile.EditUserProfile(context.Background(), &profile)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ControllerProfile) SearchUserByNickname(w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query().Get("nickname")
	if nickname == "" {
		http.Error(w, "Nickname is required", http.StatusBadRequest)
		return
	}

	users, err := c.UsecaseProfile.FindUsersByNickname(context.Background(), models.Nickname(nickname))
	if err != nil {
		http.Error(w, "Failed to search users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
