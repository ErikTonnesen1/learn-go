package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Create a deferred function (which will always run in the event of a panic)
		defer func() {
			//use built-in recover() function to check if a panic occurred
			//if panic did happen, reocver() will return the panic value
			//if panid didn't happen, it will return nil
			pv := recover()
			if pv != nil {
				//IF there was a panic, set the headers below. This acts as a trigger to make Go's http server automatically close the current connection after the response has been sent
				w.Header().Set("Connection", "close")

				//The value returned by recover() has the type any, so we use the fmt.Errorf() with the %v verb to coerce it into an error and call our serveErrorResponse() helper
				//In turn this will log the error at the error level and return a 500 response
				app.serveErrorResponse(w, r, fmt.Errorf("%v", pv))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
