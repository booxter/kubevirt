// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteChannelReleaseParams creates a new DeleteChannelReleaseParams object
// with the default values initialized.
func NewDeleteChannelReleaseParams() *DeleteChannelReleaseParams {
	var ()
	return &DeleteChannelReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteChannelReleaseParamsWithTimeout creates a new DeleteChannelReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteChannelReleaseParamsWithTimeout(timeout time.Duration) *DeleteChannelReleaseParams {
	var ()
	return &DeleteChannelReleaseParams{

		timeout: timeout,
	}
}

// NewDeleteChannelReleaseParamsWithContext creates a new DeleteChannelReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteChannelReleaseParamsWithContext(ctx context.Context) *DeleteChannelReleaseParams {
	var ()
	return &DeleteChannelReleaseParams{

		Context: ctx,
	}
}

// NewDeleteChannelReleaseParamsWithHTTPClient creates a new DeleteChannelReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteChannelReleaseParamsWithHTTPClient(client *http.Client) *DeleteChannelReleaseParams {
	var ()
	return &DeleteChannelReleaseParams{
		HTTPClient: client,
	}
}

/*DeleteChannelReleaseParams contains all the parameters to send to the API endpoint
for the delete channel release operation typically these are written to a http.Request
*/
type DeleteChannelReleaseParams struct {

	/*Channel
	  channel name

	*/
	Channel string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  full package name

	*/
	Package string
	/*Release
	  Release name

	*/
	Release string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete channel release params
func (o *DeleteChannelReleaseParams) WithTimeout(timeout time.Duration) *DeleteChannelReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete channel release params
func (o *DeleteChannelReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete channel release params
func (o *DeleteChannelReleaseParams) WithContext(ctx context.Context) *DeleteChannelReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete channel release params
func (o *DeleteChannelReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete channel release params
func (o *DeleteChannelReleaseParams) WithHTTPClient(client *http.Client) *DeleteChannelReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete channel release params
func (o *DeleteChannelReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannel adds the channel to the delete channel release params
func (o *DeleteChannelReleaseParams) WithChannel(channel string) *DeleteChannelReleaseParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the delete channel release params
func (o *DeleteChannelReleaseParams) SetChannel(channel string) {
	o.Channel = channel
}

// WithNamespace adds the namespace to the delete channel release params
func (o *DeleteChannelReleaseParams) WithNamespace(namespace string) *DeleteChannelReleaseParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete channel release params
func (o *DeleteChannelReleaseParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the delete channel release params
func (o *DeleteChannelReleaseParams) WithPackage(packageVar string) *DeleteChannelReleaseParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the delete channel release params
func (o *DeleteChannelReleaseParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WithRelease adds the release to the delete channel release params
func (o *DeleteChannelReleaseParams) WithRelease(release string) *DeleteChannelReleaseParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the delete channel release params
func (o *DeleteChannelReleaseParams) SetRelease(release string) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteChannelReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel
	if err := r.SetPathParam("channel", o.Channel); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	// path param release
	if err := r.SetPathParam("release", o.Release); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
