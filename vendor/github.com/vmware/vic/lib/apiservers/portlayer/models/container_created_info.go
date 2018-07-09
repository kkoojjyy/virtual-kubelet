package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ContainerCreatedInfo container created info
// swagger:model ContainerCreatedInfo
type ContainerCreatedInfo struct {

	// handle
	// Required: true
	Handle string `json:"handle"`

	// id
	// Required: true
	ID string `json:"id"`
}

// Validate validates this container created info
func (m *ContainerCreatedInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHandle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerCreatedInfo) validateHandle(formats strfmt.Registry) error {

	if err := validate.RequiredString("handle", "body", string(m.Handle)); err != nil {
		return err
	}

	return nil
}

func (m *ContainerCreatedInfo) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}