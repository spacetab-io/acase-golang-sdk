package acase_sdk

type Auth struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

// Auth and request preferences
type Source struct {
	RequestorID				string	`xml:"RequestorID"`
	RequestorPreferences	string	`xml:"RequestorPreferences"`
}

type RequestDetails struct {

}

type AcaseResponse struct {
	Errors	[]RespError
}

type RespError struct {
	Message	string
	Code	string
}
