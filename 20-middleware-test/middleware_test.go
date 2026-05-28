package middlewaretest

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Execute")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Execute")
		fmt.Fprint(w, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Execute")
		panic("Ups")
	})

	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorhandler := &ErrorHandler{
		Handler: LogMiddleware,
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: errorhandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
