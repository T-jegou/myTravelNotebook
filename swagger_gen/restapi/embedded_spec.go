// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "MyTravelNotebook is an application aim to provide a simple way to create a booktrip.\n",
    "title": "MyTravelNotebook",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/travel": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Returns a list of travels.",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Travel"
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
          "travel"
        ],
        "summary": "Add a new travel to the shelf",
        "operationId": "addTravel",
        "parameters": [
          {
            "description": "Travel object that need to be added to the shelf",
            "name": "nameTravel",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Travel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Succesful operation"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/travel/{id}": {
      "get": {
        "description": "Returns a single travel",
        "produces": [
          "application/json"
        ],
        "tags": [
          "travel"
        ],
        "summary": "Find travel by ID",
        "operationId": "getTravelById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "id's travel to be retrieve",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Travel"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Travel not found"
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
          "travel"
        ],
        "summary": "Updated travel",
        "operationId": "updateTravel",
        "parameters": [
          {
            "type": "integer",
            "description": "id's travel that need to be updated",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Travel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Travel succesfully updated"
          },
          "400": {
            "description": "Invalid Travel supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "travel"
        ],
        "summary": "Deletes a travel",
        "operationId": "deleteTravelById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "id's travel to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Travel successfuly deleted"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Travel not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Travel": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "country": {
          "description": "Country of the travel",
          "type": "string",
          "example": "Lithuania"
        },
        "creationDate": {
          "type": "string",
          "format": "date-time"
        },
        "descriptionTravel": {
          "description": "Notes and resume of your trip !",
          "type": "string",
          "example": "I visit Cracovie, Katowice and Warsaw, but to be honest my favorite city was Gdansk"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "nameTravel": {
          "description": "Name of your trip",
          "type": "string",
          "example": "10 days in Poland"
        }
      }
    }
  },
  "x-tagGroups": [
    {
      "name": "MyTravelNotebook management",
      "tags": [
        "travel"
      ]
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "MyTravelNotebook is an application aim to provide a simple way to create a booktrip.\n",
    "title": "MyTravelNotebook",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/travel": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Returns a list of travels.",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Travel"
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
          "travel"
        ],
        "summary": "Add a new travel to the shelf",
        "operationId": "addTravel",
        "parameters": [
          {
            "description": "Travel object that need to be added to the shelf",
            "name": "nameTravel",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Travel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Succesful operation"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/travel/{id}": {
      "get": {
        "description": "Returns a single travel",
        "produces": [
          "application/json"
        ],
        "tags": [
          "travel"
        ],
        "summary": "Find travel by ID",
        "operationId": "getTravelById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "id's travel to be retrieve",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Travel"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Travel not found"
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
          "travel"
        ],
        "summary": "Updated travel",
        "operationId": "updateTravel",
        "parameters": [
          {
            "type": "integer",
            "description": "id's travel that need to be updated",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Travel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Travel succesfully updated"
          },
          "400": {
            "description": "Invalid Travel supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "travel"
        ],
        "summary": "Deletes a travel",
        "operationId": "deleteTravelById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "id's travel to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Travel successfuly deleted"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Travel not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Travel": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "country": {
          "description": "Country of the travel",
          "type": "string",
          "example": "Lithuania"
        },
        "creationDate": {
          "type": "string",
          "format": "date-time"
        },
        "descriptionTravel": {
          "description": "Notes and resume of your trip !",
          "type": "string",
          "example": "I visit Cracovie, Katowice and Warsaw, but to be honest my favorite city was Gdansk"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "nameTravel": {
          "description": "Name of your trip",
          "type": "string",
          "example": "10 days in Poland"
        }
      }
    }
  },
  "x-tagGroups": [
    {
      "name": "MyTravelNotebook management",
      "tags": [
        "travel"
      ]
    }
  ]
}`))
}
