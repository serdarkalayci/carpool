docker compose up -d --build
npm start

1. Docker compose'a eklenecek 
2. degisik env.lar yazilacak prodda calistir 
3. i8s yapilacak
4. input validations yapilacak
5. error handling duzeltilecek


Serdara sorular:

UTC kaydedilmeli?
trip list bossa empty trip donmeli


Extra notlar
soru isaretleri hallet
1. Kullanici kayi, donen hata mesajini parse et!

{
"headers": {
"normalizedNames": {},
"lazyUpdate": null
},
"status": 422,
"statusText": "Unprocessable Entity",
"url": "http://localhost:4200/api/user",
"ok": false,
"name": "HttpErrorResponse",
"message": "Http failure response for http://localhost:4200/api/user: 422 Unprocessable Entity",
"error": [
"Key: 'AddUserRequest.Email' Error: Field validation for 'Email' failed on the 'email' tag",
"Key: 'AddUserRequest.Phone' Error: Field validation for 'Phone' failed on the 'e164' tag"
]
}

{
"headers": {
"normalizedNames": {},
"lazyUpdate": null
},
"status": 500,
"statusText": "Internal Server Error",
"url": "http://localhost:4200/api/user",
"ok": false,
"name": "HttpErrorResponse",
"message": "Http failure response for http://localhost:4200/api/user: 500 Internal Server Error",
"error": {
"error": "email and/or phone number already exists"
}
}

2. Kayit formu validation ekle
3. add trip tarih secmek icin date picker kullan
4. 
