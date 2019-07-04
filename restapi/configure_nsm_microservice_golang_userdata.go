// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/artprocessors/nsm-microservice-golang-userdata/internal/auth"

	"github.com/artprocessors/nsm-microservice-golang-userdata/internal/utils"

	"github.com/artprocessors/nsm-microservice-golang-userdata/internal/userdata"

	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	op "github.com/artprocessors/nsm-microservice-golang-userdata/restapi/operations"
)

//go:generate swagger generate server --target ../../nsm-microservice-golang-userdata --name NsmMicroserviceGolangUserdata --spec ../spec/userdata-swagger.yaml --model-package restapi/models

func init() {
	utils.GetLogger().Info("initialised logger")
}

func configureFlags(api *op.NsmMicroserviceGolangUserdataAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *op.NsmMicroserviceGolangUserdataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api-key" header is set
	api.APIKeyAuth = auth.ApiKeyAuth

	// // Applies when the Authorization header is set with the Basic scheme
	api.BasicAuth = auth.BasicAuth

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	//api.APIAuthorizer = security.Authorized()

	userDataService := userdata.New()

	api.GetHealthHandler = op.GetHealthHandlerFunc(func(params op.GetHealthParams, principal interface{}) middleware.Responder {
		return userDataService.GetHealth(params)
	})
	api.GetInfoHandler = op.GetInfoHandlerFunc(func(params op.GetInfoParams, principal interface{}) middleware.Responder {
		return userDataService.GetInfo(params)
	})

	api.GetUserHandler = op.GetUserHandlerFunc(func(params op.GetUserParams, principal interface{}) middleware.Responder {
		return userDataService.GetUser(params)
	})
	api.PostUserHandler = op.PostUserHandlerFunc(func(params op.PostUserParams, principal interface{}) middleware.Responder {
		return userDataService.PostUser(params)
	})

	api.GetInteractionsHandler = op.GetInteractionsHandlerFunc(func(params op.GetInteractionsParams, principal interface{}) middleware.Responder {
		return userDataService.GetInteractions(params)
	})
	api.PostInteractionHandler = op.PostInteractionHandlerFunc(func(params op.PostInteractionParams, principal interface{}) middleware.Responder {
		return userDataService.PostInteraction(params)
	})

	api.GetGameOnResultsHandler = op.GetGameOnResultsHandlerFunc(func(params op.GetGameOnResultsParams, principal interface{}) middleware.Responder {
		return userDataService.GetGameOnResults(params)
	})
	api.PostGameOnResultsHandler = op.PostGameOnResultsHandlerFunc(func(params op.PostGameOnResultsParams, principal interface{}) middleware.Responder {
		return userDataService.PostGameOnResults(params)
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
