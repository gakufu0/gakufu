package main

import(
    "time"
)

type Music struct{
    ID              uint64      `json:"id" gorm:"primary_key`
    MusicId         string      `json:"music_id"`
    MusicName       string      `json:"music_name"`
    Content         string      `json:"content"`
    Description     string      `json:"description"`
    FavoritedUsers  []string    `json:"favorited_users"`
    CreatedAt       time.Time   `json:"createdAt"`
    CreatedUser     string      `json:"created_user"`
    Tags            []string    `json:"tags"`
    CreatingMusics  []CreatingMusic `json:"creating_musics"`
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
    MusicId         string `json:"music_id"`
    MusicName       string `json:"music_name"`
    Content         string `json:"content"`
    Description     string `json:"description"`
    Tags            []string `json:"tags"`
}
