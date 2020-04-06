// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// AuthV1LoginReader is a Reader for the AuthV1Login structure.
type AuthV1LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthV1LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthV1LoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAuthV1LoginNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAuthV1LoginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthV1LoginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthV1LoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthV1LoginOK creates a AuthV1LoginOK with default headers values
func NewAuthV1LoginOK() *AuthV1LoginOK {
	return &AuthV1LoginOK{}
}

/*AuthV1LoginOK handles this case with default header values.

A successful response.
*/
type AuthV1LoginOK struct {
	Payload *service_model.V1Auth
}

func (o *AuthV1LoginOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] authV1LoginOK  %+v", 200, o.Payload)
}

func (o *AuthV1LoginOK) GetPayload() *service_model.V1Auth {
	return o.Payload
}

func (o *AuthV1LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Auth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthV1LoginNoContent creates a AuthV1LoginNoContent with default headers values
func NewAuthV1LoginNoContent() *AuthV1LoginNoContent {
	return &AuthV1LoginNoContent{}
}

/*AuthV1LoginNoContent handles this case with default header values.

No content.
*/
type AuthV1LoginNoContent struct {
	Payload interface{}
}

func (o *AuthV1LoginNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] authV1LoginNoContent  %+v", 204, o.Payload)
}

func (o *AuthV1LoginNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthV1LoginNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthV1LoginForbidden creates a AuthV1LoginForbidden with default headers values
func NewAuthV1LoginForbidden() *AuthV1LoginForbidden {
	return &AuthV1LoginForbidden{}
}

/*AuthV1LoginForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type AuthV1LoginForbidden struct {
	Payload interface{}
}

func (o *AuthV1LoginForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] authV1LoginForbidden  %+v", 403, o.Payload)
}

func (o *AuthV1LoginForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthV1LoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthV1LoginNotFound creates a AuthV1LoginNotFound with default headers values
func NewAuthV1LoginNotFound() *AuthV1LoginNotFound {
	return &AuthV1LoginNotFound{}
}

/*AuthV1LoginNotFound handles this case with default header values.

Resource does not exist.
*/
type AuthV1LoginNotFound struct {
	Payload interface{}
}

func (o *AuthV1LoginNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] authV1LoginNotFound  %+v", 404, o.Payload)
}

func (o *AuthV1LoginNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthV1LoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthV1LoginDefault creates a AuthV1LoginDefault with default headers values
func NewAuthV1LoginDefault(code int) *AuthV1LoginDefault {
	return &AuthV1LoginDefault{
		_statusCode: code,
	}
}

/*AuthV1LoginDefault handles this case with default header values.

An unexpected error response
*/
type AuthV1LoginDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the auth v1 login default response
func (o *AuthV1LoginDefault) Code() int {
	return o._statusCode
}

func (o *AuthV1LoginDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/users/token][%d] AuthV1_Login default  %+v", o._statusCode, o.Payload)
}

func (o *AuthV1LoginDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *AuthV1LoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}