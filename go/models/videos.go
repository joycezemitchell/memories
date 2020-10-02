package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Videos Data Structure
type Videos struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Video       string             `bson:"video"`
	Thumbnail   string             `bson:"thumbnail"`
	CreatedAt   string             `bson:"createdAt"`
	UpdatedAt   string             `bson:"updatedAt"`
}
