package mango_db

import "go.mongodb.org/mongo-driver/mongo"

type storage struct {
	profile    *mongo.Collection
	friendship *mongo.Collection
}

type ProfileRepositoryMongo struct {
	db storage
}

func NewProfileRepository(db *mongo.Database) *ProfileRepositoryMongo {
	return &ProfileRepositoryMongo{
		db: storage{
			profile:    db.Collection("profile"),
			friendship: db.Collection("friendship"),
		},
	}
}
