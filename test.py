import requests

resp = requests.post("http://127.0.0.1:8080/registration/", json={
    "login": "123123",
    "password": "123123123"
})

print(resp.status_code)
print(resp.content)