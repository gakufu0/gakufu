import requests

url = "http://localhost:1323/create"

data = {
    "UserId" : "ueken",
    "AccountName" : "test",
    "Contet" : "aaaa"
}

res = requests.post(url,data=data)
print(res.text)
