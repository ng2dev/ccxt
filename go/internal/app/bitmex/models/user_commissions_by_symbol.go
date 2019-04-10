// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserCommissionsBySymbol user commissions by symbol
// swagger:model UserCommissionsBySymbol
type UserCommissionsBySymbol struct {

	// symbol
	Symbol *UserCommission `json:"__symbol__,omitempty"`

	// user commissions by symbol
	// Required: true
	UserCommissionsBySymbol map[string]*UserCommission `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *UserCommissionsBySymbol) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// symbol
		Symbol *UserCommission `json:"__symbol__,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv UserCommissionsBySymbol

	rcv.Symbol = stage1.Symbol

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "__symbol__")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]*UserCommission)
		for k, v := range stage2 {
			var toadd *UserCommission
			if err := json.Unmarshal(v, toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.UserCommissionsBySymbol = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m UserCommissionsBySymbol) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// symbol
		Symbol *UserCommission `json:"__symbol__,omitempty"`
	}

	stage1.Symbol = m.Symbol

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.UserCommissionsBySymbol) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.UserCommissionsBySymbol)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this user commissions by symbol
func (m *UserCommissionsBySymbol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	for k := range m.UserCommissionsBySymbol {

		if err := validate.Required(k, "body", m.UserCommissionsBySymbol[k]); err != nil {
			return err
		}
		if val, ok := m.UserCommissionsBySymbol[k]; ok {
			if val != nil {
				if err := val.Validate(formats); err != nil {
					return err
				}
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCommissionsBySymbol) validateSymbol(formats strfmt.Registry) error {

	if swag.IsZero(m.Symbol) { // not required
		return nil
	}

	if m.Symbol != nil {
		if err := m.Symbol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("__symbol__")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserCommissionsBySymbol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCommissionsBySymbol) UnmarshalBinary(b []byte) error {
	var res UserCommissionsBySymbol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}