package acase_sdk

type Auth struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

type RespError struct {
	Message	string
	Code	string
}
