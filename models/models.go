package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name" validate:"required,min=2,max=30"`
	Password        *string            `json:"password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Phone           *string            `json:"phone" validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"create_at"`
	Updated_At      time.Time          `json:"update_at"`
	User_ID         string             `json:"user_id"`
	User_Cart       []ProductUser      `json:"user_cart"`
	Address_Details []Address          `json:"address_details"`
	Order_Status    []Order            `json:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        int                `json:"price"`
	Rating       *uint              `json:"rating"`
	Image        *string            `json:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house" bson:"house_name"`
	Street     *string            `json:"street" bson:"street_name"`
	City       *string            `json:"city" bson:"city_name"`
	Pincode    *string            `json:"pincode" bson:"pin_name"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time          `json:"order_at" bson:"order_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
