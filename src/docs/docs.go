// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/budget-control/api/v1/expense": {
            "get": {
                "description": "Find all expenses",
                "tags": [
                    "Expenses"
                ],
                "summary": "Find all expenses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.ExpenseResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new expense. Obs.: you cannot create two expenses with the same description in a single month.",
                "tags": [
                    "Expenses"
                ],
                "summary": "Create a new expense",
                "parameters": [
                    {
                        "description": "Expense",
                        "name": "expense",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ExpenseRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/budget-control/api/v1/expense/period/{year}/{month}": {
            "get": {
                "description": "Find all expenses by period",
                "tags": [
                    "Expenses"
                ],
                "summary": "Find all expenses by period",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Month",
                        "name": "month",
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
                                "$ref": "#/definitions/entity.ExpenseResponse"
                            }
                        }
                    }
                }
            }
        },
        "/budget-control/api/v1/expense/{id}": {
            "get": {
                "description": "Find expense by id",
                "tags": [
                    "Expenses"
                ],
                "summary": "Find expense by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Expense ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ExpenseResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an expense",
                "tags": [
                    "Expenses"
                ],
                "summary": "Update an expense",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Expense ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Expense",
                        "name": "expense",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ExpenseRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete an expense",
                "tags": [
                    "Expenses"
                ],
                "summary": "Delete an expense",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Expense ID",
                        "name": "id",
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
        "/budget-control/api/v1/health": {
            "get": {
                "description": "return server status",
                "tags": [
                    "Health"
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/budget-control/api/v1/receipt": {
            "get": {
                "description": "Find all receipts",
                "tags": [
                    "Receipts"
                ],
                "summary": "Find all receipts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.ReceiptResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new receipt. Obs.: you cannot create two receipts with the same description in a single month.",
                "tags": [
                    "Receipts"
                ],
                "summary": "Create a new receipt",
                "parameters": [
                    {
                        "description": "Receipt",
                        "name": "receipt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ReceiptRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/budget-control/api/v1/receipt/period/{year}/{month}": {
            "get": {
                "description": "Find all receipts by Period",
                "tags": [
                    "Receipts"
                ],
                "summary": "Find all receipts by Period",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Month",
                        "name": "month",
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
                                "$ref": "#/definitions/entity.ReceiptResponse"
                            }
                        }
                    }
                }
            }
        },
        "/budget-control/api/v1/receipt/{id}": {
            "get": {
                "description": "Find a receipt by id",
                "tags": [
                    "Receipts"
                ],
                "summary": "Find a receipt by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receipt id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ReceiptResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a receipt",
                "tags": [
                    "Receipts"
                ],
                "summary": "Update a receipt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receipt id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Receipt",
                        "name": "receipt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ReceiptRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a receipt",
                "tags": [
                    "Receipts"
                ],
                "summary": "Delete a receipt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receipt id",
                        "name": "id",
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
        "/budget-control/api/v1/summary/{year}/{month}": {
            "get": {
                "description": "get month balance sumary",
                "tags": [
                    "Sumary"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BalanceSumaryResponse"
                        }
                    }
                }
            }
        },
        "/budget-control/api/v1/user": {
            "post": {
                "description": "create a new user",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.UserRequest"
                        }
                    }
                }
            }
        },
        "/budget-control/api/v1/user/authenticate": {
            "post": {
                "description": "authenticate user",
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.ExpenseRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "entity.ExpenseResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "entity.ReceiptRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "entity.ReceiptResponse": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "entity.UserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.BalanceSumaryResponse": {
            "type": "object",
            "properties": {
                "categoryBalance": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "number"
                    }
                },
                "monthBalance": {
                    "type": "number"
                },
                "totalExpense": {
                    "type": "number"
                },
                "totalReceipt": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.3.1",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Alura Backend Challenge 2nd Edition API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}