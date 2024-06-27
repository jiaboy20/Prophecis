// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FlowJSONResponse flow Json response
// swagger:model FlowJsonResponse
type FlowJSONResponse struct {

	// 实验的工作流描述，所谓的flowjson
	FlowJSON string `json:"flow_json,omitempty"`
}

// Validate validates this flow Json response
func (m *FlowJSONResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlowJSONResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowJSONResponse) UnmarshalBinary(b []byte) error {
	var res FlowJSONResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
