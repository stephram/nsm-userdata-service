// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/stephram/nsm-userdata-service/restapi/operations"
)

//go:generate swagger generate server --target ../../nsm-microservice-golang-userdata --name NsmUserdataService --spec ../spec/userdata-swagger.yaml --model-package restapi/models

func configureFlags(api *operations.NsmUserdataServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NsmUserdataServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api-key" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (ApiKey) api-key from header param [api-key] has not yet been implemented")
	}
	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuth = func(user string, pass string) (interface{}, error) {
		return nil, errors.NotImplemented("basic auth  (Basic) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.GetGameOnResultsHandler == nil {
		api.GetGameOnResultsHandler = operations.GetGameOnResultsHandlerFunc(func(params operations.GetGameOnResultsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetGameOnResults has not yet been implemented")
		})
	}
	if api.GetHealthHandler == nil {
		api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetHealth has not yet been implemented")
		})
	}
	if api.GetInfoHandler == nil {
		api.GetInfoHandler = operations.GetInfoHandlerFunc(func(params operations.GetInfoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetInfo has not yet been implemented")
		})
	}
	if api.GetInteractionsHandler == nil {
		api.GetInteractionsHandler = operations.GetInteractionsHandlerFunc(func(params operations.GetInteractionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetInteractions has not yet been implemented")
		})
	}
	if api.GetUserHandler == nil {
		api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetUser has not yet been implemented")
		})
	}
	if api.PostGameOnResultsHandler == nil {
		api.PostGameOnResultsHandler = operations.PostGameOnResultsHandlerFunc(func(params operations.PostGameOnResultsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .PostGameOnResults has not yet been implemented")
		})
	}
	if api.PostInteractionHandler == nil {
		api.PostInteractionHandler = operations.PostInteractionHandlerFunc(func(params operations.PostInteractionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .PostInteraction has not yet been implemented")
		})
	}
	if api.PostUserHandler == nil {
		api.PostUserHandler = operations.PostUserHandlerFunc(func(params operations.PostUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .PostUser has not yet been implemented")
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
