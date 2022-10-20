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
        "/api/group/top": {
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
        "/api/group/top/{id}": {
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
        "types.ChildGroups": {
            "type": "object",
            "properties": {
                "endpoint_id": {
                    "type": "array",
                    "items": {
                        "type": "integer"
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
