package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"user_profile_service/internal/app/controllers/friendship"
	"user_profile_service/internal/app/controllers/profile"
	repository "user_profile_service/internal/app/repositories/mango_db"
	friendship_usecase "user_profile_service/internal/app/usecases/friendsfip"
	profile_usecase "user_profile_service/internal/app/usecases/profile"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// repository
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database("profile_service")

	repo := repository.NewProfileRepository(db)

	// usecases
	profileUsecase := profile_usecase.NewProfileUsecase(profile_usecase.DepsProfile{
		ProfileRepository: repo,
	})

	friendshipUsecase := friendship_usecase.NewFriendshipUsecase(friendship_usecase.FriendshipDeps{
		FriendshipRepository: repo,
	})

	profile.NewProfileHandler(profile.DepsProfile{
		UsecaseProfile: profileUsecase,
	})

	friendship.NewFriendshipHandler(friendship.Deps{
		FriendshipUsecase: friendshipUsecase,
	})

}
