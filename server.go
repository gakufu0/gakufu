package main

import (
    _"fmt"
    _"net/http"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _"github.com/mattn/go-sqlite3"
    "html/template"
)

func main(){

    db,err := gorm.Open("sqlite3","./database/dev.sqlite3")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    db.DB()
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Music{})
    db.AutoMigrate(&CreatingMusic{})

    e := echo.New()
    t := &Template{
        templates: template.Must(template.ParseGlob("public/*.html")),
    }
    e.Renderer = t

    e.GET("/",func (c echo.Context) error{
        return c.File("./public/test.html")
    })

    e.POST("/create",func (c echo.Context) error{
        user := new(User)
        c.Bind(user)
        if user == nil {
            return c.JSON(400, response{Message: "user data is null", Code: 400})
        }

        db.NewRecord(user)
        db.Create(&user)
        return c.JSON(200, response{Message: "user create successful", Code:200})
    })

    e.POST("/:userid/delete", func (c echo.Context) error{
        user := new(User)

        userid := c.Param("userid")
        ret := db.Where("user_id = ?",userid).First(&user)
        if ret.Error != nil {
            return c.JSON(404, response{Message: "user not found", Code:404})
        }

        db.Delete(&user)

        return c.JSON(200, response{Message: "OK", Code: 200})
    })

    e.Logger.Fatal(e.Start(":1323"))
}

