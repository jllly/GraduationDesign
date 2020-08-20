package dao

type User struct {
	Phonenumber int `json:"phonenumber"`
	Password    string
	Id          int `json:"id"`
	Position    int
	UserName    string `json:"user_name"`
}


type Rep struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}