package models

type Person struct {
	PersNo    int    `json:"persNo"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}

type PersonExtra struct {
	PersNo       int    `json:"persNo"`
	Lastname     string `json:"lastname"`
	Firstname    string `json:"firstname"`
	Username     string `json:"username"`
	FunctionNo   int    `json:"functionNo"`
	FunctionName string `json:"functionName"`
	CityNo       int    `json:"cityNo"`
	CityName     string `json:"cityName"`
}

type AuthPerson struct {
	PersNo     int    `json:"persNo"`
	Username   string `json:"username"`
	FunctionNo int    `json:"functionNo"`
}

type Function struct {
	FunctionNo   int    `json:"functionNo"`
	FunctionName string `json:"functionName"`
}
