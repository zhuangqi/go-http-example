package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go-http-example/app/model"
	"net/http"
	"strconv"
	"time"
)

const (
	//获取单个用户SQL
	GetUserSQL = "select id, username, password,gmt_create from user where id =?"
	//创建用户SQL
	CreateUserSQL = "insert into user (username, password, gmt_create) values (?, ?, ?)"
)

// 获取单个用户
func GetUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var user model.User
	err := db.Get(&user, GetUserSQL, id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	marshal, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(marshal)
}

//创建用户
func CreateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := db.Exec(CreateUserSQL, user.Username, user.Password, time.Now())
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	user.Id = id
	marshal, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(marshal)
}
