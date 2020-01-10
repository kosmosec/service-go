package web

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

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

func (a *App) Handle(method, url string, h Handler) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)

		if err != nil {
			a.log.Printf("ERROR : %+v", err)

			if err := RespondError(w, err); err != nil {
				a.log.Printf("ERROR : %v", err)
			}
		}
	}
	a.mux.MethodFunc(method, url, fn)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}


