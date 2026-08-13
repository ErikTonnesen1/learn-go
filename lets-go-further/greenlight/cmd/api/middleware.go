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
			//if panid didn't happen 

		}
	})
}
