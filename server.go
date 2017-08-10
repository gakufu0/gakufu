package main

import (
    "fmt"
    _"net/http"
    "net"
    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _"github.com/mattn/go-sqlite3"
    _"html/template"
    "os"
    "io"
    logrus "github.com/Sirupsen/logrus"
)

func saveFile(path string, c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    if string(path[len(path)-1]) == string("/") {
        path += file.Filename
    }
    dst, err := os.Create(path)
    if err != nil {
        return err
    }
    defer dst.Close()

    if _, err = io.Copy(dst, src); err != nil {
        return err
    }
    return nil
}

func main(){

    db,err := gorm.Open("sqlite3","./database/dev.sqlite3")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    db.SingularTable(true)
    db.LogMode(true)

    db.DB()
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Music{})
    db.AutoMigrate(&Notice{})
    db.AutoMigrate(&Follow{})
    db.AutoMigrate(&Favorite{})
    db.AutoMigrate(&History{})
    db.AutoMigrate(&CreatingMusic{})

    e := echo.New()
    e.Static("/", "assets")

    e.GET("/music/picture/:userid/:imageName",func (c echo.Context) error{
        userid := c.Param("userid")
        imageName := c.Param("imageName")
        return c.File("./assets/music/picture/"+userid+"/"+imageName)
    })

    e.GET("/",func (c echo.Context) error{
      return c.File("./public/create_user.html")
    })
    e.GET("/:userid",func (c echo.Context) error{
        user := new(User)
        userid := c.Param("userid")
        //レコードがないときも0を返してエラーになってない
        if userid == "favicon.ico" {
          return nil
        }
        if err := db.Where("user_id = ?",userid).First(&user).Error; err != nil{
          fmt.Printf("%v",err)
          fmt.Printf("aaaaaa")
          return c.File("./public/create_user.html")
        }
        fmt.Printf("aaa")
        return c.File("./public/index.html")
    })

    e.GET("/:userid/music/new",func (c echo.Context) error{
        var music []Music
        db.Limit(20).Find(&music)
        return c.JSON(200,music)
    })

    e.GET("/:userid/notice", func (c echo.Context) error{
        userid := c.Param("userid")

        var notice []Notice
        db.Where("user_id = ?",userid).Limit(20).Find(&notice)
        return c.JSON(200, notice)
    })

    e.GET("/:userid/fav", func (c echo.Context) error{
        userid := c.Param("userid")

        fav := new(Favorite)
        db.Where("user_id = ?",userid).Limit(20).Find(&fav)
        return c.JSON(200,fav)
    })

    e.GET("/:userid/history", func (c echo.Context) error{
        userid := c.Param("userid")

        history := new(History)
        db.Where("user_id = ?",userid).Limit(20).Find(&history)
        return c.JSON(200, history)
    })

    e.POST("/:userid/music",func (c echo.Context) error{
        music := new(Music)

        music.MusicName = c.FormValue("music_name")
        music.Description = c.FormValue("description")
        music.Tags = c.FormValue("tags")

        userid := c.Param("userid")
        music.CreateUser = userid

        logrus.Warn()
        p,_ := os.Getwd()
        //TODO 毎回作るのは雑魚
        os.Mkdir(p+"/assets/music/picture/"+userid, 0777)
        err := saveFile(p+"/assets/music/picture/" + userid + "/" + music.MusicName,c)
        if err != nil {
            return c.JSON(400, response{Message: "can not saved file", Code: 400})
        }

        if music.MusicName == ""{
            return c.JSON(400, response{Message: "music data is not enough", Code:400})
        }

        music.Content = "/" + userid + "/" + music.MusicName

        db.NewRecord(music)
        db.Create(&music)
        return c.JSON(200, response{Message: "successful music create",Code:200})
    })

    e.POST("/:userid/CreatingMusic", func (c echo.Context) error{
        music := new(CreatingMusic)
        c.Bind(music)
        userid := c.Param("userid")
        music.CreateUser = userid

        if music.MusicId == "" || music.MusicName == "" || music.Content == ""{
            return c.JSON(400, response{Message: "music data is not enough", Code:400})
        }

        db.NewRecord(music)
        db.Create(&music)
        return c.JSON(200, response{Message: "successful music create", Code:200})
    })

    e.POST("/:userid/music/delete/:musicid", func (c echo.Context) error{
        musicid := c.Param("musicid")
        userid  := c.Param("userid")

        music := new(Music)
        ret := db.Where("music_id = ?",musicid).First(&music)
        if ret.Error == nil {
            return c.JSON(404, response{Message: "user not found", Code:404})
        }

        if music.CreateUser != userid {
            return c.JSON(400, response{Message: "not you created",Code:200})
        }

        db.Delete(&music)
        return c.JSON(200, response{Message: "OK", Code: 200})
    })

    e.POST("/createuser",func (c echo.Context) error{
        user := new(User)
        c.Bind(user)
        if user == nil {
            return c.JSON(400, response{Message: "user data is null", Code: 400})
        }

        userDB := new(User)
        db.Where("user_id = ?",user.UserId).First(&userDB)
        if userDB.UserId != "" {
            return c.JSON(400, response{Message: "user_id already used", Code: 400})
        }

        db.NewRecord(user)
        db.Create(&user)
        return c.JSON(200, response{Message: "successful user create", Code:200})
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

    e.POST("/:userid/fav", func (c echo.Context) error{
        fav := new(Favorite)
        favorited := new(Favorite)
        c.Bind(fav)
        fmt.Printf("aaaaaa")
        fmt.Printf("%v",fav.MusicId)

        if err := db.Where("music_id = ?",fav.MusicId).First(&favorited).Error; err != nil{
          userid := c.Param("userid")
          fav.UserId = userid

          db.NewRecord(fav)
          db.Create(&fav)
          return c.JSON(200, response{Message: "successful favorite", Code:200})
        }
        return c.JSON(400,response{Message: "this music is already favorited", Code:400})
    })

    e.POST("/:userid/history", func (c echo.Context) error{
        history := new(History)
        c.Bind(history)

        if history.MusicId == "" || history.Unixtime == 0 {
            return c.JSON(400, response{Message: "music data is not enough", Code:400})
        }

        userid := c.Param("userid")
        history.UserId = userid

        db.NewRecord(history)
        db.Create(&history)
        return c.JSON(200, response{Message: "successfull", Code:200})
    })

    addrs, err := net.InterfaceAddrs()
    fmt.Println(addrs)
    e.Logger.Fatal(e.Start(":1323"))
}
