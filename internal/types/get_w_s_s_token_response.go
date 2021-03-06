// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetWSSTokenResponse get w s s token response
//
// swagger:model getWSSTokenResponse
type GetWSSTokenResponse struct {

	// token
	// Example: f0f8e8f8-f0f8-f0f8-f0f8-f0f8f0f8f0f8
	// Required: true
	// Format: uuid
	Token *strfmt.UUID `json:"token"`
}

// Validate validates this get w s s token response
func (m *GetWSSTokenResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWSSTokenResponse) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if err := validate.FormatOf("token", "body", "uuid", m.Token.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get w s s token response based on context it is used
func (m *GetWSSTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetWSSTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWSSTokenResponse) UnmarshalBinary(b []byte) error {
	var res GetWSSTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
