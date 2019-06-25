// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/artprocessors/nsm-microservice-golang-userdata/restapi/operations"
)

//go:generate swagger generate server --target ../../nsm-microservice-golang-userdata --name NsmMicroserviceGolangUserdata --spec ../spec/userdata-swagger.yaml --model-package restapi/models

func configureFlags(api *operations.NsmMicroserviceGolangUserdataAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NsmMicroserviceGolangUserdataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GetAvatarForTokenIDHandler == nil {
		api.GetAvatarForTokenIDHandler = operations.GetAvatarForTokenIDHandlerFunc(func(params operations.GetAvatarForTokenIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetAvatarForTokenID has not yet been implemented")
		})
	}
	if api.GetHealthHandler == nil {
		api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetHealth has not yet been implemented")
		})
	}
	if api.GetInteractionsHandler == nil {
		api.GetInteractionsHandler = operations.GetInteractionsHandlerFunc(func(params operations.GetInteractionsParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetInteractions has not yet been implemented")
		})
	}
	if api.GetStatusByTokenIDHandler == nil {
		api.GetStatusByTokenIDHandler = operations.GetStatusByTokenIDHandlerFunc(func(params operations.GetStatusByTokenIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetStatusByTokenID has not yet been implemented")
		})
	}
	if api.PostAvatarForTokenIDHandler == nil {
		api.PostAvatarForTokenIDHandler = operations.PostAvatarForTokenIDHandlerFunc(func(params operations.PostAvatarForTokenIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .PostAvatarForTokenID has not yet been implemented")
		})
	}
	if api.PostInteractionsHandler == nil {
		api.PostInteractionsHandler = operations.PostInteractionsHandlerFunc(func(params operations.PostInteractionsParams) middleware.Responder {
			return middleware.NotImplemented("operation .PostInteractions has not yet been implemented")
		})
	}

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
func configureServer(s *http.Server, scheme, addr string) {
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
