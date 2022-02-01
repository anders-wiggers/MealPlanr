# Routers

Store information about All project routes

## Routes

# `User`

User consist of a user in the system 
```json
{
    "name":"Baron Boss",
    "id":"100001",
    "profilePicture":"http://linktoimg.com/10001.png",
}
```

`POST /user` creates aa user \
`GET /user` with token provices user information \
`UPDATE /user` changes user basic information

# `user/recipes`

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

`POST user/recipes` Create a new recipe \
`GET user/recipes` Return all recipes of a user \

```json
{
    "recipes": [
        100001,
        100002,
        100003,
        100005
    ]
}
```

`GET user/recipe?id=1000010` Return a specific recipy \
`DELETE user/recipes` Delete one or multiple recipes \

# `planner`

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
            [...]
        }
    }
}
```

`GET planner` return a plan for a week \
`POST planner` edit or upload a week plan \

# `planner/shoppinglist`

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

`GET planner/shoppinglist` returns a shopping list for the week. \
`UPDATE planner/shoppinglist` update a shopping list. \
