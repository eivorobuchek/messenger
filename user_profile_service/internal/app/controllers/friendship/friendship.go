package friendship

import (
	"context"
	"encoding/json"
	"net/http"
	friendsfip_models "user_profile_service/internal/app/usecases/friendsfip"
)

func (h *Controller) AddFriend(w http.ResponseWriter, r *http.Request) {
	var request friendsfip_models.AddFriendRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.FriendshipUsecase.AddFriend(context.Background(), &request); err != nil {
		http.Error(w, "Failed to add friend", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Controller) RemoveFriend(w http.ResponseWriter, r *http.Request) {
	var request friendsfip_models.DeleteFriendRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.FriendshipUsecase.RemoveFriend(context.Background(), &request); err != nil {
		http.Error(w, "Failed to remove friend", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Controller) UpdateFriendRequest(w http.ResponseWriter, r *http.Request) {
	var request friendsfip_models.UpdateFriendshipStatus

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.FriendshipUsecase.UpdateFriendRequestStatus(context.Background(), &request); err != nil {
		http.Error(w, "Failed to update friendship status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Controller) GetFriends(w http.ResponseWriter, r *http.Request) {
	var request friendsfip_models.GetFriendsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := request.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	friends, err := h.FriendshipUsecase.GetFriends(context.Background(), &request)
	if err != nil {
		http.Error(w, "Failed to retrieve friends", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(friends)
}
