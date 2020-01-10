package web

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

// App is the entrypoint into your application
type App struct {
	log *log.Logger
	mux *chi.Mux
}

func NewApp(log *log.Logger) *App {
	return &App{
		log: log,
		mux: chi.NewRouter(),
	}
}

func (a *App) Handle(method, url string, h http.HandlerFunc) {
	a.mux.MethodFunc(method, url, h)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}


