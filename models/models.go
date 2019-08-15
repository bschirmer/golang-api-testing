package models

type Pet struct {
	Id        int      `bson:"id" json:"id"`
	Category  Category `bson:"category" json:"category"`
	Name      string   `bson:"name" json:"name"`
	PhotoUrls []string `bson:"photoUrls" json:"photoUrls"`
	Tags      []Tag    `bson:"tags" json:"tags"`
	Status    string   `bson:"status" json:"status"`
}

type Category struct {
	Id   int    `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type Tag struct {
	Id   int    `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type Order struct {
	Id       int64  `bson:"id" json:"id"`
	PetId    int64  `bson:"petId" json:"petId"`
	Quantity int32  `bson:"quantity" json:"quantity"`
	ShipDate string `bson:"shipDate" json:"shipDate"`
	Status   string `bson:"status" json:"status"`
	Complete bool   `bson:"complete" json:"complete"`
}

type User struct {
	Id         int64  `bson:"id" json:"id"`
	UserName   string `bson:"userName" json:"userName"`
	FirstName  string `bson:"firstName" json:"firstName"`
	LastName   string `bson:"lastName" json:"lastName"`
	Email      string `bson:"email" json:"email"`
	Password   string `bson:"password" json:"password"`
	Phone      string `bson:"phone" json:"phone"`
	UserStatus int32  `bson:"userStatus" json:"userStatus"`
}

type ApiResponse struct {
	Code    int32  `bson:"code" json:"code"`
	Type    string `bson:"type" json:"type"`
	Message string `bson:"message" json:"message"`
}
