package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID			 primitive.ObjectID      `bson:"_id"`
	Fname         string       			 `bson:"fname" validate:"required"`
	Lname		  string				 `bson:"email" validate:"required"`
	Phone		  string				 `bson:"phone" validate:"required"`
	Password	  string				 `bson:"password" validate:"required"`
	Created_at    time.Time				 `bson:"created_at"`
	Updated_at	  time.Time				 `bson:"updated_at"`
	Role		  []uint               	 `bson:"role" validate:"required,eq=1|eq=2|eq=3,default=1"`
	Orders        []string             	 `bson:"orders" validate:"required"`
	RefreshToken  string				 `bson:"refreshToken" validate:"required"`
	Token         string				 `bson:"token" validate:"required"`

}