package main

import(
    "time"
)

type Music struct{
    ID              uint64 `json:"id" gorm:"primary_key`
    MusicName       string `json:"music_name"`
    Content         string `json:"content"`
    Description     string `json:"description"`
    FavoritedUsers  string `json:"favorited_users"`
    CreatedAt       time.Time `json:"createdAt"`
    CreatedUser     string `json:"created_user"`
    Tags            string `json:"tags"`
}
