package mango_db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user_profile_service/internal/app/models"
)

// EditUserProfile редактирует профиль пользователя
func (r *ProfileRepositoryMongo) EditUserProfile(ctx context.Context, user models.Profile) error {
	filter := bson.M{"user_id": user.UserID}
	update := bson.M{"$set": bson.M{
		"nickname":   user.Nickname,
		"bio":        user.Bio,
		"avatar":     user.Avatar,
		"updated_at": time.Now(),
	}}

	_, err := r.db.profile.UpdateOne(ctx, filter, update)
	return err
}

// FindUsersByNickname ищет пользователей по никнейму
func (r *ProfileRepositoryMongo) FindUsersByNickname(ctx context.Context, nickname models.Nickname) ([]models.Profile, error) {
	filter := bson.M{"nickname": primitive.Regex{Pattern: string(nickname), Options: "i"}}
	cursor, err := r.db.profile.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.Profile
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
