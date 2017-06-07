package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/ridewindx/loong/pkg/api/upload/restapi/operations"
	"github.com/ridewindx/loong/pkg/api/upload/restapi/operations/file_upload"
	"github.com/ridewindx/loong/pkg/api/upload/restapi/operations/supercharged_upload"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ..\pkg\api\upload --name LoongUploadApi --spec ..\api\upload-v2.yaml

func configureFlags(api *operations.LoongUploadAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LoongUploadAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinConsumer = runtime.ByteStreamConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.OAuth2SecurityAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (OAuth2Security) has not yet been implemented")
	}

	api.SuperchargedUploadAbortUploadSessionHandler = supercharged_upload.AbortUploadSessionHandlerFunc(func(params supercharged_upload.AbortUploadSessionParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.AbortUploadSession has not yet been implemented")
	})
	api.SuperchargedUploadCommitUploadSessionHandler = supercharged_upload.CommitUploadSessionHandlerFunc(func(params supercharged_upload.CommitUploadSessionParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.CommitUploadSession has not yet been implemented")
	})
	api.SuperchargedUploadCreateUploadSessionHandler = supercharged_upload.CreateUploadSessionHandlerFunc(func(params supercharged_upload.CreateUploadSessionParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.CreateUploadSession has not yet been implemented")
	})
	api.SuperchargedUploadCreateUploadSessionNewVersionHandler = supercharged_upload.CreateUploadSessionNewVersionHandlerFunc(func(params supercharged_upload.CreateUploadSessionNewVersionParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.CreateUploadSessionNewVersion has not yet been implemented")
	})
	api.SuperchargedUploadGetUploadSessionHandler = supercharged_upload.GetUploadSessionHandlerFunc(func(params supercharged_upload.GetUploadSessionParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.GetUploadSession has not yet been implemented")
	})
	api.SuperchargedUploadListPartsHandler = supercharged_upload.ListPartsHandlerFunc(func(params supercharged_upload.ListPartsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.ListParts has not yet been implemented")
	})
	api.FileUploadUploadFileHandler = file_upload.UploadFileHandlerFunc(func(params file_upload.UploadFileParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation file_upload.UploadFile has not yet been implemented")
	})
	api.FileUploadUploadFileVersionHandler = file_upload.UploadFileVersionHandlerFunc(func(params file_upload.UploadFileVersionParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation file_upload.UploadFileVersion has not yet been implemented")
	})
	api.SuperchargedUploadUploadPartHandler = supercharged_upload.UploadPartHandlerFunc(func(params supercharged_upload.UploadPartParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation supercharged_upload.UploadPart has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
