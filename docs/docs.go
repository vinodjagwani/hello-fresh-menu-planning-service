// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Vinod Jagwani",
            "email": "vkj.agwani86@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/menu-planning-service/api/v1/comments": {
            "post": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Create Comment",
                "parameters": [
                    {
                        "description": "Comment create request body",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CommentCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Comment create response body",
                        "schema": {
                            "$ref": "#/definitions/dto.CommentCreateResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/ingredients": {
            "post": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Create ingredient",
                "parameters": [
                    {
                        "description": "Ingredient create request body",
                        "name": "menuPlan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.IngredientCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ingredient create response body",
                        "schema": {
                            "$ref": "#/definitions/dto.IngredientCreateResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/ingredients/{ingredientId}": {
            "get": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Query ingredient by ingredient id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ingredient ID",
                        "name": "ingredientId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ingredient detail query response body",
                        "schema": {
                            "$ref": "#/definitions/dto.RecipeDetailQueryResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Delete ingredient by ingredient id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ingredient ID",
                        "name": "ingredientId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/login": {
            "post": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User Login Request Body",
                        "name": "menuPlan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User Login Response Body",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/recipes": {
            "post": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipe"
                ],
                "summary": "Create recipe",
                "parameters": [
                    {
                        "description": "Recipe create request body",
                        "name": "recipeCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RecipeCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recipe create response body",
                        "schema": {
                            "$ref": "#/definitions/dto.RecipeCreateResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/recipes/{recipeId}": {
            "get": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipe"
                ],
                "summary": "Query recipe by recipe id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recipe detail query response body",
                        "schema": {
                            "$ref": "#/definitions/dto.RecipeDetailQueryResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipe"
                ],
                "summary": "Update recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Recipe update request body",
                        "name": "recipeUpdateRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RecipeCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recipe update response body",
                        "schema": {
                            "$ref": "#/definitions/dto.RecipeUpdateResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipe"
                ],
                "summary": "Delete recipe by recipe id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/signUp": {
            "post": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "User signup",
                "parameters": [
                    {
                        "description": "User SignUp Request Body",
                        "name": "menuPlan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserSignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User SignUp Response Body",
                        "schema": {
                            "$ref": "#/definitions/dto.UserSignUpResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/weekly-menu-plan": {
            "post": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MenuPlan"
                ],
                "summary": "Create weekly menu plan",
                "parameters": [
                    {
                        "description": "Menu plan request body",
                        "name": "menuPlan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.MenuPlanCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Menu plan response body",
                        "schema": {
                            "$ref": "#/definitions/dto.MenuPlanCreateResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/weekly-menu-plan/{menuPlanId}": {
            "get": {
                "description": "body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MenuPlan"
                ],
                "summary": "Query weekly menu plan by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Menu Plan ID",
                        "name": "menuPlanId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Menu plan response body",
                        "schema": {
                            "$ref": "#/definitions/dto.MenuPlanCreateResponse"
                        }
                    }
                }
            }
        },
        "/menu-planning-service/api/v1/weekly-menu-plans": {
            "get": {
                "description": "body",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MenuPlan"
                ],
                "summary": "Query all weekly menu plans (from current week and in future)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size Number",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "All Menu plans response body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.MenuPlanQueryResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CommentCreateRequest": {
            "type": "object",
            "required": [
                "comment",
                "menuPlanId",
                "userID"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "menuPlanId": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "dto.CommentCreateResponse": {
            "type": "object",
            "required": [
                "commentId"
            ],
            "properties": {
                "commentId": {
                    "type": "string"
                }
            }
        },
        "dto.IngredientCreateRequest": {
            "type": "object",
            "required": [
                "ingredientDescription",
                "ingredientName",
                "ingredientUnit"
            ],
            "properties": {
                "ingredientDescription": {
                    "type": "string"
                },
                "ingredientName": {
                    "type": "string"
                },
                "ingredientUnit": {
                    "type": "integer"
                }
            }
        },
        "dto.IngredientCreateResponse": {
            "type": "object",
            "properties": {
                "ingredientDescription": {
                    "type": "string"
                },
                "ingredientId": {
                    "type": "string"
                },
                "ingredientName": {
                    "type": "string"
                },
                "unit": {
                    "type": "integer"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.MenuPlanCreateRequest": {
            "type": "object",
            "required": [
                "menuPlanName"
            ],
            "properties": {
                "menuPlanDescription": {
                    "type": "string"
                },
                "menuPlanName": {
                    "type": "string"
                },
                "menuPlanRecipeCreateRequest": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.MenuPlanRecipeCreateRequest"
                    }
                },
                "totalCockingTime": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                }
            }
        },
        "dto.MenuPlanCreateResponse": {
            "type": "object",
            "properties": {
                "menuPlanId": {
                    "type": "string"
                },
                "menuPlanName": {
                    "type": "string"
                }
            }
        },
        "dto.MenuPlanIngredientCreateRequest": {
            "type": "object",
            "required": [
                "ingredientDescription",
                "ingredientName"
            ],
            "properties": {
                "ingredientDescription": {
                    "type": "string"
                },
                "ingredientName": {
                    "type": "string"
                },
                "ingredientUnit": {
                    "type": "integer"
                }
            }
        },
        "dto.MenuPlanIngredientQueryDetail": {
            "type": "object",
            "properties": {
                "ingredientDescription": {
                    "type": "string"
                },
                "ingredientId": {
                    "type": "string"
                },
                "ingredientName": {
                    "type": "string"
                },
                "ingredientUnit": {
                    "type": "integer"
                }
            }
        },
        "dto.MenuPlanQueryResponse": {
            "type": "object",
            "properties": {
                "menuPlanDescription": {
                    "type": "string"
                },
                "menuPlanId": {
                    "type": "string"
                },
                "menuPlanName": {
                    "type": "string"
                },
                "menuPlanRecipeQueryDetail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.MenuPlanRecipeQueryDetail"
                    }
                },
                "totalCockingTime": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                }
            }
        },
        "dto.MenuPlanRecipeCreateRequest": {
            "type": "object",
            "required": [
                "recipeName"
            ],
            "properties": {
                "menuPlanIngredientCreateRequest": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.MenuPlanIngredientCreateRequest"
                    }
                },
                "recipeClassification": {
                    "type": "string"
                },
                "recipeDescription": {
                    "type": "string"
                },
                "recipeName": {
                    "type": "string"
                }
            }
        },
        "dto.MenuPlanRecipeQueryDetail": {
            "type": "object",
            "properties": {
                "menuPlanIngredientDetail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.MenuPlanIngredientQueryDetail"
                    }
                },
                "recipeClassification": {
                    "type": "string"
                },
                "recipeDescription": {
                    "type": "string"
                },
                "recipeId": {
                    "type": "string"
                },
                "recipeInstructions": {
                    "type": "string"
                },
                "recipeName": {
                    "type": "string"
                }
            }
        },
        "dto.RecipeCreateRequest": {
            "type": "object",
            "required": [
                "recipeClassification",
                "recipeDescription",
                "recipeInstructions",
                "recipeName"
            ],
            "properties": {
                "recipeClassification": {
                    "type": "string"
                },
                "recipeDescription": {
                    "type": "string"
                },
                "recipeInstructions": {
                    "type": "string"
                },
                "recipeName": {
                    "type": "string"
                }
            }
        },
        "dto.RecipeCreateResponse": {
            "type": "object",
            "properties": {
                "recipeClassification": {
                    "type": "string"
                },
                "recipeDescription": {
                    "type": "string"
                },
                "recipeId": {
                    "type": "string"
                },
                "recipeName": {
                    "type": "string"
                }
            }
        },
        "dto.RecipeDetailQueryResponse": {
            "type": "object",
            "properties": {
                "recipeClassification": {
                    "type": "string"
                },
                "recipeDescription": {
                    "type": "string"
                },
                "recipeId": {
                    "type": "string"
                },
                "recipeName": {
                    "type": "string"
                }
            }
        },
        "dto.RecipeUpdateResponse": {
            "type": "object",
            "properties": {
                "recipeClassification": {
                    "type": "string"
                },
                "recipeDescription": {
                    "type": "string"
                },
                "recipeId": {
                    "type": "string"
                },
                "recipeInstructions": {
                    "type": "string"
                },
                "recipeName": {
                    "type": "string"
                }
            }
        },
        "dto.UserSignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "userName",
                "userType"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                },
                "userType": {
                    "type": "string"
                }
            }
        },
        "dto.UserSignUpResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "HelloFresh Australia MenuPlan RESTFul Apis",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
