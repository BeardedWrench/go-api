# Go API

## Starting the Application

```console
buffalo dev
```



# Endpoints

## Get all users
### Request 
`GET /users`
### Response
```
Array of users in the database.
```

## Get user by ID
### Request
`GET /users/{id}`
### Response
```
A user object
```

## Create new user
### Request
`POST /users`
```JSON
{
  "first_name": "Bob",
  "last_name": "Dylan",
  "email": "email@test.com",
  "password": "somePassword",
  "verified": false,
  "admin": false
}
```

> **Note**
> The `password` field will be encrypted once submitted. 

### Response
```JSON
{
  "id": "43229b75-4674-4e59-9d25-045d18a80658",
  "first_name": "Bob",
  "last_name": "Dylan",
  "password": "9eEfKnFjLeiceoNxwaNu/wt7Mh2RyvknAu5CguEFg3I=",
  "email": "email@test.com",
  "verified": false,
  "admin": false
}
```

## Update a user
### Request
`PUT /users/{id}`
```JSON
{
  "first_name": "Bob",
  "last_name": "Marley",
  "email": "email@test.com",
  "password": "myNewPassword",
  "verified": false,
  "admin": false
}
```
### Response
```JSON
{
  "message": "User with id: c992c506-5402-40b3-8b17-95faef04d90e updated."
}
```

## Delete a user
### Request
`DELETE /users/{id}`
### Response
```JSON
{
  "message": "User with id: c992c506-5402-40b3-8b17-95faef04d90e deleted."
}
```
