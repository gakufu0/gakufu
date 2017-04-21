import requests

url = "http://localhost:1323/test/delete"

data = {
    "UserId" : "test",
    "AccountName" : "test",
    "Contet" : "aaaa"
}

res = requests.post(url,data=data)
print(res.text)
