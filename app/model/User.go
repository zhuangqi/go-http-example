package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	GmtCrate int64  `json:"gmt_create"`
}
