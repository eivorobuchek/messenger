package mango_db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"user_profile_service/internal/app/models"
)

// AddFriendRequest добавляет запрос на дружбу
func (r *ProfileRepositoryMongo) AddFriendRequest(ctx context.Context, friendID, userID models.UserId) error {
	friendRequest := models.FriendRequest{
		RequesterID: friendID,
		ReceiverID:  userID,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := r.db.friendship.InsertOne(ctx, friendRequest)
	return err
}

// RemoveFriendRequest удаляет запрос на дружбу или друга из коллекции friend_requests
func (r *ProfileRepositoryMongo) RemoveFriendRequest(ctx context.Context, userID, friendID models.UserId) error {
	filter := bson.M{
		"requester_id": friendID,
		"receiver_id":  userID,
	}
	_, err := r.db.friendship.DeleteOne(ctx, filter)
	return err
}

// UpdateFriendRequestStatus обновляет статус запроса на дружбу (accept/decline)
func (r *ProfileRepositoryMongo) UpdateFriendRequestStatus(ctx context.Context, friendID, userId models.UserId, status string) error {
	filter := bson.M{"requester_id": friendID, "receiver_id": userId}
	update := bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}}
	_, err := r.db.friendship.UpdateOne(ctx, filter, update)
	return err
}

// GetFriends возвращает список подтверждённых друзей пользователя
func (r *ProfileRepositoryMongo) GetFriends(ctx context.Context, userID models.UserId) ([]models.Profile, error) {
	// Находим подтвержденные запросы в друзья
	filter := bson.M{
		"$or": []bson.M{
			{"requester_id": userID, "status": "accepted"},
			{"receiver_id": userID, "status": "accepted"},
		},
	}

	cursor, err := r.db.friendship.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var friendIDs []models.UserId
	for cursor.Next(ctx) {
		var request models.FriendRequest
		if err := cursor.Decode(&request); err != nil {
			return nil, err
		}
		if request.RequesterID == userID {
			friendIDs = append(friendIDs, request.ReceiverID)
		} else {
			friendIDs = append(friendIDs, request.RequesterID)
		}
	}

	// Получаем профили друзей
	var friends []models.Profile
	friendFilter := bson.M{"user_id": bson.M{"$in": friendIDs}}
	cursor, err = r.db.profile.Find(ctx, friendFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &friends); err != nil {
		return nil, err
	}
	return friends, nil
}
