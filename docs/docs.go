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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addresses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "retrieve a list of all addresses in the system",
                "operationId": "fetch-all-addrs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Address"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "store a new address",
                "operationId": "add-addr",
                "parameters": [
                    {
                        "description": "new address data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.addUpdateAddressBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Address"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            }
        },
        "/addresses/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "retrieve an address by Id",
                "operationId": "fetch-addr",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "update an existing address by Id",
                "operationId": "update-addr",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated address data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.addUpdateAddressBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Address"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "remove an existing address by Id",
                "operationId": "delete-addr",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "retrieve a list of all users in the system",
                "operationId": "fetch-all-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "add a new user",
                "operationId": "add-user",
                "parameters": [
                    {
                        "description": "new user data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.addUpdateUserBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "retrieve a user by Id",
                "operationId": "fetch-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "modify an existing user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new user data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.addUpdateUserBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "users"
                ],
                "summary": "delete a user by Id, including any addresses associated with the user",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            }
        },
        "/users/{id}/addresses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users",
                    "addresses"
                ],
                "summary": "retrieve a list of addresses by the user's Id",
                "operationId": "fetch-addrs-for-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Address"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ApiError": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.addUpdateAddressBody": {
            "type": "object",
            "required": [
                "userId"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "controllers.addUpdateUserBody": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "models.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
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
	Title:            "Go + Gin API Practice",
	Description:      "A small go/gin web app providing a simple REST API for managing users and addresses",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
