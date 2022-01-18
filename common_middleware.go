package sicgolib

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ContentTypeMiddleware ensures all routes in the project uses `application/json` Content-Type in the header
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// CorsMiddleware allows our API to whitelist some host to be able to interact with our API.
func CorsMiddleware(whitelistedUrls map[string]bool) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			// this code block's flow:
			// check if request method is options
			// get incoming request's origin url
			// check if it is one of the whitelistedUrls
			// if it does, then add allowOrigin to one of the whitelistedUrls+

			rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")

			requestOriginUrl := r.Header.Get("Origin")
			log.Printf("INFO CorsMiddleware: received request from %s %v", requestOriginUrl, whitelistedUrls[requestOriginUrl])
			if whitelistedUrls[requestOriginUrl] {
				rw.Header().Set("Access-Control-Allow-Origin", requestOriginUrl)
			}

			if r.Method != http.MethodOptions {
				next.ServeHTTP(rw, r)
				return
			}

			rw.Write([]byte("okok"))
		})
	}
}

// ErrorHandlingMiddleware catches panics or exception that is thrown by every routes that came, it will return the proper response to user.
// this way, there will no need of a very verboose and complicated error handling flow, as everything else will be handled in the backend side
func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case *BaseResponse:
					rw.WriteHeader(v.Code)
					NewBaseResponse(
						v.Code,
						v.Message,
						v.Errors,
						v.Data,
					).ToJSON(rw)
				default:
					rw.WriteHeader(500)
					NewBaseResponse(
						500,
						RESPONSE_ERROR_RUNTIME_MESSAGE,
						NewErrorResponseData(NewErrorResponseValue("msg", "runtime error")),
						nil,
					)
				}
			}
		}()
		next.ServeHTTP(rw, r)
	})
}
