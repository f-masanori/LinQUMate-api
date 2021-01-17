package main

import (
	"fmt"
	gormHandler "linqumate/infra/mysql"
	"linqumate/interface/api"
	mysqlUserRepo "linqumate/repository/mysql"
	"linqumate/usecase"

	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New() //echoというwebフレームワークを使って実装します

	e.Use(middleware.Logger())  //リクエストがきた時にターミナルにリクエスト内容をプリントしてくれる
	e.Use(middleware.Recover()) //panic(例外処理のようなもの)が起きた時に自動的に復旧してくれる

	/* userContorllerを作成 */
	dbhandler := gormHandler.NewMysqlHandler() //ここで違うDBハンドラーにすればデータの保存先を簡単に変更できる
	userRepositroy := mysqlUserRepo.NewMysqlUserRepository(dbhandler)
	userUsecase := usecase.NewUserUsecase(&userRepositroy)
	userController := api.NewUserController(userUsecase)

	e.GET("/a", hello)
	e.GET("/user", userController.UserCreate)

	e.Logger.Fatal(e.Start(":1991"))
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "linqumate"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	fmt.Println(CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println("エラー")
		panic(err.Error())
	}
	return db
}

// Handler
func hello(c echo.Context) error {
	db := gormConnect()
	defer db.Close()

	return c.String(http.StatusOK, "Hello, World!")
}
