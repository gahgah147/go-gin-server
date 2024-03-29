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
        "/api/file": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查詢全部照片跟影片",
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "查询错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "上傳照片/影片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "傳送人",
                        "name": "Sender",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "上傳檔案",
                        "name": "Image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "機器編號",
                        "name": "Machine_Id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "位置",
                        "name": "Location",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "上傳檔案錯誤",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/file/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改 {id} 上傳檔案",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "上傳檔案ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "傳送人",
                        "name": "Sender",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "上傳檔案",
                        "name": "Image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "機器編號",
                        "name": "Machine_Id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "位置",
                        "name": "Location",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "修改上傳檔案錯誤",
                        "schema": {
                            "type": "string"
                        }
                    }
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
