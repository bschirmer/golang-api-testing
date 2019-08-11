package models

type Pet struct {
	Id        int      `bson:"Id" json:"Id"`
	Category  Category `bson:"Category" json:"Category"`
	Name      string   `bson:"Name" json:"Name"`
	PhotoUrls []string `bson:"PhotoUrls" json:"PhotoUrls"`
	Tags      []Tag    `bson:"Tag" json:"Tag"`
	Status    string   `bson:"Status" json:"Status"`
}

type Category struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type Tag struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type Order struct {
	Id       int64  `json:"Id"`
	PetId    int64  `json:"PetId"`
	Quantity int32  `json:"Quantity"`
	ShipDate string `json:"ShipDate"`
	Status   string `json:"Status"`
	Complete bool   `json:"Complete"`
}

type User struct {
	Id         int64  `json:"Id"`
	UserName   string `json:"UserName"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	Phone      string `json:"Phone"`
	UserStatus int32  `json:"UserStatus"`
}

type ApiResponse struct {
	Code    int32  `json:"Code"`
	Type    string `json:"Type"`
	Message string `json:"Message"`
}
