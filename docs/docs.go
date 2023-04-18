// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Bagus Septianto",
            "url": "http://bagusseptianto.com/",
            "email": "mygram@bagusseptianto.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comment/CreateComment": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creating a Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Create a Comment",
                "parameters": [
                    {
                        "description": "Comment Fields",
                        "name": "Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"photo_id\":1,\"message\":\"Lorem ipsum dolor sit amet\"}"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Comment",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/comment/DeleteComment/{CommentID}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deleting a Comment by CommentID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Delete a Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "CommentID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/comment/GetAll": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Read all Comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Comment"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/comment/GetOne/{CommentID}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reading a Comment based on CommentID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Read a Comment by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "CommentID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/comment/UpdateComment/{CommentID}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updating a Comment by CommentID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Update a Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "CommentID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Comment Fields",
                        "name": "Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"photo_id\":1,\"message\":\"Lorem ipsum dolor sit amet\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Comment",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/photo/CreatePhoto": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creating a Photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Create a Photo",
                "parameters": [
                    {
                        "description": "Photo Fields",
                        "name": "Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"title\":\"Lorem Ipsum\",\"caption\":\"Lorem ipsum dolor sit amet\",\"photo_url\":\"https://css.category.id/cat.jpg\"}"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Photo",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/photo/DeletePhoto/{PhotoID}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deleting a Photo by PhotoID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Delete a Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "PhotoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/photo/GetAll": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Read all Photos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Photo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/photo/GetOne/{PhotoID}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reading a Photo based on PhotoID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Read a Photo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "PhotoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/photo/UpdatePhoto/{PhotoID}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updating a Photo by PhotoID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Update a Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "PhotoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Photo Fields",
                        "name": "Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"title\":\"Lorem Ipsum\",\"caption\":\"Lorem ipsum dolor sit amet\",\"photo_url\":\"https://css.category.id/cat.jpg\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Photo",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/socialmedia/CreateSocialMedia": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creating a SocialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMedia"
                ],
                "summary": "Create a SocialMedia",
                "parameters": [
                    {
                        "description": "SocialMedia Fields",
                        "name": "SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"name\":\"bob\",\"social_media_url\":\"https://bagusseptianto.com\"}"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "SocialMedia",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/socialmedia/DeleteSocialMedia/{SocialMediaID}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deleting a SocialMedia by SocialMediaID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMedia"
                ],
                "summary": "Delete a SocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SocialMedia ID",
                        "name": "SocialMediaID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/socialmedia/GetAll": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMedia"
                ],
                "summary": "Read all SocialMedia",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SocialMedia"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/socialmedia/GetOne/{SocialMediaID}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reading a SocialMedia based on SocialMediaID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMedia"
                ],
                "summary": "Read a SocialMedia by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SocialMedia ID",
                        "name": "SocialMediaID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/socialmedia/UpdateSocialMedia/{SocialMediaID}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updating a SocialMedia by SocialMediaID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMedia"
                ],
                "summary": "Update a SocialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SocialMedia ID",
                        "name": "SocialMediaID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "SocialMedia Fields",
                        "name": "SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"name\":\"bob\",\"social_media_url\":\"https://bagusseptianto.com\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "SocialMedia",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Logging in user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "user credential",
                        "name": "UserCredential",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"username\":\"username\",\"password\":\"12345678\"}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "invaild username/password"
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Creating new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "User",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\"username\":\"username\",\"password\":\"12345678\",\"email\":\"email@bagusseptianto.com\",\"age\":21}"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "required": [
                "message",
                "photo_id"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "required": [
                "photo_url",
                "title"
            ],
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "required": [
                "name",
                "social_media_url"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Bearer Token. Please add \"Bearer \" + your token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "My Gram",
	Description:      "DTS FGA KOMINFO Hacktiv8 Golang-6",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}