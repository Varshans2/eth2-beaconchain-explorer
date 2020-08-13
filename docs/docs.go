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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/healthz": {
            "get": {
                "description": "Health endpoint for montitoring if the explorer is in sync",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health of the explorer",
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
        "/api/v1/block/{slotOrHash}": {
            "get": {
                "description": "Returns a block by its slot or root hash",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot or root hash or the string latest",
                        "name": "slotOrHash",
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
        "/api/v1/block/{slot}/attestations": {
            "get": {
                "description": "Returns the attestations included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the attestations included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
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
        "/api/v1/block/{slot}/attesterslashings": {
            "get": {
                "description": "Returns the attester slashings included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the attester slashings included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
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
        "/api/v1/block/{slot}/deposits": {
            "get": {
                "description": "Returns the deposits included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the deposits included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
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
        "/api/v1/block/{slot}/proposerslashings": {
            "get": {
                "description": "Returns the proposer slashings included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the proposer slashings included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
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
        "/api/v1/block/{slot}/voluntaryexits": {
            "get": {
                "description": "Returns the voluntary exits included in a specific block",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Get the voluntary exits included in a specific block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block slot",
                        "name": "slot",
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
        "/api/v1/epoch/{epoch}": {
            "get": {
                "description": "Returns information for a specified epoch by the epoch number or the latest epoch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Epoch"
                ],
                "summary": "Get epoch by number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Epoch number or the string latest",
                        "name": "epoch",
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
        "/api/v1/epoch/{epoch}/blocks": {
            "get": {
                "description": "Returns all blocks for a specified epoch",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Epoch"
                ],
                "summary": "Get epoch blocks by epoch number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Epoch number or the string latest",
                        "name": "epoch",
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
        "/api/v1/eth1deposit/{txhash}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Eth1"
                ],
                "summary": "Get an eth1 deposit by its eth1 transaction hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Eth1 transaction hash",
                        "name": "txhash",
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
        "/api/v1/validator/{index}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get a validator by its index or public key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indices, comma separated",
                        "name": "index",
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
        "/api/v1/validator/{index}/balancehistory": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get the balance history (last 100 epochs) of a validator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indices, comma separated",
                        "name": "index",
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
        "/api/v1/validator/{index}/performance": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validator"
                ],
                "summary": "Get the current performance of a validator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Up to 100 validator indices, comma separated",
                        "name": "index",
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Beaconcha.in ETH2 API",
	Description: "High performance API for querying information from the Ethereum 2.0 beacon chain\nThe API is currently free to use. A fair use policy applies. Calls are rate limited to\n10 requests / 1 minute / IP. All API results are cached for 1 minute.\nIf you required a higher usage plan please get in touch with us at support@beaconcha.in.",
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
