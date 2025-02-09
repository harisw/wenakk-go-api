// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Hari Setiawan",
            "email": "haristw16@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "Get all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    }
                }
            }
        },
        "/categories/{slug}": {
            "get": {
                "description": "Get category by slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get category by slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    }
                }
            }
        },
        "/origins": {
            "get": {
                "description": "Get all origins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "origins"
                ],
                "summary": "Get all origins",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Origin"
                            }
                        }
                    }
                }
            }
        },
        "/origins/{slug}": {
            "get": {
                "description": "Get origin by slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "origins"
                ],
                "summary": "Get origin by slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Origin Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Origin"
                        }
                    }
                }
            }
        },
        "/recipes": {
            "get": {
                "description": "Get all recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get all recipes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Recipe"
                            }
                        }
                    }
                }
            }
        },
        "/recipes/category/{slug}": {
            "get": {
                "description": "Get all recipes paginated by its category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get all recipes by Categories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "OrderBy",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CategoryRecipesResponse"
                        }
                    }
                }
            }
        },
        "/recipes/origin/{slug}": {
            "get": {
                "description": "Get all recipes paginated by its origin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get all recipes by Origins",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Origin Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.OriginRecipesResponse"
                        }
                    }
                }
            }
        },
        "/recipes/{recipeId}": {
            "get": {
                "description": "Get recipe by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get recipe by id",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Recipe"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Category": {
            "type": "object",
            "properties": {
                "Img": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Slug": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Origin": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "Slug": {
                    "type": "string"
                }
            }
        },
        "models.Recipe": {
            "type": "object",
            "properties": {
                "Calories": {
                    "type": "number"
                },
                "Category": {
                    "$ref": "#/definitions/models.Category"
                },
                "Date_published": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "Images": {
                    "type": "string",
                    "example": "{\"key\":\"value\"}"
                },
                "Ingredients": {
                    "type": "string"
                },
                "Instructions": {
                    "type": "string",
                    "example": "{\"key\":\"value\"}"
                },
                "Keywords": {
                    "type": "string",
                    "example": "{\"key\":\"value\"}"
                },
                "Name": {
                    "type": "string"
                },
                "Origin": {
                    "$ref": "#/definitions/models.Origin"
                },
                "Protein": {
                    "type": "number"
                },
                "Rating": {
                    "type": "number"
                },
                "Recipe_id": {
                    "type": "integer"
                },
                "Recipe_yield": {
                    "type": "string"
                },
                "Total_time": {
                    "type": "string"
                }
            }
        },
        "types.CategoryRecipesResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.Category"
                },
                "recipes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Recipe"
                    }
                }
            }
        },
        "types.OriginRecipesResponse": {
            "type": "object",
            "properties": {
                "origin": {
                    "$ref": "#/definitions/models.Origin"
                },
                "recipes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Recipe"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Wenakk API",
	Description:      "This is API Documentation for Wenakk API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
