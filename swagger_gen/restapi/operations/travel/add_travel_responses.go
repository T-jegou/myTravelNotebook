// Code generated by go-swagger; DO NOT EDIT.

package travel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddTravelOKCode is the HTTP code returned for type AddTravelOK
const AddTravelOKCode int = 200

/*AddTravelOK Succesful operation

swagger:response addTravelOK
*/
type AddTravelOK struct {
}

// NewAddTravelOK creates AddTravelOK with default headers values
func NewAddTravelOK() *AddTravelOK {

	return &AddTravelOK{}
}

// WriteResponse to the client
func (o *AddTravelOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AddTravelMethodNotAllowedCode is the HTTP code returned for type AddTravelMethodNotAllowed
const AddTravelMethodNotAllowedCode int = 405

/*AddTravelMethodNotAllowed Invalid input

swagger:response addTravelMethodNotAllowed
*/
type AddTravelMethodNotAllowed struct {
}

// NewAddTravelMethodNotAllowed creates AddTravelMethodNotAllowed with default headers values
func NewAddTravelMethodNotAllowed() *AddTravelMethodNotAllowed {

	return &AddTravelMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddTravelMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
