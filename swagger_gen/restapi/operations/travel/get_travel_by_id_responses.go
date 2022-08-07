// Code generated by go-swagger; DO NOT EDIT.

package travel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/T-jegou/myTravelNotebook/swagger_gen/models"
)

// GetTravelByIDOKCode is the HTTP code returned for type GetTravelByIDOK
const GetTravelByIDOKCode int = 200

/*GetTravelByIDOK successful operation

swagger:response getTravelByIdOK
*/
type GetTravelByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Travel `json:"body,omitempty"`
}

// NewGetTravelByIDOK creates GetTravelByIDOK with default headers values
func NewGetTravelByIDOK() *GetTravelByIDOK {

	return &GetTravelByIDOK{}
}

// WithPayload adds the payload to the get travel by Id o k response
func (o *GetTravelByIDOK) WithPayload(payload *models.Travel) *GetTravelByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get travel by Id o k response
func (o *GetTravelByIDOK) SetPayload(payload *models.Travel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTravelByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTravelByIDBadRequestCode is the HTTP code returned for type GetTravelByIDBadRequest
const GetTravelByIDBadRequestCode int = 400

/*GetTravelByIDBadRequest Invalid ID supplied

swagger:response getTravelByIdBadRequest
*/
type GetTravelByIDBadRequest struct {
}

// NewGetTravelByIDBadRequest creates GetTravelByIDBadRequest with default headers values
func NewGetTravelByIDBadRequest() *GetTravelByIDBadRequest {

	return &GetTravelByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetTravelByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetTravelByIDNotFoundCode is the HTTP code returned for type GetTravelByIDNotFound
const GetTravelByIDNotFoundCode int = 404

/*GetTravelByIDNotFound Travel not found

swagger:response getTravelByIdNotFound
*/
type GetTravelByIDNotFound struct {
}

// NewGetTravelByIDNotFound creates GetTravelByIDNotFound with default headers values
func NewGetTravelByIDNotFound() *GetTravelByIDNotFound {

	return &GetTravelByIDNotFound{}
}

// WriteResponse to the client
func (o *GetTravelByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}