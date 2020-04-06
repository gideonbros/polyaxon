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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ProjectSearchesV1ListProjectSearchNamesReader is a Reader for the ProjectSearchesV1ListProjectSearchNames structure.
type ProjectSearchesV1ListProjectSearchNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectSearchesV1ListProjectSearchNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectSearchesV1ListProjectSearchNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectSearchesV1ListProjectSearchNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewProjectSearchesV1ListProjectSearchNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectSearchesV1ListProjectSearchNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectSearchesV1ListProjectSearchNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectSearchesV1ListProjectSearchNamesOK creates a ProjectSearchesV1ListProjectSearchNamesOK with default headers values
func NewProjectSearchesV1ListProjectSearchNamesOK() *ProjectSearchesV1ListProjectSearchNamesOK {
	return &ProjectSearchesV1ListProjectSearchNamesOK{}
}

/*ProjectSearchesV1ListProjectSearchNamesOK handles this case with default header values.

A successful response.
*/
type ProjectSearchesV1ListProjectSearchNamesOK struct {
	Payload *service_model.V1ListSearchesResponse
}

func (o *ProjectSearchesV1ListProjectSearchNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/searches/names][%d] projectSearchesV1ListProjectSearchNamesOK  %+v", 200, o.Payload)
}

func (o *ProjectSearchesV1ListProjectSearchNamesOK) GetPayload() *service_model.V1ListSearchesResponse {
	return o.Payload
}

func (o *ProjectSearchesV1ListProjectSearchNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListSearchesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1ListProjectSearchNamesNoContent creates a ProjectSearchesV1ListProjectSearchNamesNoContent with default headers values
func NewProjectSearchesV1ListProjectSearchNamesNoContent() *ProjectSearchesV1ListProjectSearchNamesNoContent {
	return &ProjectSearchesV1ListProjectSearchNamesNoContent{}
}

/*ProjectSearchesV1ListProjectSearchNamesNoContent handles this case with default header values.

No content.
*/
type ProjectSearchesV1ListProjectSearchNamesNoContent struct {
	Payload interface{}
}

func (o *ProjectSearchesV1ListProjectSearchNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/searches/names][%d] projectSearchesV1ListProjectSearchNamesNoContent  %+v", 204, o.Payload)
}

func (o *ProjectSearchesV1ListProjectSearchNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1ListProjectSearchNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1ListProjectSearchNamesForbidden creates a ProjectSearchesV1ListProjectSearchNamesForbidden with default headers values
func NewProjectSearchesV1ListProjectSearchNamesForbidden() *ProjectSearchesV1ListProjectSearchNamesForbidden {
	return &ProjectSearchesV1ListProjectSearchNamesForbidden{}
}

/*ProjectSearchesV1ListProjectSearchNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ProjectSearchesV1ListProjectSearchNamesForbidden struct {
	Payload interface{}
}

func (o *ProjectSearchesV1ListProjectSearchNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/searches/names][%d] projectSearchesV1ListProjectSearchNamesForbidden  %+v", 403, o.Payload)
}

func (o *ProjectSearchesV1ListProjectSearchNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1ListProjectSearchNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1ListProjectSearchNamesNotFound creates a ProjectSearchesV1ListProjectSearchNamesNotFound with default headers values
func NewProjectSearchesV1ListProjectSearchNamesNotFound() *ProjectSearchesV1ListProjectSearchNamesNotFound {
	return &ProjectSearchesV1ListProjectSearchNamesNotFound{}
}

/*ProjectSearchesV1ListProjectSearchNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ProjectSearchesV1ListProjectSearchNamesNotFound struct {
	Payload interface{}
}

func (o *ProjectSearchesV1ListProjectSearchNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/searches/names][%d] projectSearchesV1ListProjectSearchNamesNotFound  %+v", 404, o.Payload)
}

func (o *ProjectSearchesV1ListProjectSearchNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectSearchesV1ListProjectSearchNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectSearchesV1ListProjectSearchNamesDefault creates a ProjectSearchesV1ListProjectSearchNamesDefault with default headers values
func NewProjectSearchesV1ListProjectSearchNamesDefault(code int) *ProjectSearchesV1ListProjectSearchNamesDefault {
	return &ProjectSearchesV1ListProjectSearchNamesDefault{
		_statusCode: code,
	}
}

/*ProjectSearchesV1ListProjectSearchNamesDefault handles this case with default header values.

An unexpected error response
*/
type ProjectSearchesV1ListProjectSearchNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the project searches v1 list project search names default response
func (o *ProjectSearchesV1ListProjectSearchNamesDefault) Code() int {
	return o._statusCode
}

func (o *ProjectSearchesV1ListProjectSearchNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/searches/names][%d] ProjectSearchesV1_ListProjectSearchNames default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectSearchesV1ListProjectSearchNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ProjectSearchesV1ListProjectSearchNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}