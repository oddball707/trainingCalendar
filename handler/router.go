package handler

import (
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

func (h *Handler) NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(CORS)

	router.Get("/api/health", h.HealthHandler)
	router.Get("/api/readiness", h.ReadinessHandler)
	router.Post("/api/create", h.CreateIcal)
	router.Post("/api/show", h.CreateSchedule)

	return router
}


// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func (h *Handler) FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}


func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

