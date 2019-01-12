// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-12 23:23:45.509927 +0800 CST m=+0.049873736

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server celler server.",
        "title": "Swagger Example API",
        "termsOfService": "https://razeen.me",
        "contact": {
            "name": "Razeen",
            "url": "https://razeen.me",
            "email": "me@razeen.me"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "127.0.0.1:8080",
    "basePath": "/api/v1",
    "paths": {
        "/file/{id}": {
            "get": {
                "description": "获取文件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "文件处理"
                ],
                "summary": "获取某个文件",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文件ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
            }
        },
        "/hello": {
            "get": {
                "description": "向你说Hello",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "测试SayHello",
                "parameters": [
                    {
                        "type": "string",
                        "description": "人名",
                        "name": "who",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\": \"hello Razeen\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"who are you\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "文件列表",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件处理"
                ],
                "summary": "查看文件列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Files"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登入",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登陆"
                ],
                "summary": "登陆",
                "parameters": [
                    {
                        "type": "string",
                        "default": "admin",
                        "description": "用户名",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"login success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"user or password error\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "上传文件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件处理"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.File"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.File": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "len": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "main.Files": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.File"
                    }
                },
                "len": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}