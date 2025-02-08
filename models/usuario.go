package models

type Usuario struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Nombre   string `json:"nombre" bson:"nombre"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
