Echo Framework Example
==========================

Introduction
------------
A Simple REST API example using Echo framework featuring:
- Custom middleware for logging utilizing logrus library and filerotate plugin.
- GORM ORM library with Mysql database

Installation
--------------------
1. Clone/download Repository
2. Restore database schema using ``mdstest.sql`` file.
3. Edit database access credentials config at ``config/app.toml`` or create user and grant access using following SQL command:
``GRANT SELECT,INSERT,UPDATE,DELETE ON 'mdstest'.* TO 'user'@'*' IDENTIFIED BY 'password';``
``FLUSH PRIVILEGES;``
4. Compile code using ``go build``
5. Execute the compiled binary (``./mdstest`` on Linux)

API Documentation
=================

Note: hostname assumed is ``localhost:1323``
### 1. Create User

URL: `http://localhost:1323/User/Create`

METHOD: `HTTP POST`

Variables in request body:
- ``user_id`` : id of user to create
- ``user_name``: name of user
- ``user_password`` : user's password
- ``repeat_password``: repeat user's password

e.g success response:
```javascript
{
    "response_code": "OK",
    "response_message": "User created",
    "data": null
}
````

### 2. Update User

URL: `http://localhost:1323/User/Update`

METHOD: `HTTP PUT`

Variables in request body:
- ``user_id`` : id of user to create
- ``user_name``: name of user
- ``user_status``: status of user (valida values are: A, I, D)
- ``user_password`` : user's password
- ``repeat_password``: repeat user's password

e.g success response:
```javascript
{
    "response_code": "OK",
    "response_message": "User updated",
    "data": null
}
````

### 3. Delete User

URL: `http://localhost:1323/User/Delete/{user_id}`

METHOD: `HTTP DELETE`

Variables in request query:
- ``user_id`` : id of user to delete

e.g success response:
```javascript
{
    "response_code": "OK",
    "response_message": "User deleted",
    "data": null
}
````

### 4. Query user data

URL: `http://localhost:1323/User/{user_id}`

METHOD: `HTTP GET`

Variables in request query:
- ``user_id`` : id of user to query

e.g.success response
```javascript
{
    "response_code": "OK",
    "response_message": "",
    "data": {
        "lastUpdated": "2019-02-17 19:18:59",
        "user_id": "test",
        "user_name": "test user",
        "user_password": "$2a$10$0wzryOZrGtmgvtqRc1MB0OWMB5cRoC0/1/6fj2QE8bO.fa1V.fZ6a",
        "user_status": "A",
        "last_updated": "2019-02-17T19:18:59+07:00",
        "user_setting": [
            {
                "lastUpdated": "2019-02-18 09:42:25",
                "setting_id": 1,
                "UserId": "test",
                "User": null,
                "setting_key": "key123",
                "setting_value": "value123",
                "last_updated": "2019-02-18T09:42:25+07:00"
            },
            {
                "lastUpdated": "2019-02-18 09:42:35",
                "setting_id": 2,
                "UserId": "test",
                "User": null,
                "setting_key": "key456",
                "setting_value": "value456",
                "last_updated": "2019-02-18T09:42:35+07:00"
            }
        ]
    }
}
````

### 5. Create User Setting

URL: `http://localhost:1323/UserSetting/Create`

METHOD: `HTTP POST`

Variables in request body:
- ``user_id`` : id of user who owns setting
- ``setting_key``: key name of setting
- ``setting_value``: value of setting

e.g. success response:
```javascript
{
    "response_code": "OK",
    "response_message": "User setting created",
    "data": null
}
````

### 6. Update User Setting

URL: `http://localhost:1323/UserSetting/Update`

METHOD: `HTTP PUT`

Variables in request body:
- ``user_id`` : id of user who owns setting
- ``setting_key``: key name of setting
- ``setting_value``: value of setting

e.g. success response:
```javascript
{
    "response_code": "OK",
    "response_message": "User setting updated",
    "data": null
}
````

### 7. Delete User Setting

URL: `http://localhost:1323/UserSetting/Delete/{user_id}/{setting_key}`

METHOD: `HTTP DELETE`

Variables in request query:
- ``user_id`` : id of user who owns setting
- ``setting_key``: key name of setting

e.g. success response:
```javascript
{
    "response_code": "OK",
    "response_message": "User setting deleted",
    "data": null
}
````