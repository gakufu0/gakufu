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
    e.GET("/",index)
    e.GET("/hello",helloworld)
    e.POST("/create",create_user)
    e.POST("/:userid/delete",delete_user)
    e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error{
    return c.File( "./public/test.html" )
}

func create_user(c echo.Context) (err error){
    u := new(User)
    if err = c.Bind(u); err != nil {
        return
    }

    db,err := gorm.Open("sqlite3","./database/dev.sqlite3")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    db.DB()
    db.AutoMigrate(&User{})
    db.Create(&u)

    oota := new(User)
    db.First(&oota)
    fmt.Println(oota)

    return c.JSON(200, u)
}

func delete_user(c echo.Context) (err error){
    user := c.Param("userid")
    db,err := gorm.Open("sqlite3","./database/dev.sqlite3")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    db.DB()
    db.Where("user_id = ?",user).Find(&User{}).Delete(&User{})

    return c.String(200,"ok")
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
