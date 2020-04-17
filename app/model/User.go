package model

import "time"

type User struct {
	Id       int64     `db:"id"json:"id"`
	Username string    `db:"username"json:"username"`
	Password string    `db:"password"json:"password"`
	GmtCrate time.Time `db:"gmt_create"json:"gmt_create"`
}
