package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-http-example/app/model"
	"net/http"
	"strconv"
	"time"
)

// 获取单个用户
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := model.User{Id: id, Username: "wdz", Password: "123345", GmtCrate: time.Now().Unix()}
	marshal, _ := json.Marshal(user)
	_, _ = w.Write(marshal)
}
