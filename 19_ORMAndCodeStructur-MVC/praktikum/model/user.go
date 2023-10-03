package model

type User struct {
	Id    int    `json:"id"		form:"id"`
	Age   int    `json:"age"	form:"age"`
	Email string `json:"email"	form:"email"`
	Name  string `json:"name"	form:"name"`
}
