// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Interaction interaction
// swagger:model Interaction
type Interaction struct {

	// interaction
	Interaction string `json:"interaction,omitempty"`

	// token id
	TokenID string `json:"token_id,omitempty"`
}

// Validate validates this interaction
func (m *Interaction) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Interaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interaction) UnmarshalBinary(b []byte) error {
	var res Interaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
