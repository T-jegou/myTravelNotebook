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

// UpdateTravelRequest update travel request
//
// swagger:model updateTravelRequest
type UpdateTravelRequest struct {

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

// Validate validates this update travel request
func (m *UpdateTravelRequest) Validate(formats strfmt.Registry) error {
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

func (m *UpdateTravelRequest) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("country", "body", *m.Country, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTravelRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 1); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTravelRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update travel request based on context it is used
func (m *UpdateTravelRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTravelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTravelRequest) UnmarshalBinary(b []byte) error {
	var res UpdateTravelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
