package router

import (
	"log"
	"net/http"
	"os"

	"github.com/semihsemih/kv-store/internal/api"
	"github.com/semihsemih/kv-store/internal/storage"
)

// Router Define "Router" type
type Router struct {
	router        *http.ServeMux
	store         storage.Storage
	serverLogFile *os.File
}

// New Initialize new Router
func New(s storage.Storage) *Router {
	// Create and open file to write details about incoming http requests
	serverLogFile, err := os.OpenFile("./store/server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Server log file can not opened.")
	}

	// Initialize new router
	r := &Router{
		router:        http.NewServeMux(),
		store:         s,
		serverLogFile: serverLogFile,
	}

	r.initRoutes()

	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Check request methods and only allow "GET" method
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Set where to write the request details and write request log
	log.SetOutput(r.serverLogFile)
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)

	r.router.ServeHTTP(w, req)
}

// Set routes in app
func (r *Router) initRoutes() {
	// Handle root path
	r.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// The "/" pattern matches everything, so we need to check that we're at the root here.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		_, err := w.Write([]byte("This application is a simple key value store. Please read README file."))
		if err != nil {
			return
		}
	})

	r.router.Handle("/set", api.Set(r.store))
	r.router.Handle("/get", api.Get(r.store))
	r.router.Handle("/flush", api.Flush(r.store))
}
