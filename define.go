package main

import(
    "time"
    "io"
    _"fmt"
    _"net/http"
    "html/template"
    "github.com/labstack/echo"
)

type Music struct{
    ID              uint64      `json:"id" gorm:"primary_key`
    MusicId         string      `json:"music_id"`
    MusicName       string      `json:"music_name"`
    Content         string      `json:"content"`
    Description     string      `json:"description"`
    FavoritedUsers  string      `json:"favorited_users"`
    CreatedAt       time.Time   `json:"createdAt"`
    CreateUser      string      `json:"create_user"`
    Tags            string      `json:"tags"`
}

type User struct{
    ID              uint64 `json:"id" gorm:"primary_key"`
    UserId          string `json:"user_id"`
    AccountName     string `json:"account_name"`
    Follow          string `json:"follow"`
    Follower        string `json:"follower"`
    Notification    string `json:"notification"`
    Favorited       string `json:"favorited"`
    History         string `json:"history"`
}

type CreatingMusic struct{
    ID              uint64 `json:"id" gorm:"primary_key`
    MusicId         string `json:"music_id"`
    MusicName       string `json:"music_name"`
    Content         string `json:"content"`
    Description     string `json:"description"`
    CreateUser      string `json:"create_user"`
    Tags            string `json:"tags"`
}

type response struct {
    Code    int
    Message string
}

func authentication(){
    //TODO ここに認証処理
    return;
}

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}
