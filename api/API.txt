# API

`req` = request
`res` = response
`stc` = status code

## users

##### signup
 ```
POST /user
 ```
`req`	email, username, password
`res` success(ture or false), message
`stc` 201, 400, 500

##### email auth
 ```
GET /user/:auth_password
 ```
`res` success, message (if success == true res session(id,username,icon_url))
`stc` 200, 400, 500

##### login
 ```
POST /user/:username
 ```
`req`	email or username
`res` success, message (if success == true res session(id,username,icon_url))
`stc` 200, 400, 500

##### logout
```
DELETE /user/:username
```
`res` success, message
`stc` 204, 400, 401, 403, 500


##### get user info
```
GET /user/:username
```
`res` success, message (if success == true response ※something)
`stc` 200, 400, 401, 403, 500

##### change user info
```
PUT /user/:username
```
`req`	※something
`res` success, message
`stc` 200,400,401,403,500

##### change password
```
PUT /user/:username/password
```
`req`	password, new_password
`res` success, message (if success == true res a new session(id,username,icon_url))
`stc` 200, 400, 500

※something = email, username, icon_url, location, profile, github, level, follower_count, following_count...


## user resource
##### list all resource of someone
```
GET /user/:username/all
```
`req`	page
`stc` 200, 400, 500

##### get one resource
```
GET /user/:username/※topics/:resource-id
```
`stc` 200, 400, 500

##### upload/modify one resource
```
POST /user/:username/※topics/:resource-id
```
`stc` 200, 400, 500

##### delete one resource
```
DELETE /user/:username/※topics/:resource-id
```
`stc` 204,400,401,403,500

※topics = texts / courses / notes / questions


## comments
##### Show comments
```
GET /user/:username/resource/:resource-id/comments
```
`stc` 200, 400, 500

##### Post a comment
```
POST /user/:username/resource/:resource-id/comments
```
`stc` 201, 400, 500

##### Delete a comment
```
DELETE /user/:username/resource/:resource-id/comments
```
`stc` 204, 400, 401, 403, 500


## Message
##### Show all message
```
GET /user/:username/message/
```
`stc` 200, 400, 500

##### Show message of @
```
GET /user/:username/message/at
```
`stc` 200, 400, 500

##### Show message of like
```
GET /user/:username/message/like
```
`stc` 200, 400, 500

##### Show message of system
```
GET /user/:username/message/system
```
`stc` 200, 400, 500

##### Show message of comments
```
GET /user/:username/message/comment
```
`stc` 200, 400, 500

##### Show message of follow
```
GET /user/:username/message/follow
```
`stc` 200, 400, 500

##### Show message of favourite
```
GET /user/:username/message/favourite
```
`stc` 200, 400, 500

##### Show message of subscription
```
GET /user/:username/message/subscription
```
`stc` 200, 400, 500

##### Show message of pravite mail
```
GET /user/:username/message/mail
```
`stc` 200, 400, 500

## Other
##### follow one user
```
GET /user/:username/follow
```
`stc` 200, 400, 500

##### give resource a like
```
GET /user/:username/resource/:resource-id/like
```
`stc` 200, 400, 500

##### Add resource to my collection
```
GET /user/:username/resource/:resource-id/favourite
```
`stc` 200, 400, 500

##### Subscribe to a resource-id
```
GET /user/:username/resource/:resource-id/subscription
```
`stc` 200, 400, 500


※ status code
200	ok
201	created
400	bad request
401	Unauthorized
403	Forbidden
500	Internal Server Error
