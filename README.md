# `bowditch`

API to fetch [Cal Dining menus](https://caldining.berkeley.edu/menus/) in JSON format.

## Routes
`GET /:restaurant` - fetches all meals for the restaurant specified

`GET /:restaurant/:meal` - fetches specified meal for the restaurant specified

## Example
### `GET /crossroads/lunch`
```json
{
  "name": "Lunch",
  "sections": [
    {
      "name": "BIG C OMNI",
      "items": [
        {
          "name": "Chipotle Chicken Pasta"
        },
        {
          "name": "Hawaiian Rolls"
        },
        {
          "name": "Roasted Tomatoes with Garlic"
        }
      ]
    },
    {
      "name": "BIG C VEGAN",
      "items": [
        {
          "name": "Tomato Chipotle Pasta"
        },
        {
          "name": "Dinner Roll"
        },
        {
          "name": "Sauteed Asparagus, Zucchini, and Onions"
        }
      ]
    },
    {
      "name": "GOLDEN GRILL",
      "items": [
        {
          "name": "Chili Cheese Hot Dog"
        },
        {
          "name": "Vegan Chili Dog"
        },
        {
          "name": "Waffle Fries"
        }
      ]
    },
    {
      "name": "BEAR FIT",
      "items": [
        {
          "name": "Hominy and Squash Bear Fit Bowl"
        },
        {
          "name": "Halal Beef"
        },
        {
          "name": "Halal Chicken"
        }
      ]
    },
    {
      "name": "SOUPS",
      "items": [
        {
          "name": "Turkey Chili"
        },
        {
          "name": "Bean Chili"
        }
      ]
    },
    {
      "name": "SALADS",
      "items": [
        {
          "name": "Mixed Green Salad with Balsamic Vinaigrette"
        }
      ]
    },
    {
      "name": "DELI",
      "items": [
        {
          "name": "Peanut Butter and Strawberry Jelly on Wheat bread"
        },
        {
          "name": "Turkey and Swiss Sandwich on Wheat Bread"
        }
      ]
    },
    {
      "name": "DESSERT",
      "items": [
        {
          "name": "Vegan Oatmeal Raisin Cookies"
        },
        {
          "name": "Pecan Pie"
        },
        {
          "name": "Carrot Cake"
        }
      ]
    }
  ]
}
```

### `GET /crossroads`
```json
{
  "data": [
    {
      "name": "Breakfast",
      "sections": [
        {
          "name": "BIG C OMNI",
          "items": [
            {
              "name": "Huevos Rancheros Tostada"
            },
            {
              "name": "Chorizo"
            },
            {
              "name": "Roasted Carrots, Zucchini, Eggplant and Onions"
            }
          ]
        },
        {
          "name": "BIG C VEGAN",
          "items": [
            {
              "name": "Vegan Tostada"
            },
            {
              "name": "Squash, Corn, and Onions Saute"
            }
          ]
        },
        {
          "name": "HOT MORNING GRAINS",
          "items": [
            {
              "name": "Oatmeal"
            },
            {
              "name": "Quinoa Coconut Porridge"
            }
          ]
        },
        {
          "name": "BREAKFAST BREAD",
          "items": [
            {
              "name": "Assorted Mini Danishes"
            },
            {
              "name": "Assorted Scones"
            }
          ]
        },
        {
          "name": "BREAKFAST SPECIAL",
          "items": [
            {
              "name": "Yogurt Parfait Bar"
            }
          ]
        },
        {
          "name": "DAILY RICE",
          "items": [
            {
              "name": "White Rice"
            },
            {
              "name": "Brown Rice"
            }
          ]
        }
      ]
    },
    {
      "name": "Lunch",
      "sections": [
        {
          "name": "BIG C OMNI",
          "items": [
            {
              "name": "Chipotle Chicken Pasta"
            },
            {
              "name": "Hawaiian Rolls"
            },
            {
              "name": "Roasted Tomatoes with Garlic"
            }
          ]
        },
        {
          "name": "BIG C VEGAN",
          "items": [
            {
              "name": "Tomato Chipotle Pasta"
            },
            {
              "name": "Dinner Roll"
            },
            {
              "name": "Sauteed Asparagus, Zucchini, and Onions"
            }
          ]
        },
        {
          "name": "GOLDEN GRILL",
          "items": [
            {
              "name": "Chili Cheese Hot Dog"
            },
            {
              "name": "Vegan Chili Dog"
            },
            {
              "name": "Waffle Fries"
            }
          ]
        },
        {
          "name": "BEAR FIT",
          "items": [
            {
              "name": "Hominy and Squash Bear Fit Bowl"
            },
            {
              "name": "Halal Beef"
            },
            {
              "name": "Halal Chicken"
            }
          ]
        },
        {
          "name": "SOUPS",
          "items": [
            {
              "name": "Turkey Chili"
            },
            {
              "name": "Bean Chili"
            }
          ]
        },
        {
          "name": "SALADS",
          "items": [
            {
              "name": "Mixed Green Salad with Balsamic Vinaigrette"
            }
          ]
        },
        {
          "name": "DELI",
          "items": [
            {
              "name": "Peanut Butter and Strawberry Jelly on Wheat bread"
            },
            {
              "name": "Turkey and Swiss Sandwich on Wheat Bread"
            }
          ]
        },
        {
          "name": "DESSERT",
          "items": [
            {
              "name": "Vegan Oatmeal Raisin Cookies"
            },
            {
              "name": "Pecan Pie"
            },
            {
              "name": "Carrot Cake"
            }
          ]
        }
      ]
    },
    {
      "name": "Dinner",
      "sections": [
        {
          "name": "BIG C OMNI",
          "items": [
            {
              "name": "Roast Turkey with Pineapple Chimichurri"
            },
            {
              "name": "Cilantro Lime Quinoa"
            },
            {
              "name": "Steamed Broccoli"
            }
          ]
        },
        {
          "name": "BIG C VEGAN",
          "items": [
            {
              "name": "Vegan Wild Rice Stuffed Bell Pepper"
            },
            {
              "name": "Cilantro Lime Quinoa"
            },
            {
              "name": "Steamed Broccoli"
            }
          ]
        },
        {
          "name": "GOLDEN GRILL",
          "items": [
            {
              "name": "All Beef Corn Dog"
            },
            {
              "name": "Cauliflower Tempura"
            },
            {
              "name": "Housemade Lemon Pepper Potato Chips"
            }
          ]
        },
        {
          "name": "BEAR FIT",
          "items": [
            {
              "name": "Hominy and Squash Bear Fit Bowl"
            },
            {
              "name": "Halal Beef"
            },
            {
              "name": "Halal Chicken"
            }
          ]
        },
        {
          "name": "SOUPS",
          "items": [
            {
              "name": "Turkey Chili"
            },
            {
              "name": "Bean Chili"
            }
          ]
        },
        {
          "name": "SALADS",
          "items": [
            {
              "name": "Mixed Green Salad with Balsamic Vinaigrette"
            }
          ]
        },
        {
          "name": "DELI",
          "items": [
            {
              "name": "Peanut Butter and Strawberry Jelly on Wheat bread"
            },
            {
              "name": "Turkey and Swiss Sandwich on Wheat Bread"
            }
          ]
        },
        {
          "name": "DESSERT",
          "items": [
            {
              "name": "Vegan Oatmeal Raisin Cookies"
            },
            {
              "name": "Pecan Pie"
            },
            {
              "name": "Carrot Cake"
            }
          ]
        }
      ]
    }
  ]
}
```

## Running locally
```go
$ go mod download
$ go build
$ ./bowditch
```
