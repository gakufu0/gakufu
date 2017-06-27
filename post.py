import requests

url = "http://localhost:1323/mstn/music"

data = {
    "MusicId" : "ueken",
    "MusicName" : "test",
    "Content" : open("music_box.png","rb")
}
res = requests.post(url,data=data)
print(res.text)
