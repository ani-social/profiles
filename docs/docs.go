// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://anij.bytecats.codes/tos/",
        "contact": {
            "name": "API Support",
            "url": "http://www.anij.bytecats.codes/support",
            "email": "support@swagger.io"
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
        "/users": {
            "get": {
                "description": "Retrieve a list of all users",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all users",
                "operationId": "get-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User to be created",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "description": "Retrieve a user by their username",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a specific user",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to be fetched",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "404": {
                        "description": "User not found"
                    }
                }
            },
            "put": {
                "description": "Update a user's data by their username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to be updated",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user data",
                        "name": "updatedUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "404": {
                        "description": "User not found"
                    }
                }
            },
            "delete": {
                "description": "Delete a user by their username",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a user",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to be deleted",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "404": {
                        "description": "User not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Achievement": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "main.Profile": {
            "type": "object",
            "properties": {
                "achievements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Achievement"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "philosophy": {
                    "type": "string"
                },
                "socialLinks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Social"
                    }
                }
            }
        },
        "main.Social": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "main.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/main.Profile"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "pipebomb.bytecats.codes",
	BasePath:         "/profiles/api/",
	Schemes:          []string{"https"},
	Title:            "User Profile API",
	Description:      "This is a simple API for managing user profiles.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
