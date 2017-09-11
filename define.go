package main

import(
  "time"
  "io"
  _"fmt"
  _"net/http"
  "html/template"
  "github.com/labstack/echo"
  _"github.com/jinzhu/gorm"
)

type Music struct{
  MusicId         string    `json:"music_id"`
  MusicName       string    `json:"music_name"`
  Content         string    `json:"content"; gorm"size:256"`
  Description     string    `json:"description"`
  CreatedAt       time.Time `json:"createdAt"`
  CreateUser      string    `json:"create_user"`
  Tags            string    `json:"tags"`
  Views           uint64    `json:"views"`
  Status          string    `json:"status"`
}

type User struct{
  UserId          string `json:"user_id"`
  AccountName     string `json:"account_name"`
  Password        string `json:"password"`
}

type Notice struct{
  UserId          string  `json:"user_id"`
  Content         string  `json:"content"`
  Unixtime        uint64  `json:"unixtime"`
}

type Follow struct{
  Follow          uint64 `json:"follow"`
  Follower        uint64 `json:"follower"`
}

type Favorite struct{
  UserId          string `json:"user_id"`
  MusicId         string `json:"music_id"`
}

type History struct{
  MusicId         uint64  `json:"music_id"`
  UserId          string  `json:"user_id"`
  Unixtime        uint64  `json:"unixtime" default:0`
}

type response struct {
  Code    int
  Message string
}

type MusicIdTemp struct{
  MusicId string
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
