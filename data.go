package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "shahidupatel@gmail.com",
		Id:    "user1",
		Name:  "Shahid Patel",
		Token: "123",
	},
	"user2": {
		Email: "shahidpatel@protonmail.com",
		Id:    "user2",
		Name:  "Shahid Patel",
		Token: "310",
	},
}
