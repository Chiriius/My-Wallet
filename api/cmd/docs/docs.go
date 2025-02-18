// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/login": {
            "post": {
                "summary": "Login User",
                "description": "Login with email and password",
                "parameters": [
                    {
                        "name": "user",
                        "in": "body",
                        "description": "User Login Request",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/LoginUserResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "summary": "Create User",
                "description": "Creates a new user",
                "parameters": [
                    {
                        "name": "user",
                        "in": "body",
                        "description": "User",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created",
                        "schema": {
                            "$ref": "#/definitions/CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "summary": "Get User",
                "description": "Retrieve user information by ID",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User retrieved",
                        "schema": {
                            "$ref": "#/definitions/GetUserResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "summary": "Update User",
                "description": "Updates an existing user",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "user",
                        "in": "body",
                        "description": "User",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated",
                        "schema": {
                            "$ref": "#/definitions/UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "summary": "Delete User",
                "description": "Deletes a user by ID",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "User deleted",
                        "schema": {
                            "$ref": "#/definitions/DeleteUserResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "summary": "Soft Delete User",
                "description": "Marks a user as deleted (soft delete) without removing it from the database",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "User soft deleted",
                        "schema": {
                            "$ref": "#/definitions/DeleteUserResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "LoginUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "passwordExample"
                }
            }
        },
        "LoginUserResponse": {
            "type": "object",
            "properties": {
                "Match": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "CreateUserRequest": {
            "type": "object",
            "properties": {
                "dni": {
                    "type": "integer",
                    "example": 1002842747
                },
                "type_dni": {
                    "type": "string",
                    "example": "CC"
                },
                "name": {
                    "type": "string",
                    "example": "User"
                },
                "email": {
                    "type": "string",
                    "example": "user@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "passwordExample"
                },
                "address": {
                    "type": "string",
                    "example": "Cra 00 #00-00"
                },
                "phone": {
                    "type": "integer",
                    "example": 30179423800
                }
            }
        },
        "CreateUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "UpdateUserRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "123"
                },
                "dni": {
                    "type": "integer",
                    "example": 1002842747
                },
                "type_dni": {
                    "type": "string",
                    "example": "CC"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "newpassword123"
                },
                "address": {
                    "type": "string",
                    "example": "123 Main St, City"
                },
                "phone": {
                    "type": "integer",
                    "example": 1234567890
                }
            }
        },
        "UpdateUserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/User"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "GetUserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/User"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "dni": {
                    "type": "integer"
                },
                "type_dni": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                }
            }
        },
        "DeleteUserResponse": {
            "type": "object",
            "properties": {},
            "description": "Respuesta vacía en caso de éxito, sin cuerpo"
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
    Version:          "1.0",
    Host:             "localhost:8081",
    BasePath:         "/",
    Schemes:          []string{},
    Title:            "My Wallet API",
    Description:      "This is a sample server for a wallet API.",
    SwaggerTemplate:  docTemplate,
}


func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
	