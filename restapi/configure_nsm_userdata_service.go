// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

	"github.com/stephram/nsm-userdata-service/internal/utils"

	"github.com/stephram/nsm-userdata-service/internal/auth"

	"github.com/stephram/nsm-userdata-service/internal/userdata"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/stephram/nsm-userdata-service/restapi/operations"
)

//go:generate swagger generate server --target ../../nsm-userdata-service --name NsmUserdataService --spec ../spec/userdata-swagger.yaml --model-package restapi/models

func init() {
	utils.GetLogger()
}

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
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api-key" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return auth.ApiKeyAuth(token)
	}
	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuth = func(user string, pass string) (interface{}, error) {
		return auth.BasicAuth(user, pass)
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	userdata := userdata.New()

	api.GetGameOnResultsHandler = operations.GetGameOnResultsHandlerFunc(func(params operations.GetGameOnResultsParams, principal interface{}) middleware.Responder {
		return userdata.GetGameOnResults(params)
	})
	api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams, principal interface{}) middleware.Responder {
		return userdata.GetHealth(params)
	})
	api.GetInfoHandler = operations.GetInfoHandlerFunc(func(params operations.GetInfoParams, principal interface{}) middleware.Responder {
		return userdata.GetInfo(params)
	})
	api.GetInteractionsHandler = operations.GetInteractionsHandlerFunc(func(params operations.GetInteractionsParams, principal interface{}) middleware.Responder {
		return userdata.GetInteractions(params)
	})
	api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams, principal interface{}) middleware.Responder {
		return userdata.GetUser(params)
	})
	api.PostGameOnResultsHandler = operations.PostGameOnResultsHandlerFunc(func(params operations.PostGameOnResultsParams, principal interface{}) middleware.Responder {
		return userdata.PostGameOnResults(params)
	})
	api.PostInteractionHandler = operations.PostInteractionHandlerFunc(func(params operations.PostInteractionParams, principal interface{}) middleware.Responder {
		return userdata.PostInteraction(params)
	})

	api.PostUserHandler = operations.PostUserHandlerFunc(func(params operations.PostUserParams, principal interface{}) middleware.Responder {
		return userdata.PostUser(params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func fileServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/userdata") {
			next.ServeHTTP(w, r)
		} else {
			http.FileServer(http.Dir("./swagger-ui/index.html")).ServeHTTP(w, r)
		}
	})
}

// See one of the many discussion about how to best acheive swagger-ui servings.
// A good place to stat: https://github.com/go-swagger/go-swagger/issues/477
func uiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/api/help" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui"))).ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
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
	// return fileServerMiddleware(handler)
	return uiMiddleware(handler)
	// return handler
}
