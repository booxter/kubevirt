// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/operator-framework/go-appr/models"
)

// DeleteChannelReleaseReader is a Reader for the DeleteChannelRelease structure.
type DeleteChannelReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteChannelReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteChannelReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteChannelReleaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteChannelReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteChannelReleaseOK creates a DeleteChannelReleaseOK with default headers values
func NewDeleteChannelReleaseOK() *DeleteChannelReleaseOK {
	return &DeleteChannelReleaseOK{}
}

/*DeleteChannelReleaseOK handles this case with default header values.

successful operation
*/
type DeleteChannelReleaseOK struct {
	Payload []*models.Channel
}

func (o *DeleteChannelReleaseOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/channels/{channel}/{release}][%d] deleteChannelReleaseOK  %+v", 200, o.Payload)
}

func (o *DeleteChannelReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChannelReleaseUnauthorized creates a DeleteChannelReleaseUnauthorized with default headers values
func NewDeleteChannelReleaseUnauthorized() *DeleteChannelReleaseUnauthorized {
	return &DeleteChannelReleaseUnauthorized{}
}

/*DeleteChannelReleaseUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type DeleteChannelReleaseUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteChannelReleaseUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/channels/{channel}/{release}][%d] deleteChannelReleaseUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteChannelReleaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChannelReleaseNotFound creates a DeleteChannelReleaseNotFound with default headers values
func NewDeleteChannelReleaseNotFound() *DeleteChannelReleaseNotFound {
	return &DeleteChannelReleaseNotFound{}
}

/*DeleteChannelReleaseNotFound handles this case with default header values.

Resource not found
*/
type DeleteChannelReleaseNotFound struct {
	Payload *models.Error
}

func (o *DeleteChannelReleaseNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/channels/{channel}/{release}][%d] deleteChannelReleaseNotFound  %+v", 404, o.Payload)
}

func (o *DeleteChannelReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
