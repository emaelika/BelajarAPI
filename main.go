package main

import (
	"21-api/config"
	"21-api/controller/todo"
	"21-api/controller/user"
	"21-api/model"
	"21-api/routes"
"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("start")
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	m := model.UserModel{Connection: db} // bagian yang menghungkan coding kita ke database / bagian dimana kita ngoding untk ke DB
	c := user.UserController{Model: m}
	// bagian yang menghandle segala hal yang berurusan dengan HTTP / echo
	tm := model.TodoModel{Connection: db}
	tc := todo.TodoController{Model: tm}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, c, tc)
	e.Logger.Fatal(e.Start(":8000"))
}
