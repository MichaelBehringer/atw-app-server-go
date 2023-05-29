package models

type Person struct {
	PersNo    int    `json:"persNo"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}

type AuthPerson struct {
	PersNo     int    `json:"persNo"`
	Username   string `json:"username"`
	FunctionNo int    `json:"functionNo"`
}
