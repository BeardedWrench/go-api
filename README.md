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
  "FirstName": "Bob",
  "LastName": "Dylan",
  "Email": "email@test.com",
  "Password": "somePassword",
  "Verified": false,
  "Admin": false
}
```

> **Note**
> The `password` field will be encrypted once submitted. 

### Response
```JSON
{
  "ID": "43229b75-4674-4e59-9d25-045d18a80658",
  "FirstName": "Bob",
  "LastName": "Dylan",
  "Password": "9eEfKnFjLeiceoNxwaNu/wt7Mh2RyvknAu5CguEFg3I=",
  "Email": "email@test.com",
  "Verified": false,
  "Admin": false
}
```

## Update a user
### Request
`PUT /users/{id}`
```JSON
{
  "FirstName": "Bob",
  "LastName": "Marley",
  "Email": "email@test.com",
  "Password": "myNewPassword",
  "Verified": false,
  "Admin": false
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
