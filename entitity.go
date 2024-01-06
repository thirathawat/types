package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoEntity struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"update_at"`
	DeletedAt *time.Time         `bson:"deleted_at"`
}
