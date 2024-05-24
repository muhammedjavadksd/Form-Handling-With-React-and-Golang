package models

type FormRequest struct {
	FullName string `json:"full_name"`
	Phone    int    `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type FormResponse struct {
	Id int32
}
