package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Name        string
	Queries     []string
}

type routes []route

var pathQueryParam = []string{
	"path", "{path}",
}

var apiRoutes = routes{
	route{Method: "GET", Pattern: "/api/users", HandlerFunc: getUsersHandler},
	route{Method: "GET", Pattern: "/api/users/{ID}", HandlerFunc: getUserHandler},
	route{Method: "POST", Pattern: "/api/users", HandlerFunc: createUserHandler},
	route{Method: "PUT", Pattern: "/api/users/{ID}", HandlerFunc: updateUserHandler},
	route{Method: "DELETE", Pattern: "/api/users/{ID}", HandlerFunc: deleteUserHandler},
}

func logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func main() {
	var port = "9000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Println("Server running on port " + port)
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range apiRoutes {
		handler := logger(route.HandlerFunc)
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler).
			Queries(route.Queries...)
	}
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
