package gutils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjID converts a hexadecimal string to a MongoDB ObjectID.
func ObjID(hex string) primitive.ObjectID {
	// Attempt to convert the hex string to an ObjectID
	objID, err := primitive.ObjectIDFromHex(hex)

	if err != nil {
		return primitive.NilObjectID
	}

	return objID
}
