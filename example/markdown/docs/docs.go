// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-05 14:27:50.77484837 +0300 EEST m=+0.030950916

package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/apexswing/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
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
        "/admin/user/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "List users from the store",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.UsersCollection"
                            }
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Add a new user to the store",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
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
                    "admin"
                ],
                "summary": "Add a new user to the store",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    }
                }
            }
        },
        "/admin/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Read user from the store",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.User"
                        }
                    },
                    "400": {
                        "description": "We need ID!!",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "404": {
                        "description": "Can not find ID",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.APIError": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "errorCode": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "api.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "api.UsersCollection": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "email": {
                        "type": "string"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "password": {
                        "type": "string"
                    }
                }
            }
        }
    },
    "tags": [
        {
            "description": "# Admin TAG API documentation\n\n**Admin** functions goes here \n\nFor more info please read [link](/docs/readme.md).\n\n",
            "name": "admin"
        }
    ]
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "# General API documentation\n\n**Warning** this api is not production ready. Use at your own risk.\n\nIn order to re-generate the documentation you need to run\n\n`swago init --md .`",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
