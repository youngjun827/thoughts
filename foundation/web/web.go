// Package web contains a small web framework extension to the Chi Router.
package web

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

type App struct {
    Mux *chi.Mux
    shutdown chan os.Signal
	mw []Middleware
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
    r := chi.NewRouter()
    return &App{
        Mux:      r,
        shutdown: shutdown,
		mw: mw,
    }
}

func HandlerAdapter(handler Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := handler(r.Context(), w, r)
        if err != nil {
            // Handle the error as needed
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
    }
}

func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {
    wrappedHandler := wrapMiddleware(mw, handler)
    wrappedHandler = wrapMiddleware(a.mw, wrappedHandler)

    customHandler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		err := wrappedHandler(ctx, w, r)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
        return nil
    }

    a.Mux.MethodFunc(method, path, HandlerAdapter(customHandler))
}
