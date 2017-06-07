package access_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GenerateAccessTokenURL generates an URL for the generate access token operation
type GenerateAccessTokenURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GenerateAccessTokenURL) WithBasePath(bp string) *GenerateAccessTokenURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GenerateAccessTokenURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GenerateAccessTokenURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/token"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/oauth2"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GenerateAccessTokenURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GenerateAccessTokenURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GenerateAccessTokenURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GenerateAccessTokenURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GenerateAccessTokenURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GenerateAccessTokenURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}