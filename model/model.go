package model

type user struct {
	userId   int    `json:"user_id"`
	userName string `json:"user_name"`
	PhoneNo     int    `json:"phone_number"`
	bill_due     int    `json:"bill_due"`
	bill_number      int    `json:"bill_number"`
}
