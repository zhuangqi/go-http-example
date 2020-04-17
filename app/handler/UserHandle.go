package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go-http-example/app/model"
	"net/http"
	"strconv"
)

// 获取单个用户
func GetUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var user model.User
	err := db.Get(&user, "select id, username, password,gmt_create from user where id =?", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	marshal, _ := json.Marshal(user)
	_, _ = w.Write(marshal)
}

//创建用户
func CreateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {

}
