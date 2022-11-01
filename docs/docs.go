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
            "name": "ArioDev",
            "url": "https://ariodev.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/childgroup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Create the child group",
                "operationId": "childgroup_create",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ChildGroupCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfull",
                        "schema": {
                            "$ref": "#/definitions/types.ChildGroupUpdate"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Server Faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "503": {
                        "description": "Queue is full",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        },
        "/api/childgroup/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Delete the child group",
                "operationId": "childgroup_delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfull",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Server Faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Update the child group",
                "operationId": "childgroup_update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ChildGroupUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfull",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Server Faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        },
        "/api/endpoint/all": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Get the supported endpoints",
                "operationId": "endpoints_get",
                "responses": {
                    "200": {
                        "description": "Successfull",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.EndpointGetAll"
                            }
                        }
                    }
                }
            }
        },
        "/api/endpoint/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Endpoints"
                ],
                "summary": "Get the supported endpoint",
                "operationId": "endpoints_get_one",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Endpoint id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfull",
                        "schema": {
                            "$ref": "#/definitions/types.EndpointGetAll"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        },
        "/api/middlegroup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Create the middle group",
                "operationId": "middlegroup_create",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.MiddleGroupCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Server faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        },
        "/api/middlegroup/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Delete the middle group",
                "operationId": "middlegroup_delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Update the middle group",
                "operationId": "middlegroup_update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updating data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.MiddleGroupUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Server faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        },
        "/api/topgroup": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Get TopGroups",
                "operationId": "topgroups_get",
                "responses": {
                    "200": {
                        "description": "TopGroup",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.GetTopGroups"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Add a new Top group",
                "operationId": "topgroup_create",
                "parameters": [
                    {
                        "description": "TopGroup name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "TopGroup created",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Creating faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        },
        "/api/topgroup/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Delete TopGroup",
                "operationId": "topgroups_delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TopGroup ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Result",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups Management"
                ],
                "summary": "Update TopGroups",
                "operationId": "topgroups_update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TopGroup ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "TopGroup updated",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    },
                    "500": {
                        "description": "Updating faild",
                        "schema": {
                            "$ref": "#/definitions/types.CreateTopGroupRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "schema.Endpoint": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "currencies": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "formats": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "resolutions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "tier": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "types.ChildGroupCreate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "endpoint_id": {
                    "type": "integer"
                },
                "middlegroup_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.ChildGroupUpdate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.ChildGroups": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "endpoint": {
                    "$ref": "#/definitions/schema.Endpoint"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.CreateTopGroupReq": {
            "description": "Required data to create the Top Group",
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name",
                    "type": "string"
                }
            }
        },
        "types.CreateTopGroupRes": {
            "description": "Top Group response data",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "types.EndpointGetAll": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "currencies": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "formats": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "resolutions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "tier": {
                    "type": "integer"
                }
            }
        },
        "types.GetTopGroups": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "middle_groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.MiddleGroup"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.MiddleGroup": {
            "type": "object",
            "properties": {
                "child_groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.ChildGroups"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.MiddleGroupCreate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "topgroup_id": {
                    "type": "integer"
                }
            }
        },
        "types.MiddleGroupUpdate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Glassnode Service",
	Description:      "Glassnode clone for use in Iran",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
