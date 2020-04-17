package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go-http-example/app/handler"
	"log"
	"net/http"
)

type App struct {
	Db     *sqlx.DB
	Router *mux.Router
}

//初始化
func (a *App) Initialize() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	a.Db = database
	a.loadRoute()
}

//装载路由
func (a *App) loadRoute() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", a.handlerRequest(handler.GetUser)).Methods(http.MethodGet)
	r.HandleFunc("/user", a.handlerRequest(handler.CreateUser)).Methods(http.MethodPost)
	a.Router = r
}

//启动
func (a *App) Run(host string) {
	http.Handle("/", a.Router)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *sqlx.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handlerRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.Db, w, r)
	}
}
