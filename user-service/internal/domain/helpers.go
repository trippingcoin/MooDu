package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, ErrInvalidID
	}
	return objID, nil
}
