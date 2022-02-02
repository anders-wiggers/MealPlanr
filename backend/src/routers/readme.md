# Routers

Store information about All project routes

## Routes

# `User`

User consist of a user in the system, the format of a user can be seen here:
```json
{
    "name":"Baron Boss",
    "id":"100001",
    "profilePicture":"http://linktoimg.com/10001.png",
}
```

---

## `POST /user` 
creates a new user in the system  user

**I/O TO BE DETERMINDED**

---
## `GET /user` 
with token provices user information

`BODY`
```json
{
    "bearerToken":"5230-SF20b-&21c1-%8vs1x41sd"
}
```

`OUTPUT`
```json
{
    "name":"Baron Boss",
    "id":"100001",
    "profilePicture":"http://linktoimg.com/10001.png",
}
```
---
## `PATCH /user` 
changes user basic information

`BODY`
```json
{
    "bearerToken":"5230-SF20b-&21c1-%8vs1x41sd",
    "InfoToChange":"NewInformation"
}
```

`OUTPUT`
```json
{
    "status":"succes"
}
```


# `user/recipes`



## `POST /user/recipes` 
Create a new recipe for user given a bearerToken

`BODY`
```json
{
    "id":"100001-10",
    "bearerToken":"5230-SF20b-&21c1-%8vs1x41sd",
    "recipe": {
        "title":"Boiled Goose",
        "uid":"100001-10",
        "date":"22-12-2022",
        "author":"Anders Wiggers",
        "basicInfo": {
            "prebTime":"10 mins",
            "cook":"25 mins",
            "difficulty":"Easy",
        },
        "coverImage":"http://linktoimg.com/10001.png",
        "ingredients":[
            {
                "Goose": {
                    "amount":"2000kg",
                    "purchasable":true
                }
            },
            {
                "Water": {
                    "amount":"200ml",
                    "purchasable":false
                }
            }
        ],
        "Nutrition":{
            "kcal":"100",
            "fat":"12g",
            "saturates":"4g",
            "carbs":"100g",
            "sugars":"0g",
            "fibre":"3g",
            "protein":"20g",
            "salt":"0.5g"
        } ,
        "Method":[
            "warm water",
            "put goose in water",
            "take goose out of water"
        ]
    }
}
```

`OUTPUT`
```json
{
    "status":"succes"
}
````
---
## `GET /user/recipes` 
Return all recipes of a user

```json
{
    "recipes": [
        100001,
        100002,
        100003,
        100004
    ]
}
```

---

## `GET /user/recipe?id=1000010` 
Return a specific recipe

`OUTPUT`
```json
{
    "title":"Boiled Goose",
    "uid":"100001-10",
    "date":"22-12-2022",
    "author":"Anders Wiggers",
    "basicInfo": {
        "prebTime":"10 mins",
        "cook":"25 mins",
        "difficulty":"Easy",
    },
    "coverImage":"http://linktoimg.com/10001.png",
    "ingredients":[
        {
            "Goose": {
                "amount":"2000kg",
                "purchasable":true
            }
        },
        {
            "Water": {
                "amount":"200ml",
                "purchasable":false
            }
        }
    ],
    "Nutrition":{
        "kcal":"100",
        "fat":"12g",
        "saturates":"4g",
        "carbs":"100g",
        "sugars":"0g",
        "fibre":"3g",
        "protein":"20g",
        "salt":"0.5g"
    } ,
    "Method":[
        "warm water",
        "put goose in water",
        "take goose out of water"
    ]
}
```
---
## `DELETE /user/recipes` 
Delete one or multiple recipes 

`INPUT`
```json
{
    "id":"100001-10",
    "bearerToken":"5230-SF20b-&21c1-%8vs1x41sd"
}
```

`OUTPUT`
```json
{
    "status":"succes"
}
```


# `planner`



## `GET /planner` 
Return a plan for a week 


`OUTPUT`
```json 
{
    "week":"1",
    "year":"2022",
    "days":{
        "monday":{
            "date":"02-01-2022",
            "breakfast":{
                "status":"planned",
                "meal_id":110001
            },
            "launch":{
                "status":"notplanned",
                "meal_id":
            },
            "dinner":{
                "status":"planned",
                "meal_id":110001
            }
        },
        "tuesday":{
            
        }
    }
}
```

## `POST /planner` 
Edit or upload a week plan 

`INPUT`
```json
{
    "bearerToken":"5230-SF20b-&21c1-%8vs1x41sd",
    "week":"1",
    "year":"2022",
    "days":{
        "monday":{
            "date":"02-01-2022",
            "breakfast":{
                "status":"planned",
                "meal_id":110001
            },
            "launch":{
                "status":"notplanned",
                "meal_id":
            },
            "dinner":{
                "status":"planned",
                "meal_id":110001
            }
        },
        "tuesday":{
            
        }
    }
}
```

`OUTPUT`
```json
{
    "status":"succes"
}
```


# `planner/shoppinglist`

## `GET /planner/shoppinglist` 
Returns a shopping list for the week.

`OUTPUT`
```json
{
    "Goose":{
        "amount":"500kg",
        "checked":false
    },
    "Speice":{
        "amount":"5g",
        "checked":true
    },
    "Egg":{
        "amount":"12",
        "checked":false
    },
    "Skyr":{
        "amount":"2l",
        "checked":false
    }
}
```


## `PATCH /planner/shoppinglist` 
update a shopping list.
`INPUT`
```json
{
    "Goose":{
        "amount":"500kg",
        "checked":true
    }
}
```

`OUTPUT`
```json
{
    "status":"succes"
}
```



# `login`


## `POST /login` 
given a username and password logs the user in and returns a token.

`INPUT`
```json
[...]/login
{
    "username":"someusername",
    "password":"somesecretpassword"
}
```

`OUTPUT`
```json
{
    "status":"succes",
    "authToken":"5230-SF20b-&21c1-%8vs1x41sd"
}
```