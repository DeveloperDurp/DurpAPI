// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
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
        "/openai/general": {
            "get": {
                "description": "Ask ChatGPT a general question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "openai"
                ],
                "summary": "Gerneral ChatGPT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ask ChatGPT a general question",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/openai/travelagent": {
            "get": {
                "description": "Ask ChatGPT for suggestions as if it was a travel agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "openai"
                ],
                "summary": "Travel Agent ChatGPT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ask ChatGPT for suggestions as a travel agent",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token/generateTokenHandler": {
            "get": {
                "description": "Gets the PSU Data from unraid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "token"
                ],
                "summary": "Generate JWT Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Secret Token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/unraid/powerusage": {
            "get": {
                "description": "Gets the PSU Data from unraid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "unraid"
                ],
                "summary": "Unraid PSU Stats",
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "DurpAPI",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
