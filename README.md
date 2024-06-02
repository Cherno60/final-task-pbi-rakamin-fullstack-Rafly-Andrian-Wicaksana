# GoLang API #

## Features ##
1. [Json Web Token](https://jwt.io/) (for Auth Purposes)
2. [Validation](https://github.com/asaskevich/govalidator) (for validate user register)
3. [GodotEnv](https://github.com/joho/godotenv) (All sensitive data is now saved in .env)
4. UUID (instead of using simple increment ID , using UUID is much simpler, and using [Faker](https://github.com/jaswdr/faker) to automatically generate it)
5. Salted Secret_key (the secret key is salted so it much safer)

### Preparations ###
1. Make sure there is .env is present in the project, if there is no .env, rename env file to .env
2. Setup the database username and password in the .env file
4. Make sure there is Mysql database named db_goapi in your local machine

#### Run the project ####
To run the project just run this from cmd inside the project
```golang
go run main.go
```
after you run it, you can access it in `localhost:9090`

## Endpoint ##
Any endpoint of this API can be accessed in here

## User API ##


### User Registration ###
**Endpoint : POST**

URL : /users/register

Request Body:
```json
{
  "username": "username",
  "email": "example@gmail.com",
  "password": "password"
}
```
Response Body success :
```json
{
    "status": 200,
    "messages": "User Created",
    "errors": null
}
```

### User Login ###
**Endpoint : POST**

URL : /users/login

Request Body:
```json
{
	"email": "email@email.com",
	"password": "password"
}
```
Response Body success :
```json
{
    "status": 200,
    "messages": "Logined successfully!, Hello {Username}"
}
```

### User Update ###
**Endpoint : PUT**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /users/edit/{userid}

Request Body:
```json
{
  "username": "newUsername",
  "email": "new@gmail.com",
}
```
Response Body success :
```json
{
    "Errors": null,
    "Messages": {
        "Username": "newUsername",
        "Email": "new@gmail.com",
        "Password": ""
    },
    "Status": 200
}
```

### User Delete ###
**Endpoint : DELETE**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /users/delete/{userid}
Request Parameter: 
`{userid}`
Response Body success :
```json
{
    "status": 200,
    "messages": "User Deleted",
    "errors": null
}{
    "Message": "User Logouted Successfully"
}
```

### User Logout ###
**Endpoint : GET**

URL : /users/logout

Response Body success :
```json
{
    "Message": "User Logouted Successfully"
}
```

## Photo API ##


### Photo Index ###
**Endpoint : GET**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /photos

Response Body success :
```json
{
    "Errors": null,
    "Photo Data": [
        {
            "uuid": "f4abf47c-07f1-4a7e-b250-9af9420af671",
            "title": "Photos",
            "caption": "loresdsadmsd",
            "photo_url": "google.com",
            "user_id": "38157368-9e2d-4889-9c6f-1c6f8a74a6d5",
            "User": null,
            "created_at": "2024-06-01T21:53:18.766+07:00",
            "updated_at": "2024-06-01T21:53:18.766+07:00"
        }
    ],
    "Status": 200
}
```

### Photo Create ###
**Endpoint : POST**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /photos/add

Request Body:
```json
{
   "title" : "Photos",
  "caption": "loresdsadmsd",
  "photo_url": "google.com"
}
```
Response Body success :
```json
{
    "status": 200,
    "messages": "Photo Created",
    "errors": null
}
```

### Photo Edit ###
**Endpoint : PUT**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /photos/edit/{photoid}

Request Body:
```json
{
  "title" : "newTitle",
  "caption": "CaptionNew",
  "photoUrl": "newUrl.com"
}
```
Response Body success :
```json
{
    "Errors": null,
    "Messages": {
        "Title": "newTitle",
        "Caption": "CaptionNew",
        "PhotoUrl": "newUrl.com"
    },
    "Status": 200
}
```


### Photo Delete ###
**Endpoint : DELETE**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /users/delete/{userid}
Request Parameter: 
`{photoid}`
Response Body success :
```json
{
    "Errors": null,
    "Messages": "Photo with caption CaptionNew has been deleted",
    "Status": 200
}
```


