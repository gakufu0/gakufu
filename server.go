package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _"github.com/mattn/go-sqlite3"
)

func main(){
    e := echo.New()
    e.GET("/",helloworld)
    e.Logger.Fatal(e.Start(":1323"))
}
func helloworld(c echo.Context) error{
    u := new(Music)
    u.MusicName = "くん"
    u.Content = "は"
    u.Description = "ホモ"

    db,err := gorm.Open("sqlite3","./database/dev.sqlite3")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    db.DB()
    db.AutoMigrate(&Music{})
    db.Create(&u)

    oota := new(Music)
    db.First(&oota)
    fmt.Println(oota)

    return c.JSON(http.StatusOK,u)
}
