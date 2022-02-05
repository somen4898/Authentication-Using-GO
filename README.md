<p align="center"><br><img src="https://avatars.githubusercontent.com/u/76786120?v=4" width="128" height="128" style="border-radius: 50px;" /></p>
<h3 align="center">Ssibrahimbas Go-Auth</h3>
<p align="center">
  An authorization project using mongoDB, JWT and Go.
</p>

### API

<docgen-index>

| Types                  |
|------------------------|
| **[User]("UserType")** |

| HTTP REQUEST | Name                             |
|--------------|----------------------------------|
| **`POST`**   | [`Register`]("#register")        |
| **`POST`**   | [`Login`]("#login")              |
| **`GET`**    | [`Verify Login`]("#verifyLogin") |

</docgen-index>

<docgen-api>

<br>

<hr>

<br>

### UserType

the user's fields and their types.

| Field          | Json     | Bson       | GoLang   | Type     |
|----------------|----------|------------|----------|----------|
| **`name`**     | name     | name       | Name     | string   |
| **`gender`**   | gender   | gender     | Gender   | string   |
| **`age`**      | age      | age        | Age      | int      |
| **`email`**    | email    | email      | Email    | string   |
| **`password`** | password | password   | Password | string   |

In Golang, starting fields with a capital letter makes them public. That's why the first letters of all our fields are capitalized. If we didn't want them to be accessible from outside (private), we would write them in lowercase.

<hr>

<br>

### register 

If the user registers and the transaction is successful, a token is returned is the header of the Response.

| Params         | Type   |
|----------------|--------|
| **`name`**     | string |
| **`gender`**   | string |
| **`age`**      | int    |
| **`email`**    | string |
| **`password`** | string |

#### Example Request

```http request

POST http://localhost:8080/api/register
Content-Type: application/json

{
    "name": "Sami Salih İbrahimbaş",
    "gender": "Man",
    "age": 20,
    "email": "info@samisalihibrahimbas.com.tr",
    "password": "secretPassword"
}

```

#### Returns 

```json
{
    "success": true,
    "message": "Login Successful."
}
```

<br>

<hr>

<br>

### login

The user enters the email and password fields. If there is a registration and the passwords match, the login is provided. In the header of Response, the token is passed as in the registration process.

| Params         | Type   |
|----------------|--------|
| **`email`**    | string |
| **`password`** | string |

#### Example Request

```http request

POST http://localhost:8080/api/login
Content-Type: application/json

{
    "email": "info@samisalihibrahimbas.com.tr",
    "password": "secretPassword"
}

```

#### Returns

```json
{
    "success": true,
    "message": "Login Successful."
}
```

<br>

<hr>

<br>

### verifyLogin

Token is sent in Request's Header. If the token exists and is available, the entry is validated. Otherwise, an error is received.

#### Example Request

```http request

GET http://localhost:8080/api/verifyLogin
Authorization: Bearer dssdcvkşldscasşasdşlsadlşef

```

#### Returns

```json
{
  "success": false,
  "message": "Not Authorized"
}
```

<br>

<hr>

<br>

</docgen-api>