// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/additem": {
            "post": {
                "description": "Add an item by barcode number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "add"
                ],
                "summary": "add item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "EAN (barcode) number",
                        "name": "ean",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "MHD (expiration date)",
                        "name": "mhd",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Count (amount of item to add)",
                        "name": "count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/deleteitem": {
            "delete": {
                "description": "Delete an item by EAD, MHD, or ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "delete"
                ],
                "summary": "delete item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "EAD (barcode) number",
                        "name": "ead",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MHD (expiration date)",
                        "name": "mhd",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Item ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/getItems/{id}": {
            "get": {
                "description": "Get all food items from the database or a specific item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "get"
                ],
                "summary": "Get all food items or a specific item by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item ID to get a specific item",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Model.Item"
                            }
                        }
                    }
                }
            }
        },
        "/api/viewitems": {
            "get": {
                "description": "View items that expire in the next 3 days or earlier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "view"
                ],
                "summary": "view items expiring soon",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "Model.Item": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "itemBBD": {
                    "type": "integer"
                },
                "itemCount": {
                    "type": "integer"
                },
                "itemEAN": {
                    "type": "integer"
                },
                "itemName": {
                    "type": "string"
                },
                "updatedAt": {
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
