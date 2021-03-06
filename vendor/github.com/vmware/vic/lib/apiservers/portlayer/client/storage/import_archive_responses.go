package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ImportArchiveReader is a Reader for the ImportArchive structure.
type ImportArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImportArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewImportArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewImportArchiveConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewImportArchiveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 423:
		result := NewImportArchiveLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImportArchiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewImportArchiveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImportArchiveOK creates a ImportArchiveOK with default headers values
func NewImportArchiveOK() *ImportArchiveOK {
	return &ImportArchiveOK{}
}

/*ImportArchiveOK handles this case with default header values.

OK
*/
type ImportArchiveOK struct {
}

func (o *ImportArchiveOK) Error() string {
	return fmt.Sprintf("[POST /archive][%d] importArchiveOK ", 200)
}

func (o *ImportArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportArchiveNotFound creates a ImportArchiveNotFound with default headers values
func NewImportArchiveNotFound() *ImportArchiveNotFound {
	return &ImportArchiveNotFound{}
}

/*ImportArchiveNotFound handles this case with default header values.

Supplied target not found
*/
type ImportArchiveNotFound struct {
}

func (o *ImportArchiveNotFound) Error() string {
	return fmt.Sprintf("[POST /archive][%d] importArchiveNotFound ", 404)
}

func (o *ImportArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportArchiveConflict creates a ImportArchiveConflict with default headers values
func NewImportArchiveConflict() *ImportArchiveConflict {
	return &ImportArchiveConflict{}
}

/*ImportArchiveConflict handles this case with default header values.

Unexpected resource conflict
*/
type ImportArchiveConflict struct {
}

func (o *ImportArchiveConflict) Error() string {
	return fmt.Sprintf("[POST /archive][%d] importArchiveConflict ", 409)
}

func (o *ImportArchiveConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportArchiveUnprocessableEntity creates a ImportArchiveUnprocessableEntity with default headers values
func NewImportArchiveUnprocessableEntity() *ImportArchiveUnprocessableEntity {
	return &ImportArchiveUnprocessableEntity{}
}

/*ImportArchiveUnprocessableEntity handles this case with default header values.

Format error in supplied filter or archive
*/
type ImportArchiveUnprocessableEntity struct {
}

func (o *ImportArchiveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /archive][%d] importArchiveUnprocessableEntity ", 422)
}

func (o *ImportArchiveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportArchiveLocked creates a ImportArchiveLocked with default headers values
func NewImportArchiveLocked() *ImportArchiveLocked {
	return &ImportArchiveLocked{}
}

/*ImportArchiveLocked handles this case with default header values.

Device or resource is locked
*/
type ImportArchiveLocked struct {
}

func (o *ImportArchiveLocked) Error() string {
	return fmt.Sprintf("[POST /archive][%d] importArchiveLocked ", 423)
}

func (o *ImportArchiveLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportArchiveInternalServerError creates a ImportArchiveInternalServerError with default headers values
func NewImportArchiveInternalServerError() *ImportArchiveInternalServerError {
	return &ImportArchiveInternalServerError{}
}

/*ImportArchiveInternalServerError handles this case with default header values.

failed to export tar archive from target
*/
type ImportArchiveInternalServerError struct {
}

func (o *ImportArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /archive][%d] importArchiveInternalServerError ", 500)
}

func (o *ImportArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportArchiveDefault creates a ImportArchiveDefault with default headers values
func NewImportArchiveDefault(code int) *ImportArchiveDefault {
	return &ImportArchiveDefault{
		_statusCode: code,
	}
}

/*ImportArchiveDefault handles this case with default header values.

error
*/
type ImportArchiveDefault struct {
	_statusCode int
}

// Code gets the status code for the import archive default response
func (o *ImportArchiveDefault) Code() int {
	return o._statusCode
}

func (o *ImportArchiveDefault) Error() string {
	return fmt.Sprintf("[POST /archive][%d] ImportArchive default ", o._statusCode)
}

func (o *ImportArchiveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
