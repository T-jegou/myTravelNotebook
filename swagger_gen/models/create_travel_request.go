// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateTravelRequest create travel request
//
// swagger:model createTravelRequest
type CreateTravelRequest struct {

	// country
	// Required: true
	// Min Length: 1
	Country *string `json:"country"`

	// description
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this create travel request
func (m *CreateTravelRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTravelRequest) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("country", "body", *m.Country, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateTravelRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateTravelRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create travel request based on context it is used
func (m *CreateTravelRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateTravelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTravelRequest) UnmarshalBinary(b []byte) error {
	var res CreateTravelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
