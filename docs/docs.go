// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/domains/": {
            "post": {
                "description": "Show all folders in dirrectory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all domains",
                "operationId": "get-all-domains",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelsstruct.Domains"
                        }
                    }
                }
            }
        },
        "/api/domains/{id}": {
            "post": {
                "description": "set_flag is out parameters: 0 - new(set automatic), 1 - active, 2 - ignore",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Set flag to domain",
                "operationId": "set-flag-domain",
                "parameters": [
                    {
                        "description": "set flag",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelsstruct.Flag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/prices/": {
            "get": {
                "description": "Get all prices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all prices",
                "operationId": "get-all-prices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modelsstruct.Prices"
                        }
                    }
                }
            }
        },
        "/api/prices/set_price": {
            "post": {
                "description": "set price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Set price",
                "operationId": "set-price",
                "parameters": [
                    {
                        "description": "required",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelsstruct.Price"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "set_price\": \"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/prices/update_price": {
            "post": {
                "description": "Update price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update price",
                "operationId": "update-price",
                "parameters": [
                    {
                        "description": "required",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modelsstruct.Price"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "update_price\": \"ok\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "modelsstruct.Domain": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "flag": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "modelsstruct.Domains": {
            "type": "object",
            "properties": {
                "domains": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modelsstruct.Domain"
                    }
                }
            }
        },
        "modelsstruct.Flag": {
            "type": "object",
            "required": [
                "set_flag"
            ],
            "properties": {
                "set_flag": {
                    "type": "integer"
                }
            }
        },
        "modelsstruct.Price": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "country1": {
                    "type": "string"
                },
                "country2": {
                    "type": "string"
                },
                "geo": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "money": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "new": {
                    "type": "string"
                },
                "old": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "modelsstruct.Prices": {
            "type": "object",
            "properties": {
                "prices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modelsstruct.Price"
                    }
                }
            }
        }
    }
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
	Version:     "0.01",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "GEO Support API",
	Description: "Helper to add price.js",
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
