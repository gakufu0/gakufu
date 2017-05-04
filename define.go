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
    CreatedAt       time.Time   `json:"createdAt"`
    CreateUser      string      `json:"create_user"`
    Tags            string      `json:"tags"`
    Views           uint64      `json:"views"`
}

type User struct{
    ID              uint64 `json:"id" gorm:"primary_key"`
    UserId          string `json:"user_id"`
    AccountName     string `json:"account_name"`
    History         string `json:"history"`
}

type Notice struct{
    ID              uint    `gorm:"primary_key;"`
    UserId          string  `json:"user_id"`
    Content         string  `json:"content"`
	Unixtime        uint    `json:"unixtime"`
}

type Follow struct{
    ID              uint64 `json:"id" gorm:"primary_key"`
    Follow          string `json:"follow"`
    Follower        string `json:"follower"`
}

type Favorite struct{
    ID              uint64 `json:"id" gorm:"primary_key"`
    UserId          string `json:"user_id"`
    MusicId         string `json:"music_id"`
}

type History struct{
    ID              uint64  `json:"id" gorm:"primary_key"`
    MusicId         string  `json:"music_id"`
    UserId          string  `json:"user_id"`
	Unixtime        uint    `json:"unixtime" default:0`
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
