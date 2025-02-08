package models 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Nota struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo    string             `json:"titulo" bson:"titulo"`
	Contenido string             `json:"contenido" bson:"contenido"`
}