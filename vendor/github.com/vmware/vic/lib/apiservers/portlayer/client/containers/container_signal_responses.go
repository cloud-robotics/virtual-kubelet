package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// ContainerSignalReader is a Reader for the ContainerSignal structure.
type ContainerSignalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerSignalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerSignalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerSignalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerSignalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerSignalOK creates a ContainerSignalOK with default headers values
func NewContainerSignalOK() *ContainerSignalOK {
	return &ContainerSignalOK{}
}

/*ContainerSignalOK handles this case with default header values.

OK
*/
type ContainerSignalOK struct {
}

func (o *ContainerSignalOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/signal][%d] containerSignalOK ", 200)
}

func (o *ContainerSignalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerSignalNotFound creates a ContainerSignalNotFound with default headers values
func NewContainerSignalNotFound() *ContainerSignalNotFound {
	return &ContainerSignalNotFound{}
}

/*ContainerSignalNotFound handles this case with default header values.

Container not found
*/
type ContainerSignalNotFound struct {
	Payload *models.Error
}

func (o *ContainerSignalNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/signal][%d] containerSignalNotFound  %+v", 404, o.Payload)
}

func (o *ContainerSignalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerSignalInternalServerError creates a ContainerSignalInternalServerError with default headers values
func NewContainerSignalInternalServerError() *ContainerSignalInternalServerError {
	return &ContainerSignalInternalServerError{}
}

/*ContainerSignalInternalServerError handles this case with default header values.

Failed to signal container
*/
type ContainerSignalInternalServerError struct {
	Payload *models.Error
}

func (o *ContainerSignalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/signal][%d] containerSignalInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerSignalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
