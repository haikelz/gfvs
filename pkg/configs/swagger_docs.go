package configs

const SwaggerJSON = `
{
    "swagger": "2.0",
    "info": {
        "title": "Go Fiber Vercel Starter",
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
        "version": "1.0"
    },
    "paths": {
        "/": {
            "get": {
                "description": "Get Home",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "Get Home",
                "responses": {
                    "200": {
                        "description": "Get Home Success",
                        "schema": {
                            "$ref": "#/definitions/models.APIHomeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "Get Metrics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Metrics"
                ],
                "summary": "Get Metrics",
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.APIErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "stack": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "models.APIHomeResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "string": {
                    "type": "string"
                }
            }
        }
    }
}
`
