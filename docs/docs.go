// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "window95pill@gmail.com"
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
        "/investors": {
            "get": {
                "description": "Get investors list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "investors"
                ],
                "summary": "Get investors list",
                "responses": {
                    "200": {
                        "description": "Successfully get investors list",
                        "schema": {
                            "$ref": "#/definitions/response.ListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errorMessage": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errorMessage": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/investors/{id}": {
            "get": {
                "description": "Retrieve investor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "investors"
                ],
                "summary": "Retrieve investor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Investor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieve investor",
                        "schema": {
                            "$ref": "#/definitions/response.RetrieveResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errorMessage": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errorMessage": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.ErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "errorMessage": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.ErrorResponse": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "response.ListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.RetrieveResponse"
                    }
                }
            }
        },
        "response.RetrieveResponse": {
            "type": "object",
            "properties": {
                "cik": {
                    "description": "investor company index key",
                    "type": "string",
                    "example": "1234567890"
                },
                "companyName": {
                    "description": "investor company name",
                    "type": "string",
                    "example": "Company Name"
                },
                "holdingValue": {
                    "description": "total value of holdings",
                    "type": "integer",
                    "example": 1000000
                },
                "id": {
                    "description": "investor id",
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000"
                },
                "name": {
                    "description": "investor name",
                    "type": "string",
                    "example": "John Doe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "wdwb API",
	Description:      "API Server for wdwb",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
