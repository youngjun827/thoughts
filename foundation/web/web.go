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
}

func NewApp(shutdown chan os.Signal) *App {
    r := chi.NewRouter()
    return &App{
        Mux:      r,
        shutdown: shutdown,
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
func (a *App) Handle(method string, path string, handler Handler) {
    a.Mux.MethodFunc(method, path, HandlerAdapter(handler))
}