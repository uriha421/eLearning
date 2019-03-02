# API

`req` = request  
`res` = response  
`stc` = status code  
`msg` = message

## users

##### signup
 ```
POST /user
 ```
`req`	email, username, password  
`res` code(0 or 1), msg, data  
```
{   
　　code: 0,            //code = 0 true, code = 1 false    
　　msg: "ok",    
　　data: {  
       "session": {},     //not sure about the content of session
       "id": "xxx",
       "username": "xxx",
  }   
}    
```
`stc` 201, 400, 500  

##### login
 ```
POST /user/login
 ```
`req`	username(or email), password  
`res` code, msg, data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "session": {},
       "id": "xxx",
       "username": "xxx",
   }
}   
```
`stc` 200, 400, 500  

##### logout
```
DELETE /user/:username
```
`res` code, msg, data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}   
```  
`stc` 204, 400, 401, 403, 500  

##### get user info   
```
GET /user/:username  
```
`res` code, msg ,data
```
{   
　　code: 0,   
　　msg: "ok",   
　　data: {                            
       "id": "xxx",
       "email": "xxx",
       "username": "xxx",    
       "avaster": "xxx",    
       "location": "xxx",    
       "profile": "xxx",    
       "github": "xxx",       
       ...                  //it will be extended.
   }   
}    
```  
`stc` 200, 400, 401, 403, 500    

#### change user info
```
PUT /user/:username
```
`req` id, email, username, avaster, location, profile, github ...
`res` code, msg ,data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}   
```  

## about topics

##### homepage
```
GET /                      
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "homepage": [],      //the list of the newest upload
       "has_more": false     //whether there are next page or not
   }
}
```
`stc` 200, 400, 500

##### texts
```
GET /texts/?tab=aaa&page=x         // x is page number, aaa is tab name.
                                   // no tab show all. no page show 1.
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "texts": [],          //the list of the texts
       //"tab": ‘xx’,        //res the same as the tab in url
       //"tab_list", [],     //the list of all the tab
       "has_more": false     //whether there are next page or not
   }
}
```
`stc` 200, 400, 500  
※tab is like: front-end, back-end, bigdata, algorithm...   
tab and tab_list is not needed right now.

##### notes
```
GET /notes/?tab=aaa&page=x         // x is page number, aaa is tab name.
                                   // no tab show all. no page show 1.
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "notes": [],          //the list of the texts
       //"tab": ‘xx’,        //res the same as the tab in url
       //"tab_list", [],     //the list of all the tab
       "has_more": false     //whether there are next page or not
   }
}
```
`stc` 200, 400, 500

##### questions
```
GET /questions/?tab=aaa&page=x         // x is page number, aaa is tab name.
                                       // no tab show all. no page show 1.
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "questions": [],      //the list of the texts
       //"tab": ‘xx’,        //res the same as the tab in url
       //"tab_list", [],     //the list of all the tab
       "has_more": false     //whether there are next page or not
   }
}
```
`stc` 200, 400, 500

## about user resource
※topics = texts / notes / questions (/ courses)
##### list all resource of someone
```
GET /user/:username/index
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "resource": [],       //the list of the resource that user upload recently
       "has_more": false,     //whether there are next page or not
   }
}
```
`stc` 200, 400, 500

##### list someone's all resource of one topic
```
GET /user/:username/※topics/?page=x
```  
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {  
       "※topics": [],        //the list of the resource that user upload of this topic
       "has_more": false,     //whether there are next page or not
   }
}
```
`stc` 200, 400, 500

##### get one resource of the topic
```
GET /user/:username/※topics/:resource-id
```  
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {
       "resource": {                    
　　　　　　 "title": "xxx",
            "tab": "xxx",
            "content": "xxx",
            "views_num": "1234",
            "comments_num": "12",
            "like_num": "1234",
            "created_at": "2019-03-02 21:40",
         },       
       "comments": [],       //all comments of this resource
   }
}
```
`stc` 200, 400, 500

##### upload one resource
```
POST /user/:username/※topics/:resource-id
```
`req`	title, tab, content  
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 200, 400, 500

##### modify one resource
```
PUT /user/:username/※topics/:resource-id
```
`req`	title, tab, content  
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 200, 400, 500

##### delete one resource
```
DELETE /user/:username/※topics/:resource-id
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 204, 400, 401, 403, 500



## comments
##### Post a comment
```
POST /user/:username/resource/:resource-id/comments
```
`res` comment
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 201, 400, 500

##### Delete a comment
```
DELETE /user/:username/resource/:resource-id/comments
```
`res` comment
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 204, 400, 401, 403, 500


## Other
##### give or cancel a like of one resource
```
GET /user/:username/resource/:resource-id/like
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 200, 400, 500

##### Add or delete collection of one resource
```
GET /user/:username/resource/:resource-id/favourite
```
`res`	code，msg，data
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 200, 400, 500

## message
##### Show all the message
```
GET /user/:username/msg/
```
```
{
　　code: 0,
　　msg: "ok",
　　data: {

   }
}
```
`stc` 200, 400, 500

※ status code  
200	ok  
201	created  
400	bad request  
401	Unauthorized  
403	Forbidden  
500	Internal Server Error  


# ==================================
 the function that will be implemented next time
## Other
##### follow one user
```
GET /user/:username/follow
```
`stc` 200, 400, 500

##### Subscribe to a resource-id
```
GET /user/:username/resource/:resource-id/subscription
```
`stc` 200, 400, 500

## message
##### Show msg from ※somewhere
```
GET /user/:username/msg/※somewhere   
```
`stc` 200, 400, 500  
※somewhere = comments/like/system/follow/favourite/subscription

//##### Show msg of pravite mail   
//GET /user/:username/msg/mail   
//`stc` 200, 400, 500  

## email_auth
