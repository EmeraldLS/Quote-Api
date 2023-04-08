package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Quote_Text string             `json:"quote_text,omitempty" bson:"quote_text,omitempty"`
	Author     string             `json:"author,omitempty" bson:"author,omitempty"`
}
