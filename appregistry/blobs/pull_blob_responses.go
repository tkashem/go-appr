// Code generated by go-swagger; DO NOT EDIT.

package blobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/operator-framework/go-appr/models"
)

// PullBlobReader is a Reader for the PullBlob structure.
type PullBlobReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *PullBlobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPullBlobOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPullBlobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPullBlobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPullBlobOK creates a PullBlobOK with default headers values
func NewPullBlobOK(writer io.Writer) *PullBlobOK {
	return &PullBlobOK{
		Payload: writer,
	}
}

/*PullBlobOK handles this case with default header values.

successful operation
*/
type PullBlobOK struct {
	Payload io.Writer
}

func (o *PullBlobOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}][%d] pullBlobOK  %+v", 200, o.Payload)
}

func (o *PullBlobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPullBlobUnauthorized creates a PullBlobUnauthorized with default headers values
func NewPullBlobUnauthorized() *PullBlobUnauthorized {
	return &PullBlobUnauthorized{}
}

/*PullBlobUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type PullBlobUnauthorized struct {
	Payload *models.Error
}

func (o *PullBlobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}][%d] pullBlobUnauthorized  %+v", 401, o.Payload)
}

func (o *PullBlobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPullBlobNotFound creates a PullBlobNotFound with default headers values
func NewPullBlobNotFound() *PullBlobNotFound {
	return &PullBlobNotFound{}
}

/*PullBlobNotFound handles this case with default header values.

Package not found
*/
type PullBlobNotFound struct {
	Payload *models.Error
}

func (o *PullBlobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}][%d] pullBlobNotFound  %+v", 404, o.Payload)
}

func (o *PullBlobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
